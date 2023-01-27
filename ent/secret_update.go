// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/TinyMonitor/ent/predicate"
	"github.com/lbrictson/TinyMonitor/ent/secret"
)

// SecretUpdate is the builder for updating Secret entities.
type SecretUpdate struct {
	config
	hooks    []Hook
	mutation *SecretMutation
}

// Where appends a list predicates to the SecretUpdate builder.
func (su *SecretUpdate) Where(ps ...predicate.Secret) *SecretUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SecretUpdate) SetCreatedAt(t time.Time) *SecretUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SecretUpdate) SetNillableCreatedAt(t *time.Time) *SecretUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SecretUpdate) SetUpdatedAt(t time.Time) *SecretUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetCreatedBy sets the "created_by" field.
func (su *SecretUpdate) SetCreatedBy(s string) *SecretUpdate {
	su.mutation.SetCreatedBy(s)
	return su
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (su *SecretUpdate) SetNillableCreatedBy(s *string) *SecretUpdate {
	if s != nil {
		su.SetCreatedBy(*s)
	}
	return su
}

// SetValue sets the "value" field.
func (su *SecretUpdate) SetValue(s string) *SecretUpdate {
	su.mutation.SetValue(s)
	return su
}

// Mutation returns the SecretMutation object of the builder.
func (su *SecretUpdate) Mutation() *SecretMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SecretUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SecretUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SecretUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SecretUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SecretUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := secret.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SecretUpdate) check() error {
	if v, ok := su.mutation.Value(); ok {
		if err := secret.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Secret.value": %w`, err)}
		}
	}
	return nil
}

func (su *SecretUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   secret.Table,
			Columns: secret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: secret.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(secret.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(secret.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.CreatedBy(); ok {
		_spec.SetField(secret.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := su.mutation.Value(); ok {
		_spec.SetField(secret.FieldValue, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{secret.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SecretUpdateOne is the builder for updating a single Secret entity.
type SecretUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SecretMutation
}

// SetCreatedAt sets the "created_at" field.
func (suo *SecretUpdateOne) SetCreatedAt(t time.Time) *SecretUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SecretUpdateOne) SetNillableCreatedAt(t *time.Time) *SecretUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SecretUpdateOne) SetUpdatedAt(t time.Time) *SecretUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetCreatedBy sets the "created_by" field.
func (suo *SecretUpdateOne) SetCreatedBy(s string) *SecretUpdateOne {
	suo.mutation.SetCreatedBy(s)
	return suo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suo *SecretUpdateOne) SetNillableCreatedBy(s *string) *SecretUpdateOne {
	if s != nil {
		suo.SetCreatedBy(*s)
	}
	return suo
}

// SetValue sets the "value" field.
func (suo *SecretUpdateOne) SetValue(s string) *SecretUpdateOne {
	suo.mutation.SetValue(s)
	return suo
}

// Mutation returns the SecretMutation object of the builder.
func (suo *SecretUpdateOne) Mutation() *SecretMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SecretUpdateOne) Select(field string, fields ...string) *SecretUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Secret entity.
func (suo *SecretUpdateOne) Save(ctx context.Context) (*Secret, error) {
	var (
		err  error
		node *Secret
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Secret)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SecretMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SecretUpdateOne) SaveX(ctx context.Context) *Secret {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SecretUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SecretUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SecretUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := secret.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SecretUpdateOne) check() error {
	if v, ok := suo.mutation.Value(); ok {
		if err := secret.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Secret.value": %w`, err)}
		}
	}
	return nil
}

func (suo *SecretUpdateOne) sqlSave(ctx context.Context) (_node *Secret, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   secret.Table,
			Columns: secret.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: secret.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Secret.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, secret.FieldID)
		for _, f := range fields {
			if !secret.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != secret.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(secret.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(secret.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.CreatedBy(); ok {
		_spec.SetField(secret.FieldCreatedBy, field.TypeString, value)
	}
	if value, ok := suo.mutation.Value(); ok {
		_spec.SetField(secret.FieldValue, field.TypeString, value)
	}
	_node = &Secret{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{secret.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}