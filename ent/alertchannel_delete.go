// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/TinyMonitor/ent/alertchannel"
	"github.com/lbrictson/TinyMonitor/ent/predicate"
)

// AlertChannelDelete is the builder for deleting a AlertChannel entity.
type AlertChannelDelete struct {
	config
	hooks    []Hook
	mutation *AlertChannelMutation
}

// Where appends a list predicates to the AlertChannelDelete builder.
func (acd *AlertChannelDelete) Where(ps ...predicate.AlertChannel) *AlertChannelDelete {
	acd.mutation.Where(ps...)
	return acd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acd *AlertChannelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(acd.hooks) == 0 {
		affected, err = acd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlertChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			acd.mutation = mutation
			affected, err = acd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acd.hooks) - 1; i >= 0; i-- {
			if acd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (acd *AlertChannelDelete) ExecX(ctx context.Context) int {
	n, err := acd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acd *AlertChannelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: alertchannel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: alertchannel.FieldID,
			},
		},
	}
	if ps := acd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AlertChannelDeleteOne is the builder for deleting a single AlertChannel entity.
type AlertChannelDeleteOne struct {
	acd *AlertChannelDelete
}

// Exec executes the deletion query.
func (acdo *AlertChannelDeleteOne) Exec(ctx context.Context) error {
	n, err := acdo.acd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{alertchannel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acdo *AlertChannelDeleteOne) ExecX(ctx context.Context) {
	acdo.acd.ExecX(ctx)
}