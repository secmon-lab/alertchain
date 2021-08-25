// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/m-mizutani/alertchain/pkg/infra/ent/alert"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/tasklog"
	"github.com/m-mizutani/alertchain/pkg/infra/schema"
	"github.com/m-mizutani/alertchain/types"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	alertFields := schema.Alert{}.Fields()
	_ = alertFields
	// alertDescStatus is the schema descriptor for status field.
	alertDescStatus := alertFields[4].Descriptor()
	// alert.DefaultStatus holds the default value on creation for the status field.
	alert.DefaultStatus = types.AlertStatus(alertDescStatus.Default.(string))
	tasklogFields := schema.TaskLog{}.Fields()
	_ = tasklogFields
	// tasklogDescOptional is the schema descriptor for optional field.
	tasklogDescOptional := tasklogFields[1].Descriptor()
	// tasklog.DefaultOptional holds the default value on creation for the optional field.
	tasklog.DefaultOptional = tasklogDescOptional.Default.(bool)
	// tasklogDescStatus is the schema descriptor for status field.
	tasklogDescStatus := tasklogFields[7].Descriptor()
	// tasklog.DefaultStatus holds the default value on creation for the status field.
	tasklog.DefaultStatus = types.TaskStatus(tasklogDescStatus.Default.(string))
}
