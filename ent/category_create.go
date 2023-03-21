// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cateiru/cateiru.com/ent/category"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CategoryCreate) SetName(s string) *CategoryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNameJa sets the "name_ja" field.
func (cc *CategoryCreate) SetNameJa(s string) *CategoryCreate {
	cc.mutation.SetNameJa(s)
	return cc
}

// SetEmoji sets the "emoji" field.
func (cc *CategoryCreate) SetEmoji(s string) *CategoryCreate {
	cc.mutation.SetEmoji(s)
	return cc
}

// SetCreated sets the "created" field.
func (cc *CategoryCreate) SetCreated(t time.Time) *CategoryCreate {
	cc.mutation.SetCreated(t)
	return cc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableCreated(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetCreated(*t)
	}
	return cc
}

// SetModified sets the "modified" field.
func (cc *CategoryCreate) SetModified(t time.Time) *CategoryCreate {
	cc.mutation.SetModified(t)
	return cc
}

// SetNillableModified sets the "modified" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableModified(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetModified(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CategoryCreate) SetID(u uint32) *CategoryCreate {
	cc.mutation.SetID(u)
	return cc
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	cc.defaults()
	return withHooks[*Category, CategoryMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CategoryCreate) defaults() {
	if _, ok := cc.mutation.Created(); !ok {
		v := category.DefaultCreated()
		cc.mutation.SetCreated(v)
	}
	if _, ok := cc.mutation.Modified(); !ok {
		v := category.DefaultModified()
		cc.mutation.SetModified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoryCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Category.name"`)}
	}
	if _, ok := cc.mutation.NameJa(); !ok {
		return &ValidationError{Name: "name_ja", err: errors.New(`ent: missing required field "Category.name_ja"`)}
	}
	if _, ok := cc.mutation.Emoji(); !ok {
		return &ValidationError{Name: "emoji", err: errors.New(`ent: missing required field "Category.emoji"`)}
	}
	if _, ok := cc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "Category.created"`)}
	}
	if _, ok := cc.mutation.Modified(); !ok {
		return &ValidationError{Name: "modified", err: errors.New(`ent: missing required field "Category.modified"`)}
	}
	return nil
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		_node = &Category{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(category.Table, sqlgraph.NewFieldSpec(category.FieldID, field.TypeUint32))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.NameJa(); ok {
		_spec.SetField(category.FieldNameJa, field.TypeString, value)
		_node.NameJa = value
	}
	if value, ok := cc.mutation.Emoji(); ok {
		_spec.SetField(category.FieldEmoji, field.TypeString, value)
		_node.Emoji = value
	}
	if value, ok := cc.mutation.Created(); ok {
		_spec.SetField(category.FieldCreated, field.TypeTime, value)
		_node.Created = value
	}
	if value, ok := cc.mutation.Modified(); ok {
		_spec.SetField(category.FieldModified, field.TypeTime, value)
		_node.Modified = value
	}
	return _node, _spec
}

// CategoryCreateBulk is the builder for creating many Category entities in bulk.
type CategoryCreateBulk struct {
	config
	builders []*CategoryCreate
}

// Save creates the Category entities in the database.
func (ccb *CategoryCreateBulk) Save(ctx context.Context) ([]*Category, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Category, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CategoryCreateBulk) SaveX(ctx context.Context) []*Category {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
