// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/alert"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/attribute"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/predicate"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/reference"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/tasklog"
	"github.com/m-mizutani/alertchain/types"
)

// AlertUpdate is the builder for updating Alert entities.
type AlertUpdate struct {
	config
	hooks    []Hook
	mutation *AlertMutation
}

// Where appends a list predicates to the AlertUpdate builder.
func (au *AlertUpdate) Where(ps ...predicate.Alert) *AlertUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetTitle sets the "title" field.
func (au *AlertUpdate) SetTitle(s string) *AlertUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (au *AlertUpdate) SetNillableTitle(s *string) *AlertUpdate {
	if s != nil {
		au.SetTitle(*s)
	}
	return au
}

// ClearTitle clears the value of the "title" field.
func (au *AlertUpdate) ClearTitle() *AlertUpdate {
	au.mutation.ClearTitle()
	return au
}

// SetDescription sets the "description" field.
func (au *AlertUpdate) SetDescription(s string) *AlertUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (au *AlertUpdate) SetNillableDescription(s *string) *AlertUpdate {
	if s != nil {
		au.SetDescription(*s)
	}
	return au
}

// ClearDescription clears the value of the "description" field.
func (au *AlertUpdate) ClearDescription() *AlertUpdate {
	au.mutation.ClearDescription()
	return au
}

// SetDetector sets the "detector" field.
func (au *AlertUpdate) SetDetector(s string) *AlertUpdate {
	au.mutation.SetDetector(s)
	return au
}

// SetNillableDetector sets the "detector" field if the given value is not nil.
func (au *AlertUpdate) SetNillableDetector(s *string) *AlertUpdate {
	if s != nil {
		au.SetDetector(*s)
	}
	return au
}

// ClearDetector clears the value of the "detector" field.
func (au *AlertUpdate) ClearDetector() *AlertUpdate {
	au.mutation.ClearDetector()
	return au
}

// SetStatus sets the "status" field.
func (au *AlertUpdate) SetStatus(ts types.AlertStatus) *AlertUpdate {
	au.mutation.SetStatus(ts)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *AlertUpdate) SetNillableStatus(ts *types.AlertStatus) *AlertUpdate {
	if ts != nil {
		au.SetStatus(*ts)
	}
	return au
}

// SetSeverity sets the "severity" field.
func (au *AlertUpdate) SetSeverity(t types.Severity) *AlertUpdate {
	au.mutation.SetSeverity(t)
	return au
}

// SetNillableSeverity sets the "severity" field if the given value is not nil.
func (au *AlertUpdate) SetNillableSeverity(t *types.Severity) *AlertUpdate {
	if t != nil {
		au.SetSeverity(*t)
	}
	return au
}

// ClearSeverity clears the value of the "severity" field.
func (au *AlertUpdate) ClearSeverity() *AlertUpdate {
	au.mutation.ClearSeverity()
	return au
}

// SetDetectedAt sets the "detected_at" field.
func (au *AlertUpdate) SetDetectedAt(i int64) *AlertUpdate {
	au.mutation.ResetDetectedAt()
	au.mutation.SetDetectedAt(i)
	return au
}

// SetNillableDetectedAt sets the "detected_at" field if the given value is not nil.
func (au *AlertUpdate) SetNillableDetectedAt(i *int64) *AlertUpdate {
	if i != nil {
		au.SetDetectedAt(*i)
	}
	return au
}

// AddDetectedAt adds i to the "detected_at" field.
func (au *AlertUpdate) AddDetectedAt(i int64) *AlertUpdate {
	au.mutation.AddDetectedAt(i)
	return au
}

// ClearDetectedAt clears the value of the "detected_at" field.
func (au *AlertUpdate) ClearDetectedAt() *AlertUpdate {
	au.mutation.ClearDetectedAt()
	return au
}

// SetClosedAt sets the "closed_at" field.
func (au *AlertUpdate) SetClosedAt(i int64) *AlertUpdate {
	au.mutation.ResetClosedAt()
	au.mutation.SetClosedAt(i)
	return au
}

// SetNillableClosedAt sets the "closed_at" field if the given value is not nil.
func (au *AlertUpdate) SetNillableClosedAt(i *int64) *AlertUpdate {
	if i != nil {
		au.SetClosedAt(*i)
	}
	return au
}

// AddClosedAt adds i to the "closed_at" field.
func (au *AlertUpdate) AddClosedAt(i int64) *AlertUpdate {
	au.mutation.AddClosedAt(i)
	return au
}

// ClearClosedAt clears the value of the "closed_at" field.
func (au *AlertUpdate) ClearClosedAt() *AlertUpdate {
	au.mutation.ClearClosedAt()
	return au
}

// AddAttributeIDs adds the "attributes" edge to the Attribute entity by IDs.
func (au *AlertUpdate) AddAttributeIDs(ids ...int) *AlertUpdate {
	au.mutation.AddAttributeIDs(ids...)
	return au
}

// AddAttributes adds the "attributes" edges to the Attribute entity.
func (au *AlertUpdate) AddAttributes(a ...*Attribute) *AlertUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAttributeIDs(ids...)
}

// AddReferenceIDs adds the "references" edge to the Reference entity by IDs.
func (au *AlertUpdate) AddReferenceIDs(ids ...int) *AlertUpdate {
	au.mutation.AddReferenceIDs(ids...)
	return au
}

// AddReferences adds the "references" edges to the Reference entity.
func (au *AlertUpdate) AddReferences(r ...*Reference) *AlertUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.AddReferenceIDs(ids...)
}

// AddTaskLogIDs adds the "task_logs" edge to the TaskLog entity by IDs.
func (au *AlertUpdate) AddTaskLogIDs(ids ...int) *AlertUpdate {
	au.mutation.AddTaskLogIDs(ids...)
	return au
}

// AddTaskLogs adds the "task_logs" edges to the TaskLog entity.
func (au *AlertUpdate) AddTaskLogs(t ...*TaskLog) *AlertUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTaskLogIDs(ids...)
}

// Mutation returns the AlertMutation object of the builder.
func (au *AlertUpdate) Mutation() *AlertMutation {
	return au.mutation
}

// ClearAttributes clears all "attributes" edges to the Attribute entity.
func (au *AlertUpdate) ClearAttributes() *AlertUpdate {
	au.mutation.ClearAttributes()
	return au
}

// RemoveAttributeIDs removes the "attributes" edge to Attribute entities by IDs.
func (au *AlertUpdate) RemoveAttributeIDs(ids ...int) *AlertUpdate {
	au.mutation.RemoveAttributeIDs(ids...)
	return au
}

// RemoveAttributes removes "attributes" edges to Attribute entities.
func (au *AlertUpdate) RemoveAttributes(a ...*Attribute) *AlertUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAttributeIDs(ids...)
}

// ClearReferences clears all "references" edges to the Reference entity.
func (au *AlertUpdate) ClearReferences() *AlertUpdate {
	au.mutation.ClearReferences()
	return au
}

// RemoveReferenceIDs removes the "references" edge to Reference entities by IDs.
func (au *AlertUpdate) RemoveReferenceIDs(ids ...int) *AlertUpdate {
	au.mutation.RemoveReferenceIDs(ids...)
	return au
}

// RemoveReferences removes "references" edges to Reference entities.
func (au *AlertUpdate) RemoveReferences(r ...*Reference) *AlertUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.RemoveReferenceIDs(ids...)
}

// ClearTaskLogs clears all "task_logs" edges to the TaskLog entity.
func (au *AlertUpdate) ClearTaskLogs() *AlertUpdate {
	au.mutation.ClearTaskLogs()
	return au
}

// RemoveTaskLogIDs removes the "task_logs" edge to TaskLog entities by IDs.
func (au *AlertUpdate) RemoveTaskLogIDs(ids ...int) *AlertUpdate {
	au.mutation.RemoveTaskLogIDs(ids...)
	return au
}

// RemoveTaskLogs removes "task_logs" edges to TaskLog entities.
func (au *AlertUpdate) RemoveTaskLogs(t ...*TaskLog) *AlertUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTaskLogIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AlertUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlertMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AlertUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AlertUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AlertUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AlertUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   alert.Table,
			Columns: alert.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: alert.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldTitle,
		})
	}
	if au.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: alert.FieldTitle,
		})
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldDescription,
		})
	}
	if au.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: alert.FieldDescription,
		})
	}
	if value, ok := au.mutation.Detector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldDetector,
		})
	}
	if au.mutation.DetectorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: alert.FieldDetector,
		})
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldStatus,
		})
	}
	if value, ok := au.mutation.Severity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldSeverity,
		})
	}
	if au.mutation.SeverityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: alert.FieldSeverity,
		})
	}
	if value, ok := au.mutation.DetectedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alert.FieldDetectedAt,
		})
	}
	if value, ok := au.mutation.AddedDetectedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alert.FieldDetectedAt,
		})
	}
	if au.mutation.DetectedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: alert.FieldDetectedAt,
		})
	}
	if value, ok := au.mutation.ClosedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alert.FieldClosedAt,
		})
	}
	if value, ok := au.mutation.AddedClosedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alert.FieldClosedAt,
		})
	}
	if au.mutation.ClosedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: alert.FieldClosedAt,
		})
	}
	if au.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.AttributesTable,
			Columns: []string{alert.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAttributesIDs(); len(nodes) > 0 && !au.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.AttributesTable,
			Columns: []string{alert.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AttributesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.AttributesTable,
			Columns: []string{alert.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.ReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.ReferencesTable,
			Columns: []string{alert.ReferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reference.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedReferencesIDs(); len(nodes) > 0 && !au.mutation.ReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.ReferencesTable,
			Columns: []string{alert.ReferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reference.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ReferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.ReferencesTable,
			Columns: []string{alert.ReferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reference.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TaskLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.TaskLogsTable,
			Columns: []string{alert.TaskLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasklog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTaskLogsIDs(); len(nodes) > 0 && !au.mutation.TaskLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.TaskLogsTable,
			Columns: []string{alert.TaskLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasklog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TaskLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.TaskLogsTable,
			Columns: []string{alert.TaskLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasklog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{alert.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AlertUpdateOne is the builder for updating a single Alert entity.
type AlertUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AlertMutation
}

// SetTitle sets the "title" field.
func (auo *AlertUpdateOne) SetTitle(s string) *AlertUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (auo *AlertUpdateOne) SetNillableTitle(s *string) *AlertUpdateOne {
	if s != nil {
		auo.SetTitle(*s)
	}
	return auo
}

// ClearTitle clears the value of the "title" field.
func (auo *AlertUpdateOne) ClearTitle() *AlertUpdateOne {
	auo.mutation.ClearTitle()
	return auo
}

// SetDescription sets the "description" field.
func (auo *AlertUpdateOne) SetDescription(s string) *AlertUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (auo *AlertUpdateOne) SetNillableDescription(s *string) *AlertUpdateOne {
	if s != nil {
		auo.SetDescription(*s)
	}
	return auo
}

// ClearDescription clears the value of the "description" field.
func (auo *AlertUpdateOne) ClearDescription() *AlertUpdateOne {
	auo.mutation.ClearDescription()
	return auo
}

// SetDetector sets the "detector" field.
func (auo *AlertUpdateOne) SetDetector(s string) *AlertUpdateOne {
	auo.mutation.SetDetector(s)
	return auo
}

// SetNillableDetector sets the "detector" field if the given value is not nil.
func (auo *AlertUpdateOne) SetNillableDetector(s *string) *AlertUpdateOne {
	if s != nil {
		auo.SetDetector(*s)
	}
	return auo
}

// ClearDetector clears the value of the "detector" field.
func (auo *AlertUpdateOne) ClearDetector() *AlertUpdateOne {
	auo.mutation.ClearDetector()
	return auo
}

// SetStatus sets the "status" field.
func (auo *AlertUpdateOne) SetStatus(ts types.AlertStatus) *AlertUpdateOne {
	auo.mutation.SetStatus(ts)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *AlertUpdateOne) SetNillableStatus(ts *types.AlertStatus) *AlertUpdateOne {
	if ts != nil {
		auo.SetStatus(*ts)
	}
	return auo
}

// SetSeverity sets the "severity" field.
func (auo *AlertUpdateOne) SetSeverity(t types.Severity) *AlertUpdateOne {
	auo.mutation.SetSeverity(t)
	return auo
}

// SetNillableSeverity sets the "severity" field if the given value is not nil.
func (auo *AlertUpdateOne) SetNillableSeverity(t *types.Severity) *AlertUpdateOne {
	if t != nil {
		auo.SetSeverity(*t)
	}
	return auo
}

// ClearSeverity clears the value of the "severity" field.
func (auo *AlertUpdateOne) ClearSeverity() *AlertUpdateOne {
	auo.mutation.ClearSeverity()
	return auo
}

// SetDetectedAt sets the "detected_at" field.
func (auo *AlertUpdateOne) SetDetectedAt(i int64) *AlertUpdateOne {
	auo.mutation.ResetDetectedAt()
	auo.mutation.SetDetectedAt(i)
	return auo
}

// SetNillableDetectedAt sets the "detected_at" field if the given value is not nil.
func (auo *AlertUpdateOne) SetNillableDetectedAt(i *int64) *AlertUpdateOne {
	if i != nil {
		auo.SetDetectedAt(*i)
	}
	return auo
}

// AddDetectedAt adds i to the "detected_at" field.
func (auo *AlertUpdateOne) AddDetectedAt(i int64) *AlertUpdateOne {
	auo.mutation.AddDetectedAt(i)
	return auo
}

// ClearDetectedAt clears the value of the "detected_at" field.
func (auo *AlertUpdateOne) ClearDetectedAt() *AlertUpdateOne {
	auo.mutation.ClearDetectedAt()
	return auo
}

// SetClosedAt sets the "closed_at" field.
func (auo *AlertUpdateOne) SetClosedAt(i int64) *AlertUpdateOne {
	auo.mutation.ResetClosedAt()
	auo.mutation.SetClosedAt(i)
	return auo
}

// SetNillableClosedAt sets the "closed_at" field if the given value is not nil.
func (auo *AlertUpdateOne) SetNillableClosedAt(i *int64) *AlertUpdateOne {
	if i != nil {
		auo.SetClosedAt(*i)
	}
	return auo
}

// AddClosedAt adds i to the "closed_at" field.
func (auo *AlertUpdateOne) AddClosedAt(i int64) *AlertUpdateOne {
	auo.mutation.AddClosedAt(i)
	return auo
}

// ClearClosedAt clears the value of the "closed_at" field.
func (auo *AlertUpdateOne) ClearClosedAt() *AlertUpdateOne {
	auo.mutation.ClearClosedAt()
	return auo
}

// AddAttributeIDs adds the "attributes" edge to the Attribute entity by IDs.
func (auo *AlertUpdateOne) AddAttributeIDs(ids ...int) *AlertUpdateOne {
	auo.mutation.AddAttributeIDs(ids...)
	return auo
}

// AddAttributes adds the "attributes" edges to the Attribute entity.
func (auo *AlertUpdateOne) AddAttributes(a ...*Attribute) *AlertUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAttributeIDs(ids...)
}

// AddReferenceIDs adds the "references" edge to the Reference entity by IDs.
func (auo *AlertUpdateOne) AddReferenceIDs(ids ...int) *AlertUpdateOne {
	auo.mutation.AddReferenceIDs(ids...)
	return auo
}

// AddReferences adds the "references" edges to the Reference entity.
func (auo *AlertUpdateOne) AddReferences(r ...*Reference) *AlertUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.AddReferenceIDs(ids...)
}

// AddTaskLogIDs adds the "task_logs" edge to the TaskLog entity by IDs.
func (auo *AlertUpdateOne) AddTaskLogIDs(ids ...int) *AlertUpdateOne {
	auo.mutation.AddTaskLogIDs(ids...)
	return auo
}

// AddTaskLogs adds the "task_logs" edges to the TaskLog entity.
func (auo *AlertUpdateOne) AddTaskLogs(t ...*TaskLog) *AlertUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTaskLogIDs(ids...)
}

// Mutation returns the AlertMutation object of the builder.
func (auo *AlertUpdateOne) Mutation() *AlertMutation {
	return auo.mutation
}

// ClearAttributes clears all "attributes" edges to the Attribute entity.
func (auo *AlertUpdateOne) ClearAttributes() *AlertUpdateOne {
	auo.mutation.ClearAttributes()
	return auo
}

// RemoveAttributeIDs removes the "attributes" edge to Attribute entities by IDs.
func (auo *AlertUpdateOne) RemoveAttributeIDs(ids ...int) *AlertUpdateOne {
	auo.mutation.RemoveAttributeIDs(ids...)
	return auo
}

// RemoveAttributes removes "attributes" edges to Attribute entities.
func (auo *AlertUpdateOne) RemoveAttributes(a ...*Attribute) *AlertUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAttributeIDs(ids...)
}

// ClearReferences clears all "references" edges to the Reference entity.
func (auo *AlertUpdateOne) ClearReferences() *AlertUpdateOne {
	auo.mutation.ClearReferences()
	return auo
}

// RemoveReferenceIDs removes the "references" edge to Reference entities by IDs.
func (auo *AlertUpdateOne) RemoveReferenceIDs(ids ...int) *AlertUpdateOne {
	auo.mutation.RemoveReferenceIDs(ids...)
	return auo
}

// RemoveReferences removes "references" edges to Reference entities.
func (auo *AlertUpdateOne) RemoveReferences(r ...*Reference) *AlertUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.RemoveReferenceIDs(ids...)
}

// ClearTaskLogs clears all "task_logs" edges to the TaskLog entity.
func (auo *AlertUpdateOne) ClearTaskLogs() *AlertUpdateOne {
	auo.mutation.ClearTaskLogs()
	return auo
}

// RemoveTaskLogIDs removes the "task_logs" edge to TaskLog entities by IDs.
func (auo *AlertUpdateOne) RemoveTaskLogIDs(ids ...int) *AlertUpdateOne {
	auo.mutation.RemoveTaskLogIDs(ids...)
	return auo
}

// RemoveTaskLogs removes "task_logs" edges to TaskLog entities.
func (auo *AlertUpdateOne) RemoveTaskLogs(t ...*TaskLog) *AlertUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTaskLogIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AlertUpdateOne) Select(field string, fields ...string) *AlertUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Alert entity.
func (auo *AlertUpdateOne) Save(ctx context.Context) (*Alert, error) {
	var (
		err  error
		node *Alert
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlertMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AlertUpdateOne) SaveX(ctx context.Context) *Alert {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AlertUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AlertUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AlertUpdateOne) sqlSave(ctx context.Context) (_node *Alert, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   alert.Table,
			Columns: alert.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: alert.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Alert.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, alert.FieldID)
		for _, f := range fields {
			if !alert.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != alert.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldTitle,
		})
	}
	if auo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: alert.FieldTitle,
		})
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldDescription,
		})
	}
	if auo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: alert.FieldDescription,
		})
	}
	if value, ok := auo.mutation.Detector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldDetector,
		})
	}
	if auo.mutation.DetectorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: alert.FieldDetector,
		})
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldStatus,
		})
	}
	if value, ok := auo.mutation.Severity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: alert.FieldSeverity,
		})
	}
	if auo.mutation.SeverityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: alert.FieldSeverity,
		})
	}
	if value, ok := auo.mutation.DetectedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alert.FieldDetectedAt,
		})
	}
	if value, ok := auo.mutation.AddedDetectedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alert.FieldDetectedAt,
		})
	}
	if auo.mutation.DetectedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: alert.FieldDetectedAt,
		})
	}
	if value, ok := auo.mutation.ClosedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alert.FieldClosedAt,
		})
	}
	if value, ok := auo.mutation.AddedClosedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: alert.FieldClosedAt,
		})
	}
	if auo.mutation.ClosedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: alert.FieldClosedAt,
		})
	}
	if auo.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.AttributesTable,
			Columns: []string{alert.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAttributesIDs(); len(nodes) > 0 && !auo.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.AttributesTable,
			Columns: []string{alert.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AttributesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.AttributesTable,
			Columns: []string{alert.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attribute.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.ReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.ReferencesTable,
			Columns: []string{alert.ReferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reference.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedReferencesIDs(); len(nodes) > 0 && !auo.mutation.ReferencesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.ReferencesTable,
			Columns: []string{alert.ReferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reference.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ReferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.ReferencesTable,
			Columns: []string{alert.ReferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: reference.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TaskLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.TaskLogsTable,
			Columns: []string{alert.TaskLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasklog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTaskLogsIDs(); len(nodes) > 0 && !auo.mutation.TaskLogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.TaskLogsTable,
			Columns: []string{alert.TaskLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasklog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TaskLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   alert.TaskLogsTable,
			Columns: []string{alert.TaskLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasklog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Alert{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{alert.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
