// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cateiru/cateiru.com/ent/link"
)

// LinkCreate is the builder for creating a Link entity.
type LinkCreate struct {
	config
	mutation *LinkMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (lc *LinkCreate) SetUserID(u uint32) *LinkCreate {
	lc.mutation.SetUserID(u)
	return lc
}

// SetName sets the "name" field.
func (lc *LinkCreate) SetName(s string) *LinkCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetNameJa sets the "name_ja" field.
func (lc *LinkCreate) SetNameJa(s string) *LinkCreate {
	lc.mutation.SetNameJa(s)
	return lc
}

// SetSiteURL sets the "site_url" field.
func (lc *LinkCreate) SetSiteURL(s string) *LinkCreate {
	lc.mutation.SetSiteURL(s)
	return lc
}

// SetFaviconURL sets the "favicon_url" field.
func (lc *LinkCreate) SetFaviconURL(s string) *LinkCreate {
	lc.mutation.SetFaviconURL(s)
	return lc
}

// SetCategoryID sets the "category_id" field.
func (lc *LinkCreate) SetCategoryID(u uint32) *LinkCreate {
	lc.mutation.SetCategoryID(u)
	return lc
}

// SetCreated sets the "created" field.
func (lc *LinkCreate) SetCreated(t time.Time) *LinkCreate {
	lc.mutation.SetCreated(t)
	return lc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (lc *LinkCreate) SetNillableCreated(t *time.Time) *LinkCreate {
	if t != nil {
		lc.SetCreated(*t)
	}
	return lc
}

// SetModified sets the "modified" field.
func (lc *LinkCreate) SetModified(t time.Time) *LinkCreate {
	lc.mutation.SetModified(t)
	return lc
}

// SetNillableModified sets the "modified" field if the given value is not nil.
func (lc *LinkCreate) SetNillableModified(t *time.Time) *LinkCreate {
	if t != nil {
		lc.SetModified(*t)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LinkCreate) SetID(u uint32) *LinkCreate {
	lc.mutation.SetID(u)
	return lc
}

// Mutation returns the LinkMutation object of the builder.
func (lc *LinkCreate) Mutation() *LinkMutation {
	return lc.mutation
}

// Save creates the Link in the database.
func (lc *LinkCreate) Save(ctx context.Context) (*Link, error) {
	var (
		err  error
		node *Link
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, lc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Link)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LinkMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LinkCreate) SaveX(ctx context.Context) *Link {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LinkCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LinkCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LinkCreate) defaults() {
	if _, ok := lc.mutation.Created(); !ok {
		v := link.DefaultCreated()
		lc.mutation.SetCreated(v)
	}
	if _, ok := lc.mutation.Modified(); !ok {
		v := link.DefaultModified()
		lc.mutation.SetModified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LinkCreate) check() error {
	if _, ok := lc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Link.user_id"`)}
	}
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Link.name"`)}
	}
	if _, ok := lc.mutation.NameJa(); !ok {
		return &ValidationError{Name: "name_ja", err: errors.New(`ent: missing required field "Link.name_ja"`)}
	}
	if _, ok := lc.mutation.SiteURL(); !ok {
		return &ValidationError{Name: "site_url", err: errors.New(`ent: missing required field "Link.site_url"`)}
	}
	if _, ok := lc.mutation.FaviconURL(); !ok {
		return &ValidationError{Name: "favicon_url", err: errors.New(`ent: missing required field "Link.favicon_url"`)}
	}
	if _, ok := lc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`ent: missing required field "Link.category_id"`)}
	}
	if _, ok := lc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "Link.created"`)}
	}
	if _, ok := lc.mutation.Modified(); !ok {
		return &ValidationError{Name: "modified", err: errors.New(`ent: missing required field "Link.modified"`)}
	}
	return nil
}

func (lc *LinkCreate) sqlSave(ctx context.Context) (*Link, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (lc *LinkCreate) createSpec() (*Link, *sqlgraph.CreateSpec) {
	var (
		_node = &Link{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: link.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: link.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: link.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := lc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldName,
		})
		_node.Name = value
	}
	if value, ok := lc.mutation.NameJa(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldNameJa,
		})
		_node.NameJa = value
	}
	if value, ok := lc.mutation.SiteURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldSiteURL,
		})
		_node.SiteURL = value
	}
	if value, ok := lc.mutation.FaviconURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: link.FieldFaviconURL,
		})
		_node.FaviconURL = value
	}
	if value, ok := lc.mutation.CategoryID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: link.FieldCategoryID,
		})
		_node.CategoryID = value
	}
	if value, ok := lc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: link.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := lc.mutation.Modified(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: link.FieldModified,
		})
		_node.Modified = value
	}
	return _node, _spec
}

// LinkCreateBulk is the builder for creating many Link entities in bulk.
type LinkCreateBulk struct {
	config
	builders []*LinkCreate
}

// Save creates the Link entities in the database.
func (lcb *LinkCreateBulk) Save(ctx context.Context) ([]*Link, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Link, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinkMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LinkCreateBulk) SaveX(ctx context.Context) []*Link {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LinkCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LinkCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}