package handler

import (
	"github.com/caos/logging"
	"github.com/caos/zitadel/internal/eventstore/v1"

	es_models "github.com/caos/zitadel/internal/eventstore/v1/models"
	"github.com/caos/zitadel/internal/eventstore/v1/query"
	"github.com/caos/zitadel/internal/eventstore/v1/spooler"
	"github.com/caos/zitadel/internal/project/repository/eventsourcing/model"
	proj_view "github.com/caos/zitadel/internal/project/repository/view"
	view_model "github.com/caos/zitadel/internal/project/repository/view/model"
)

const (
	projectRoleTable = "auth.project_roles"
)

type ProjectRole struct {
	handler
	subscription *v1.Subscription
}

func newProjectRole(
	handler handler,
) *ProjectRole {
	h := &ProjectRole{
		handler: handler,
	}

	h.subscribe()

	return h
}

func (k *ProjectRole) subscribe() {
	k.subscription = k.es.Subscribe(k.AggregateTypes()...)
	go func() {
		for event := range k.subscription.Events {
			query.ReduceEvent(k, event)
		}
	}()
}

func (p *ProjectRole) ViewModel() string {
	return projectRoleTable
}

func (_ *ProjectRole) AggregateTypes() []es_models.AggregateType {
	return []es_models.AggregateType{model.ProjectAggregate}
}

func (p *ProjectRole) CurrentSequence() (uint64, error) {
	sequence, err := p.view.GetLatestProjectRoleSequence()
	if err != nil {
		return 0, err
	}
	return sequence.CurrentSequence, nil
}

func (p *ProjectRole) EventQuery() (*es_models.SearchQuery, error) {
	sequence, err := p.view.GetLatestProjectRoleSequence()
	if err != nil {
		return nil, err
	}
	return proj_view.ProjectQuery(sequence.CurrentSequence), nil
}

func (p *ProjectRole) Reduce(event *es_models.Event) (err error) {
	role := new(view_model.ProjectRoleView)
	switch event.Type {
	case model.ProjectRoleAdded:
		err = role.AppendEvent(event)
	case model.ProjectRoleChanged:
		err = role.SetData(event)
		if err != nil {
			return err
		}
		role, err = p.view.ProjectRoleByIDs(event.AggregateID, event.ResourceOwner, role.Key)
		if err != nil {
			return err
		}
		role.ChangeDate = event.CreationDate
		err = role.AppendEvent(event)
	case model.ProjectRoleRemoved:
		err = role.SetData(event)
		if err != nil {
			return err
		}
		return p.view.DeleteProjectRole(event.AggregateID, event.ResourceOwner, role.Key, event)
	case model.ProjectRemoved:
		err := p.view.DeleteProjectRolesByProjectID(event.AggregateID)
		if err == nil {
			return p.view.ProcessedProjectRoleSequence(event)
		}
	default:
		return p.view.ProcessedProjectRoleSequence(event)
	}
	if err != nil {
		return err
	}
	return p.view.PutProjectRole(role, event)
}

func (p *ProjectRole) OnError(event *es_models.Event, err error) error {
	logging.LogWithFields("SPOOL-lso9w", "id", event.AggregateID).WithError(err).Warn("something went wrong in project role handler")
	return spooler.HandleError(event, err, p.view.GetLatestProjectRoleFailedEvent, p.view.ProcessedProjectRoleFailedEvent, p.view.ProcessedProjectRoleSequence, p.errorCountUntilSkip)
}

func (p *ProjectRole) OnSuccess() error {
	return spooler.HandleSuccess(p.view.UpdateProjectRoleSpoolerRunTimestamp)
}
