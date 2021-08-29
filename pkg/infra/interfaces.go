package infra

import (
	"context"

	"github.com/m-mizutani/alertchain/pkg/infra/ent"
	"github.com/m-mizutani/alertchain/types"
)

type Clients struct {
	DB DBClient
}

type DBClient interface {
	Close() error

	GetAlert(ctx context.Context, id types.AlertID) (*ent.Alert, error)
	GetAlerts(ctx context.Context) ([]*ent.Alert, error)
	NewAlert(ctx context.Context) (*ent.Alert, error)
	UpdateAlert(ctx context.Context, id types.AlertID, alert *ent.Alert) error
	UpdateAlertStatus(ctx context.Context, id types.AlertID, status types.AlertStatus, ts int64) error
	UpdateAlertSeverity(ctx context.Context, id types.AlertID, status types.Severity, ts int64) error

	AddAttributes(ctx context.Context, id types.AlertID, newAttrs []*ent.Attribute) error
	GetAttribute(ctx context.Context, id int) (*ent.Attribute, error)

	AddAnnotation(ctx context.Context, attr *ent.Attribute, ann []*ent.Annotation) error
	AddReference(ctx context.Context, id types.AlertID, ref *ent.Reference) error
	NewTaskLog(ctx context.Context, id types.AlertID, taskName string, ts, stage int64) (*ent.TaskLog, error)
	UpdateTaskLog(ctx context.Context, task *ent.TaskLog) error
}
