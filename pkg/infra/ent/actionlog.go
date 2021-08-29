// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/actionlog"
	"github.com/m-mizutani/alertchain/types"
)

// ActionLog is the model entity for the ActionLog schema.
type ActionLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// StartedAt holds the value of the "started_at" field.
	StartedAt int64 `json:"started_at,omitempty"`
	// ExitedAt holds the value of the "exited_at" field.
	ExitedAt int64 `json:"exited_at,omitempty"`
	// Log holds the value of the "log" field.
	Log string `json:"log,omitempty"`
	// Errmsg holds the value of the "errmsg" field.
	Errmsg string `json:"errmsg,omitempty"`
	// ErrValues holds the value of the "err_values" field.
	ErrValues []string `json:"err_values,omitempty"`
	// StackTrace holds the value of the "stack_trace" field.
	StackTrace []string `json:"stack_trace,omitempty"`
	// Status holds the value of the "status" field.
	Status types.TaskStatus `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActionLogQuery when eager-loading is set.
	Edges ActionLogEdges `json:"edges"`
}

// ActionLogEdges holds the relations/edges for other nodes in the graph.
type ActionLogEdges struct {
	// Argument holds the value of the argument edge.
	Argument []*Attribute `json:"argument,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ArgumentOrErr returns the Argument value or an error if the edge
// was not loaded in eager-loading.
func (e ActionLogEdges) ArgumentOrErr() ([]*Attribute, error) {
	if e.loadedTypes[0] {
		return e.Argument, nil
	}
	return nil, &NotLoadedError{edge: "argument"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ActionLog) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case actionlog.FieldErrValues, actionlog.FieldStackTrace:
			values[i] = new([]byte)
		case actionlog.FieldID, actionlog.FieldStartedAt, actionlog.FieldExitedAt:
			values[i] = new(sql.NullInt64)
		case actionlog.FieldName, actionlog.FieldLog, actionlog.FieldErrmsg, actionlog.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ActionLog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ActionLog fields.
func (al *ActionLog) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case actionlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			al.ID = int(value.Int64)
		case actionlog.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				al.Name = value.String
			}
		case actionlog.FieldStartedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field started_at", values[i])
			} else if value.Valid {
				al.StartedAt = value.Int64
			}
		case actionlog.FieldExitedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field exited_at", values[i])
			} else if value.Valid {
				al.ExitedAt = value.Int64
			}
		case actionlog.FieldLog:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field log", values[i])
			} else if value.Valid {
				al.Log = value.String
			}
		case actionlog.FieldErrmsg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field errmsg", values[i])
			} else if value.Valid {
				al.Errmsg = value.String
			}
		case actionlog.FieldErrValues:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field err_values", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &al.ErrValues); err != nil {
					return fmt.Errorf("unmarshal field err_values: %w", err)
				}
			}
		case actionlog.FieldStackTrace:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field stack_trace", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &al.StackTrace); err != nil {
					return fmt.Errorf("unmarshal field stack_trace: %w", err)
				}
			}
		case actionlog.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				al.Status = types.TaskStatus(value.String)
			}
		}
	}
	return nil
}

// QueryArgument queries the "argument" edge of the ActionLog entity.
func (al *ActionLog) QueryArgument() *AttributeQuery {
	return (&ActionLogClient{config: al.config}).QueryArgument(al)
}

// Update returns a builder for updating this ActionLog.
// Note that you need to call ActionLog.Unwrap() before calling this method if this ActionLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (al *ActionLog) Update() *ActionLogUpdateOne {
	return (&ActionLogClient{config: al.config}).UpdateOne(al)
}

// Unwrap unwraps the ActionLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (al *ActionLog) Unwrap() *ActionLog {
	tx, ok := al.config.driver.(*txDriver)
	if !ok {
		panic("ent: ActionLog is not a transactional entity")
	}
	al.config.driver = tx.drv
	return al
}

// String implements the fmt.Stringer.
func (al *ActionLog) String() string {
	var builder strings.Builder
	builder.WriteString("ActionLog(")
	builder.WriteString(fmt.Sprintf("id=%v", al.ID))
	builder.WriteString(", name=")
	builder.WriteString(al.Name)
	builder.WriteString(", started_at=")
	builder.WriteString(fmt.Sprintf("%v", al.StartedAt))
	builder.WriteString(", exited_at=")
	builder.WriteString(fmt.Sprintf("%v", al.ExitedAt))
	builder.WriteString(", log=")
	builder.WriteString(al.Log)
	builder.WriteString(", errmsg=")
	builder.WriteString(al.Errmsg)
	builder.WriteString(", err_values=")
	builder.WriteString(fmt.Sprintf("%v", al.ErrValues))
	builder.WriteString(", stack_trace=")
	builder.WriteString(fmt.Sprintf("%v", al.StackTrace))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", al.Status))
	builder.WriteByte(')')
	return builder.String()
}

// ActionLogs is a parsable slice of ActionLog.
type ActionLogs []*ActionLog

func (al ActionLogs) config(cfg config) {
	for _i := range al {
		al[_i].config = cfg
	}
}
