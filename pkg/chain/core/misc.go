package core

import (
	"github.com/m-mizutani/alertchain/pkg/domain/interfaces"
	"github.com/m-mizutani/alertchain/pkg/domain/model"
)

type dummyScenarioLogger struct{}

func (*dummyScenarioLogger) NewAlertLogger(alert *model.Alert) interfaces.AlertLogger {
	return &dummyAlertLogger{}
}

var _ interfaces.ScenarioLogger = &dummyScenarioLogger{}

func (x *dummyScenarioLogger) LogError(err error) {}
func (x *dummyScenarioLogger) Flush() error       { return nil }

type dummyAlertLogger struct{}

// NewActionLogger implements interfaces.AlertLogger.
func (*dummyAlertLogger) NewActionLogger() interfaces.ActionLogger {
	return &dummyActionLogger{}
}

var _ interfaces.AlertLogger = &dummyAlertLogger{}

type dummyActionLogger struct{}

func (*dummyActionLogger) LogExit(log []model.Next)  {}
func (*dummyActionLogger) LogInit(log []model.Next)  {}
func (*dummyActionLogger) LogRun(log []model.Action) {}

var _ interfaces.ActionLogger = &dummyActionLogger{}
