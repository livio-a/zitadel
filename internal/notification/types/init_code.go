package types

import (
	"html"

	"github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/crypto"
	iam_model "github.com/caos/zitadel/internal/iam/model"
	"github.com/caos/zitadel/internal/notification/templates"
	es_model "github.com/caos/zitadel/internal/user/repository/eventsourcing/model"
	view_model "github.com/caos/zitadel/internal/user/repository/view/model"
)

type InitCodeEmailData struct {
	templates.TemplateData
	URL string
}

type UrlData struct {
	UserID      string
	Code        string
	PasswordSet bool
}

func SendUserInitCode(mailhtml string, text *iam_model.MessageTextView, user *view_model.NotifyUser, code *es_model.InitUserCode, systemDefaults systemdefaults.SystemDefaults, alg crypto.EncryptionAlgorithm, colors *iam_model.LabelPolicyView, apiDomain string) error {
	codeString, err := crypto.DecryptString(code.Code, alg)
	if err != nil {
		return err
	}
	url, err := templates.ParseTemplateText(systemDefaults.Notifications.Endpoints.InitCode, &UrlData{UserID: user.ID, Code: codeString, PasswordSet: user.PasswordSet})
	if err != nil {
		return err
	}
	var args = mapNotifyUserToArgs(user)
	args["Code"] = codeString

	text.Greeting, err = templates.ParseTemplateText(text.Greeting, args)
	text.Text, err = templates.ParseTemplateText(text.Text, args)
	text.Text = html.UnescapeString(text.Text)

	emailCodeData := &InitCodeEmailData{
		TemplateData: templates.GetTemplateData(apiDomain, url, text, colors),
		URL:          url,
	}
	template, err := templates.GetParsedTemplate(mailhtml, emailCodeData)
	if err != nil {
		return err
	}
	return generateEmail(user, text.Subject, template, systemDefaults.Notifications, true)
}
