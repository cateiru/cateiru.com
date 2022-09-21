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
	"github.com/cateiru/cateiru.com/ent/biography"
	"github.com/cateiru/cateiru.com/ent/predicate"
)

// BiographyUpdate is the builder for updating Biography entities.
type BiographyUpdate struct {
	config
	hooks    []Hook
	mutation *BiographyMutation
}

// Where appends a list predicates to the BiographyUpdate builder.
func (bu *BiographyUpdate) Where(ps ...predicate.Biography) *BiographyUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUserID sets the "user_id" field.
func (bu *BiographyUpdate) SetUserID(u uint32) *BiographyUpdate {
	bu.mutation.ResetUserID()
	bu.mutation.SetUserID(u)
	return bu
}

// AddUserID adds u to the "user_id" field.
func (bu *BiographyUpdate) AddUserID(u int32) *BiographyUpdate {
	bu.mutation.AddUserID(u)
	return bu
}

// SetIsPublic sets the "is_public" field.
func (bu *BiographyUpdate) SetIsPublic(b bool) *BiographyUpdate {
	bu.mutation.SetIsPublic(b)
	return bu
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (bu *BiographyUpdate) SetNillableIsPublic(b *bool) *BiographyUpdate {
	if b != nil {
		bu.SetIsPublic(*b)
	}
	return bu
}

// SetLocationID sets the "location_id" field.
func (bu *BiographyUpdate) SetLocationID(u uint32) *BiographyUpdate {
	bu.mutation.ResetLocationID()
	bu.mutation.SetLocationID(u)
	return bu
}

// AddLocationID adds u to the "location_id" field.
func (bu *BiographyUpdate) AddLocationID(u int32) *BiographyUpdate {
	bu.mutation.AddLocationID(u)
	return bu
}

// SetPosition sets the "position" field.
func (bu *BiographyUpdate) SetPosition(s string) *BiographyUpdate {
	bu.mutation.SetPosition(s)
	return bu
}

// SetJoin sets the "join" field.
func (bu *BiographyUpdate) SetJoin(t time.Time) *BiographyUpdate {
	bu.mutation.SetJoin(t)
	return bu
}

// SetLeave sets the "leave" field.
func (bu *BiographyUpdate) SetLeave(t time.Time) *BiographyUpdate {
	bu.mutation.SetLeave(t)
	return bu
}

// SetCreated sets the "created" field.
func (bu *BiographyUpdate) SetCreated(t time.Time) *BiographyUpdate {
	bu.mutation.SetCreated(t)
	return bu
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (bu *BiographyUpdate) SetNillableCreated(t *time.Time) *BiographyUpdate {
	if t != nil {
		bu.SetCreated(*t)
	}
	return bu
}

// SetModified sets the "modified" field.
func (bu *BiographyUpdate) SetModified(t time.Time) *BiographyUpdate {
	bu.mutation.SetModified(t)
	return bu
}

// Mutation returns the BiographyMutation object of the builder.
func (bu *BiographyUpdate) Mutation() *BiographyMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BiographyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	bu.defaults()
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BiographyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BiographyUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BiographyUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BiographyUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BiographyUpdate) defaults() {
	if _, ok := bu.mutation.Modified(); !ok {
		v := biography.UpdateDefaultModified()
		bu.mutation.SetModified(v)
	}
}

func (bu *BiographyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   biography.Table,
			Columns: biography.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: biography.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: biography.FieldUserID,
		})
	}
	if value, ok := bu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: biography.FieldUserID,
		})
	}
	if value, ok := bu.mutation.IsPublic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: biography.FieldIsPublic,
		})
	}
	if value, ok := bu.mutation.LocationID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: biography.FieldLocationID,
		})
	}
	if value, ok := bu.mutation.AddedLocationID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: biography.FieldLocationID,
		})
	}
	if value, ok := bu.mutation.Position(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: biography.FieldPosition,
		})
	}
	if value, ok := bu.mutation.Join(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: biography.FieldJoin,
		})
	}
	if value, ok := bu.mutation.Leave(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: biography.FieldLeave,
		})
	}
	if value, ok := bu.mutation.Created(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: biography.FieldCreated,
		})
	}
	if value, ok := bu.mutation.Modified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: biography.FieldModified,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{biography.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BiographyUpdateOne is the builder for updating a single Biography entity.
type BiographyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BiographyMutation
}

// SetUserID sets the "user_id" field.
func (buo *BiographyUpdateOne) SetUserID(u uint32) *BiographyUpdateOne {
	buo.mutation.ResetUserID()
	buo.mutation.SetUserID(u)
	return buo
}

// AddUserID adds u to the "user_id" field.
func (buo *BiographyUpdateOne) AddUserID(u int32) *BiographyUpdateOne {
	buo.mutation.AddUserID(u)
	return buo
}

// SetIsPublic sets the "is_public" field.
func (buo *BiographyUpdateOne) SetIsPublic(b bool) *BiographyUpdateOne {
	buo.mutation.SetIsPublic(b)
	return buo
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (buo *BiographyUpdateOne) SetNillableIsPublic(b *bool) *BiographyUpdateOne {
	if b != nil {
		buo.SetIsPublic(*b)
	}
	return buo
}

// SetLocationID sets the "location_id" field.
func (buo *BiographyUpdateOne) SetLocationID(u uint32) *BiographyUpdateOne {
	buo.mutation.ResetLocationID()
	buo.mutation.SetLocationID(u)
	return buo
}

// AddLocationID adds u to the "location_id" field.
func (buo *BiographyUpdateOne) AddLocationID(u int32) *BiographyUpdateOne {
	buo.mutation.AddLocationID(u)
	return buo
}

// SetPosition sets the "position" field.
func (buo *BiographyUpdateOne) SetPosition(s string) *BiographyUpdateOne {
	buo.mutation.SetPosition(s)
	return buo
}

// SetJoin sets the "join" field.
func (buo *BiographyUpdateOne) SetJoin(t time.Time) *BiographyUpdateOne {
	buo.mutation.SetJoin(t)
	return buo
}

// SetLeave sets the "leave" field.
func (buo *BiographyUpdateOne) SetLeave(t time.Time) *BiographyUpdateOne {
	buo.mutation.SetLeave(t)
	return buo
}

// SetCreated sets the "created" field.
func (buo *BiographyUpdateOne) SetCreated(t time.Time) *BiographyUpdateOne {
	buo.mutation.SetCreated(t)
	return buo
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (buo *BiographyUpdateOne) SetNillableCreated(t *time.Time) *BiographyUpdateOne {
	if t != nil {
		buo.SetCreated(*t)
	}
	return buo
}

// SetModified sets the "modified" field.
func (buo *BiographyUpdateOne) SetModified(t time.Time) *BiographyUpdateOne {
	buo.mutation.SetModified(t)
	return buo
}

// Mutation returns the BiographyMutation object of the builder.
func (buo *BiographyUpdateOne) Mutation() *BiographyMutation {
	return buo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BiographyUpdateOne) Select(field string, fields ...string) *BiographyUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Biography entity.
func (buo *BiographyUpdateOne) Save(ctx context.Context) (*Biography, error) {
	var (
		err  error
		node *Biography
	)
	buo.defaults()
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BiographyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Biography)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BiographyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BiographyUpdateOne) SaveX(ctx context.Context) *Biography {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BiographyUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BiographyUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BiographyUpdateOne) defaults() {
	if _, ok := buo.mutation.Modified(); !ok {
		v := biography.UpdateDefaultModified()
		buo.mutation.SetModified(v)
	}
}

func (buo *BiographyUpdateOne) sqlSave(ctx context.Context) (_node *Biography, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   biography.Table,
			Columns: biography.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: biography.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Biography.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, biography.FieldID)
		for _, f := range fields {
			if !biography.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != biography.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: biography.FieldUserID,
		})
	}
	if value, ok := buo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: biography.FieldUserID,
		})
	}
	if value, ok := buo.mutation.IsPublic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: biography.FieldIsPublic,
		})
	}
	if value, ok := buo.mutation.LocationID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: biography.FieldLocationID,
		})
	}
	if value, ok := buo.mutation.AddedLocationID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: biography.FieldLocationID,
		})
	}
	if value, ok := buo.mutation.Position(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: biography.FieldPosition,
		})
	}
	if value, ok := buo.mutation.Join(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: biography.FieldJoin,
		})
	}
	if value, ok := buo.mutation.Leave(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: biography.FieldLeave,
		})
	}
	if value, ok := buo.mutation.Created(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: biography.FieldCreated,
		})
	}
	if value, ok := buo.mutation.Modified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: biography.FieldModified,
		})
	}
	_node = &Biography{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{biography.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
