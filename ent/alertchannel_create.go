// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/TinyMonitor/ent/alertchannel"
	"github.com/lbrictson/TinyMonitor/ent/monitor"
)

// AlertChannelCreate is the builder for creating a AlertChannel entity.
type AlertChannelCreate struct {
	config
	mutation *AlertChannelMutation
	hooks    []Hook
}

// SetAlertChannelType sets the "alert_channel_type" field.
func (acc *AlertChannelCreate) SetAlertChannelType(s string) *AlertChannelCreate {
	acc.mutation.SetAlertChannelType(s)
	return acc
}

// SetConfig sets the "config" field.
func (acc *AlertChannelCreate) SetConfig(m map[string]interface{}) *AlertChannelCreate {
	acc.mutation.SetConfig(m)
	return acc
}

// SetCreatedAt sets the "created_at" field.
func (acc *AlertChannelCreate) SetCreatedAt(t time.Time) *AlertChannelCreate {
	acc.mutation.SetCreatedAt(t)
	return acc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acc *AlertChannelCreate) SetNillableCreatedAt(t *time.Time) *AlertChannelCreate {
	if t != nil {
		acc.SetCreatedAt(*t)
	}
	return acc
}

// SetUpdatedAt sets the "updated_at" field.
func (acc *AlertChannelCreate) SetUpdatedAt(t time.Time) *AlertChannelCreate {
	acc.mutation.SetUpdatedAt(t)
	return acc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (acc *AlertChannelCreate) SetNillableUpdatedAt(t *time.Time) *AlertChannelCreate {
	if t != nil {
		acc.SetUpdatedAt(*t)
	}
	return acc
}

// SetID sets the "id" field.
func (acc *AlertChannelCreate) SetID(s string) *AlertChannelCreate {
	acc.mutation.SetID(s)
	return acc
}

// AddMonitorIDs adds the "monitors" edge to the Monitor entity by IDs.
func (acc *AlertChannelCreate) AddMonitorIDs(ids ...string) *AlertChannelCreate {
	acc.mutation.AddMonitorIDs(ids...)
	return acc
}

// AddMonitors adds the "monitors" edges to the Monitor entity.
func (acc *AlertChannelCreate) AddMonitors(m ...*Monitor) *AlertChannelCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return acc.AddMonitorIDs(ids...)
}

// Mutation returns the AlertChannelMutation object of the builder.
func (acc *AlertChannelCreate) Mutation() *AlertChannelMutation {
	return acc.mutation
}

// Save creates the AlertChannel in the database.
func (acc *AlertChannelCreate) Save(ctx context.Context) (*AlertChannel, error) {
	var (
		err  error
		node *AlertChannel
	)
	acc.defaults()
	if len(acc.hooks) == 0 {
		if err = acc.check(); err != nil {
			return nil, err
		}
		node, err = acc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlertChannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acc.check(); err != nil {
				return nil, err
			}
			acc.mutation = mutation
			if node, err = acc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(acc.hooks) - 1; i >= 0; i-- {
			if acc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, acc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AlertChannel)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AlertChannelMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AlertChannelCreate) SaveX(ctx context.Context) *AlertChannel {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AlertChannelCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AlertChannelCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *AlertChannelCreate) defaults() {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		v := alertchannel.DefaultCreatedAt()
		acc.mutation.SetCreatedAt(v)
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		v := alertchannel.DefaultUpdatedAt()
		acc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *AlertChannelCreate) check() error {
	if _, ok := acc.mutation.AlertChannelType(); !ok {
		return &ValidationError{Name: "alert_channel_type", err: errors.New(`ent: missing required field "AlertChannel.alert_channel_type"`)}
	}
	if _, ok := acc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`ent: missing required field "AlertChannel.config"`)}
	}
	if _, ok := acc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AlertChannel.created_at"`)}
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AlertChannel.updated_at"`)}
	}
	if v, ok := acc.mutation.ID(); ok {
		if err := alertchannel.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "AlertChannel.id": %w`, err)}
		}
	}
	return nil
}

func (acc *AlertChannelCreate) sqlSave(ctx context.Context) (*AlertChannel, error) {
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AlertChannel.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (acc *AlertChannelCreate) createSpec() (*AlertChannel, *sqlgraph.CreateSpec) {
	var (
		_node = &AlertChannel{config: acc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: alertchannel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: alertchannel.FieldID,
			},
		}
	)
	if id, ok := acc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := acc.mutation.AlertChannelType(); ok {
		_spec.SetField(alertchannel.FieldAlertChannelType, field.TypeString, value)
		_node.AlertChannelType = value
	}
	if value, ok := acc.mutation.Config(); ok {
		_spec.SetField(alertchannel.FieldConfig, field.TypeJSON, value)
		_node.Config = value
	}
	if value, ok := acc.mutation.CreatedAt(); ok {
		_spec.SetField(alertchannel.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := acc.mutation.UpdatedAt(); ok {
		_spec.SetField(alertchannel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := acc.mutation.MonitorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   alertchannel.MonitorsTable,
			Columns: alertchannel.MonitorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: monitor.FieldID,
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

// AlertChannelCreateBulk is the builder for creating many AlertChannel entities in bulk.
type AlertChannelCreateBulk struct {
	config
	builders []*AlertChannelCreate
}

// Save creates the AlertChannel entities in the database.
func (accb *AlertChannelCreateBulk) Save(ctx context.Context) ([]*AlertChannel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AlertChannel, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AlertChannelMutation)
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
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AlertChannelCreateBulk) SaveX(ctx context.Context) []*AlertChannel {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AlertChannelCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AlertChannelCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}
