package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/caos/logging"
	"golang.org/x/text/language"

	"github.com/caos/zitadel/internal/api/authz"
	"github.com/caos/zitadel/internal/command"
	sd "github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/crypto"
	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/v1"
	"github.com/caos/zitadel/internal/eventstore/v1/models"
	"github.com/caos/zitadel/internal/eventstore/v1/query"
	"github.com/caos/zitadel/internal/eventstore/v1/spooler"
	"github.com/caos/zitadel/internal/i18n"
	iam_model "github.com/caos/zitadel/internal/iam/model"
	iam_es_model "github.com/caos/zitadel/internal/iam/repository/view/model"
	"github.com/caos/zitadel/internal/notification/types"
	es_model "github.com/caos/zitadel/internal/user/repository/eventsourcing/model"
	"github.com/caos/zitadel/internal/user/repository/view"
	"github.com/caos/zitadel/internal/user/repository/view/model"
)

const (
	notificationTable            = "notification.notifications"
	NotifyUserID                 = "NOTIFICATION"
	labelPolicyTableOrg          = "management.label_policies"
	labelPolicyTableDef          = "adminapi.label_policies"
	mailTemplateTableOrg         = "management.mail_templates"
	mailTemplateTableDef         = "adminapi.mail_templates"
	messageTextTableOrg          = "management.message_texts"
	messageTextTableDef          = "adminapi.message_texts"
	messageTextTypeDomainClaimed = "DomainClaimed"
	messageTextTypeInitCode      = "InitCode"
	messageTextTypePasswordReset = "PasswordReset"
	messageTextTypeVerifyEmail   = "VerifyEmail"
	messageTextTypeVerifyPhone   = "VerifyPhone"
)

type Notification struct {
	handler
	command        *command.Commands
	systemDefaults sd.SystemDefaults
	AesCrypto      crypto.EncryptionAlgorithm
	i18n           *i18n.Translator
	statikDir      http.FileSystem
	subscription   *v1.Subscription
	apiDomain      string
}

func newNotification(
	handler handler,
	command *command.Commands,
	defaults sd.SystemDefaults,
	aesCrypto crypto.EncryptionAlgorithm,
	translator *i18n.Translator,
	statikDir http.FileSystem,
	apiDomain string,
) *Notification {
	h := &Notification{
		handler:        handler,
		command:        command,
		systemDefaults: defaults,
		i18n:           translator,
		statikDir:      statikDir,
		AesCrypto:      aesCrypto,
		apiDomain:      apiDomain,
	}

	h.subscribe()

	return h
}

func (k *Notification) subscribe() {
	k.subscription = k.es.Subscribe(k.AggregateTypes()...)
	go func() {
		for event := range k.subscription.Events {
			query.ReduceEvent(k, event)
		}
	}()
}

func (n *Notification) ViewModel() string {
	return notificationTable
}

func (_ *Notification) AggregateTypes() []models.AggregateType {
	return []models.AggregateType{es_model.UserAggregate}
}

func (n *Notification) CurrentSequence() (uint64, error) {
	sequence, err := n.view.GetLatestNotificationSequence()
	if err != nil {
		return 0, err
	}
	return sequence.CurrentSequence, nil
}

func (n *Notification) EventQuery() (*models.SearchQuery, error) {
	sequence, err := n.view.GetLatestNotificationSequence()
	if err != nil {
		return nil, err
	}
	return view.UserQuery(sequence.CurrentSequence), nil
}

func (n *Notification) Reduce(event *models.Event) (err error) {
	switch event.Type {
	case es_model.InitializedUserCodeAdded,
		es_model.InitializedHumanCodeAdded:
		err = n.handleInitUserCode(event)
	case es_model.UserEmailCodeAdded,
		es_model.HumanEmailCodeAdded:
		err = n.handleEmailVerificationCode(event)
	case es_model.UserPhoneCodeAdded,
		es_model.HumanPhoneCodeAdded:
		err = n.handlePhoneVerificationCode(event)
	case es_model.UserPasswordCodeAdded,
		es_model.HumanPasswordCodeAdded:
		err = n.handlePasswordCode(event)
	case es_model.DomainClaimed:
		err = n.handleDomainClaimed(event)
	}
	if err != nil {
		return err
	}
	return n.view.ProcessedNotificationSequence(event)
}

func (n *Notification) handleInitUserCode(event *models.Event) (err error) {
	initCode := new(es_model.InitUserCode)
	if err := initCode.SetData(event); err != nil {
		return err
	}
	alreadyHandled, err := n.checkIfCodeAlreadyHandledOrExpired(event, initCode.Expiry,
		es_model.InitializedUserCodeAdded, es_model.InitializedUserCodeSent,
		es_model.InitializedHumanCodeAdded, es_model.InitializedHumanCodeSent)
	if err != nil || alreadyHandled {
		return err
	}
	ctx := getSetNotifyContextData(event.ResourceOwner)
	colors, err := n.getLabelPolicy(ctx)
	if err != nil {
		return err
	}

	template, err := n.getMailTemplate(ctx)
	if err != nil {
		return err
	}

	user, err := n.getUserByID(event.AggregateID)
	if err != nil {
		return err
	}

	text, err := n.getMessageText(user, messageTextTypeInitCode, user.PreferredLanguage)
	if err != nil {
		return err
	}

	err = types.SendUserInitCode(string(template.Template), text, user, initCode, n.systemDefaults, n.AesCrypto, colors, n.apiDomain)
	if err != nil {
		return err
	}
	return n.command.HumanInitCodeSent(ctx, event.ResourceOwner, event.AggregateID)
}

func (n *Notification) handlePasswordCode(event *models.Event) (err error) {
	pwCode := new(es_model.PasswordCode)
	if err := pwCode.SetData(event); err != nil {
		return err
	}
	alreadyHandled, err := n.checkIfCodeAlreadyHandledOrExpired(event, pwCode.Expiry,
		es_model.UserPasswordCodeAdded, es_model.UserPasswordCodeSent,
		es_model.HumanPasswordCodeAdded, es_model.HumanPasswordCodeSent)
	if err != nil || alreadyHandled {
		return err
	}
	ctx := getSetNotifyContextData(event.ResourceOwner)
	colors, err := n.getLabelPolicy(ctx)
	if err != nil {
		return err
	}

	template, err := n.getMailTemplate(ctx)
	if err != nil {
		return err
	}

	user, err := n.getUserByID(event.AggregateID)
	if err != nil {
		return err
	}

	text, err := n.getMessageText(user, messageTextTypePasswordReset, user.PreferredLanguage)
	if err != nil {
		return err
	}
	err = types.SendPasswordCode(string(template.Template), text, user, pwCode, n.systemDefaults, n.AesCrypto, colors, n.apiDomain)
	if err != nil {
		return err
	}
	return n.command.PasswordCodeSent(ctx, event.ResourceOwner, event.AggregateID)
}

func (n *Notification) handleEmailVerificationCode(event *models.Event) (err error) {
	emailCode := new(es_model.EmailCode)
	if err := emailCode.SetData(event); err != nil {
		return err
	}
	alreadyHandled, err := n.checkIfCodeAlreadyHandledOrExpired(event, emailCode.Expiry,
		es_model.UserEmailCodeAdded, es_model.UserEmailCodeSent,
		es_model.HumanEmailCodeAdded, es_model.HumanEmailCodeSent)
	if err != nil || alreadyHandled {
		return nil
	}
	ctx := getSetNotifyContextData(event.ResourceOwner)
	colors, err := n.getLabelPolicy(ctx)
	if err != nil {
		return err
	}

	template, err := n.getMailTemplate(ctx)
	if err != nil {
		return err
	}

	user, err := n.getUserByID(event.AggregateID)
	if err != nil {
		return err
	}

	text, err := n.getMessageText(user, messageTextTypeVerifyEmail, user.PreferredLanguage)
	if err != nil {
		return err
	}

	err = types.SendEmailVerificationCode(string(template.Template), text, user, emailCode, n.systemDefaults, n.AesCrypto, colors, n.apiDomain)
	if err != nil {
		return err
	}
	return n.command.HumanEmailVerificationCodeSent(ctx, event.ResourceOwner, event.AggregateID)
}

func (n *Notification) handlePhoneVerificationCode(event *models.Event) (err error) {
	phoneCode := new(es_model.PhoneCode)
	if err := phoneCode.SetData(event); err != nil {
		return err
	}
	alreadyHandled, err := n.checkIfCodeAlreadyHandledOrExpired(event, phoneCode.Expiry,
		es_model.UserPhoneCodeAdded, es_model.UserPhoneCodeSent,
		es_model.HumanPhoneCodeAdded, es_model.HumanPhoneCodeSent)
	if err != nil || alreadyHandled {
		return nil
	}
	user, err := n.getUserByID(event.AggregateID)
	if err != nil {
		return err
	}
	text, err := n.getMessageText(user, messageTextTypeVerifyPhone, user.PreferredLanguage)
	if err != nil {
		return err
	}
	err = types.SendPhoneVerificationCode(text, user, phoneCode, n.systemDefaults, n.AesCrypto)
	if err != nil {
		return err
	}
	return n.command.HumanPhoneVerificationCodeSent(getSetNotifyContextData(event.ResourceOwner), event.ResourceOwner, event.AggregateID)
}

func (n *Notification) handleDomainClaimed(event *models.Event) (err error) {
	alreadyHandled, err := n.checkIfAlreadyHandled(event.AggregateID, event.Sequence, es_model.DomainClaimed, es_model.DomainClaimedSent)
	if err != nil || alreadyHandled {
		return nil
	}
	data := make(map[string]string)
	if err := json.Unmarshal(event.Data, &data); err != nil {
		logging.Log("HANDLE-Gghq2").WithError(err).Error("could not unmarshal event data")
		return errors.ThrowInternal(err, "HANDLE-7hgj3", "could not unmarshal event")
	}
	user, err := n.getUserByID(event.AggregateID)
	if err != nil {
		return err
	}
	if user.LastEmail == "" {
		return nil
	}
	ctx := getSetNotifyContextData(event.ResourceOwner)
	colors, err := n.getLabelPolicy(ctx)
	if err != nil {
		return err
	}

	template, err := n.getMailTemplate(ctx)
	if err != nil {
		return err
	}

	text, err := n.getMessageText(user, messageTextTypeDomainClaimed, user.PreferredLanguage)
	if err != nil {
		return err
	}
	err = types.SendDomainClaimed(string(template.Template), text, user, data["userName"], n.systemDefaults, colors, n.apiDomain)
	if err != nil {
		return err
	}
	return n.command.UserDomainClaimedSent(ctx, event.ResourceOwner, event.AggregateID)
}

func (n *Notification) checkIfCodeAlreadyHandledOrExpired(event *models.Event, expiry time.Duration, eventTypes ...models.EventType) (bool, error) {
	if event.CreationDate.Add(expiry).Before(time.Now().UTC()) {
		return true, nil
	}
	return n.checkIfAlreadyHandled(event.AggregateID, event.Sequence, eventTypes...)
}

func (n *Notification) checkIfAlreadyHandled(userID string, sequence uint64, eventTypes ...models.EventType) (bool, error) {
	events, err := n.getUserEvents(userID, sequence)
	if err != nil {
		return false, err
	}
	for _, event := range events {
		for _, eventType := range eventTypes {
			if event.Type == eventType {
				return true, nil
			}
		}
	}
	return false, nil
}

func (n *Notification) getUserEvents(userID string, sequence uint64) ([]*models.Event, error) {
	query, err := view.UserByIDQuery(userID, sequence)
	if err != nil {
		return nil, err
	}

	return n.es.FilterEvents(context.Background(), query)
}

func (n *Notification) OnError(event *models.Event, err error) error {
	logging.LogWithFields("SPOOL-s9opc", "id", event.AggregateID, "sequence", event.Sequence).WithError(err).Warn("something went wrong in notification handler")
	return spooler.HandleError(event, err, n.view.GetLatestNotificationFailedEvent, n.view.ProcessedNotificationFailedEvent, n.view.ProcessedNotificationSequence, n.errorCountUntilSkip)
}

func (n *Notification) OnSuccess() error {
	return spooler.HandleSuccess(n.view.UpdateNotificationSpoolerRunTimestamp)
}

func getSetNotifyContextData(orgID string) context.Context {
	return authz.SetCtxData(context.Background(), authz.CtxData{UserID: NotifyUserID, OrgID: orgID})
}

// Read organization specific colors
func (n *Notification) getLabelPolicy(ctx context.Context) (*iam_model.LabelPolicyView, error) {
	// read from Org
	policy, err := n.view.LabelPolicyByAggregateIDAndState(authz.GetCtxData(ctx).OrgID, labelPolicyTableOrg, int32(domain.LabelPolicyStateActive))
	if errors.IsNotFound(err) {
		// read from default
		policy, err = n.view.LabelPolicyByAggregateIDAndState(n.systemDefaults.IamID, labelPolicyTableDef, int32(domain.LabelPolicyStateActive))
		if err != nil {
			return nil, err
		}
		policy.Default = true
	}
	if err != nil {
		return nil, err
	}
	return iam_es_model.LabelPolicyViewToModel(policy), err
}

// Read organization specific template
func (n *Notification) getMailTemplate(ctx context.Context) (*iam_model.MailTemplateView, error) {
	// read from Org
	template, err := n.view.MailTemplateByAggregateID(authz.GetCtxData(ctx).OrgID, mailTemplateTableOrg)
	if errors.IsNotFound(err) {
		// read from default
		template, err = n.view.MailTemplateByAggregateID(n.systemDefaults.IamID, mailTemplateTableDef)
		if err != nil {
			return nil, err
		}
		template.Default = true
	}
	if err != nil {
		return nil, err
	}
	return iam_es_model.MailTemplateViewToModel(template), err
}

// Read organization specific texts
func (n *Notification) getMessageText(user *model.NotifyUser, textType, lang string) (*iam_model.MessageTextView, error) {
	langTag := language.Make(lang)
	if langTag == language.Und {
		langTag = language.English
	}
	langBase, _ := langTag.Base()

	defaultMessageText, err := n.view.MessageTextByIDs(n.systemDefaults.IamID, textType, langBase.String(), messageTextTableDef)
	if err != nil {
		return nil, err
	}
	defaultMessageText.Default = true

	// read from Org
	orgMessageText, err := n.view.MessageTextByIDs(user.ResourceOwner, textType, langBase.String(), messageTextTableOrg)
	if errors.IsNotFound(err) {
		return iam_es_model.MessageTextViewToModel(defaultMessageText), nil
	}
	if err != nil {
		return nil, err
	}
	mergedText := mergeMessageTexts(defaultMessageText, orgMessageText)
	return iam_es_model.MessageTextViewToModel(mergedText), err
}

func (n *Notification) getUserByID(userID string) (*model.NotifyUser, error) {
	user, usrErr := n.view.NotifyUserByID(userID)
	if usrErr != nil && !errors.IsNotFound(usrErr) {
		return nil, usrErr
	}
	if user == nil {
		user = &model.NotifyUser{}
	}
	events, err := n.getUserEvents(userID, user.Sequence)
	if err != nil {
		return user, usrErr
	}
	userCopy := *user
	for _, event := range events {
		if err := userCopy.AppendEvent(event); err != nil {
			return user, nil
		}
	}
	if userCopy.State == int32(model.UserStateDeleted) {
		return nil, errors.ThrowNotFound(nil, "HANDLER-3n8fs", "Errors.User.NotFound")
	}
	return &userCopy, nil
}

func mergeMessageTexts(defaultText *iam_es_model.MessageTextView, orgText *iam_es_model.MessageTextView) *iam_es_model.MessageTextView {
	if orgText.Subject == "" {
		orgText.Subject = defaultText.Subject
	}
	if orgText.Title == "" {
		orgText.Title = defaultText.Title
	}
	if orgText.PreHeader == "" {
		orgText.PreHeader = defaultText.PreHeader
	}
	if orgText.Text == "" {
		orgText.Text = defaultText.Text
	}
	if orgText.Greeting == "" {
		orgText.Greeting = defaultText.Greeting
	}
	if orgText.ButtonText == "" {
		orgText.ButtonText = defaultText.ButtonText
	}
	if orgText.FooterText == "" {
		orgText.FooterText = defaultText.FooterText
	}
	return orgText
}
