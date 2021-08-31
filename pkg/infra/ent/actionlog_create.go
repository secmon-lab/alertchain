// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/actionlog"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/attribute"
	"github.com/m-mizutani/alertchain/pkg/infra/ent/execlog"
)

// ActionLogCreate is the builder for creating a ActionLog entity.
type ActionLogCreate struct {
	config
	mutation *ActionLogMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (alc *ActionLogCreate) SetName(s string) *ActionLogCreate {
	alc.mutation.SetName(s)
	return alc
}

// AddArgumentIDs adds the "argument" edge to the Attribute entity by IDs.
func (alc *ActionLogCreate) AddArgumentIDs(ids ...int) *ActionLogCreate {
	alc.mutation.AddArgumentIDs(ids...)
	return alc
}

// AddArgument adds the "argument" edges to the Attribute entity.
func (alc *ActionLogCreate) AddArgument(a ...*Attribute) *ActionLogCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return alc.AddArgumentIDs(ids...)
}

// AddExecLogIDs adds the "exec_logs" edge to the ExecLog entity by IDs.
func (alc *ActionLogCreate) AddExecLogIDs(ids ...int) *ActionLogCreate {
	alc.mutation.AddExecLogIDs(ids...)
	return alc
}

// AddExecLogs adds the "exec_logs" edges to the ExecLog entity.
func (alc *ActionLogCreate) AddExecLogs(e ...*ExecLog) *ActionLogCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return alc.AddExecLogIDs(ids...)
}

// Mutation returns the ActionLogMutation object of the builder.
func (alc *ActionLogCreate) Mutation() *ActionLogMutation {
	return alc.mutation
}

// Save creates the ActionLog in the database.
func (alc *ActionLogCreate) Save(ctx context.Context) (*ActionLog, error) {
	var (
		err  error
		node *ActionLog
	)
	if len(alc.hooks) == 0 {
		if err = alc.check(); err != nil {
			return nil, err
		}
		node, err = alc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = alc.check(); err != nil {
				return nil, err
			}
			alc.mutation = mutation
			if node, err = alc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(alc.hooks) - 1; i >= 0; i-- {
			if alc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = alc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, alc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (alc *ActionLogCreate) SaveX(ctx context.Context) *ActionLog {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *ActionLogCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *ActionLogCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *ActionLogCreate) check() error {
	if _, ok := alc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (alc *ActionLogCreate) sqlSave(ctx context.Context) (*ActionLog, error) {
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (alc *ActionLogCreate) createSpec() (*ActionLog, *sqlgraph.CreateSpec) {
	var (
		_node = &ActionLog{config: alc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: actionlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: actionlog.FieldID,
			},
		}
	)
	if value, ok := alc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: actionlog.FieldName,
		})
		_node.Name = value
	}
	if nodes := alc.mutation.ArgumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   actionlog.ArgumentTable,
			Columns: []string{actionlog.ArgumentColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := alc.mutation.ExecLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   actionlog.ExecLogsTable,
			Columns: []string{actionlog.ExecLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: execlog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActionLogCreateBulk is the builder for creating many ActionLog entities in bulk.
type ActionLogCreateBulk struct {
	config
	builders []*ActionLogCreate
}

// Save creates the ActionLog entities in the database.
func (alcb *ActionLogCreateBulk) Save(ctx context.Context) ([]*ActionLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*ActionLog, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActionLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *ActionLogCreateBulk) SaveX(ctx context.Context) []*ActionLog {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *ActionLogCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *ActionLogCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}