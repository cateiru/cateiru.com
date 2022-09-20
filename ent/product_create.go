// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cateiru/cateir.com/ent/product"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (pc *ProductCreate) SetUserID(u uint32) *ProductCreate {
	pc.mutation.SetUserID(u)
	return pc
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNameJa sets the "name_ja" field.
func (pc *ProductCreate) SetNameJa(s string) *ProductCreate {
	pc.mutation.SetNameJa(s)
	return pc
}

// SetDetail sets the "detail" field.
func (pc *ProductCreate) SetDetail(s string) *ProductCreate {
	pc.mutation.SetDetail(s)
	return pc
}

// SetDetailJa sets the "detail_ja" field.
func (pc *ProductCreate) SetDetailJa(s string) *ProductCreate {
	pc.mutation.SetDetailJa(s)
	return pc
}

// SetSiteURL sets the "site_url" field.
func (pc *ProductCreate) SetSiteURL(s string) *ProductCreate {
	pc.mutation.SetSiteURL(s)
	return pc
}

// SetGithubURL sets the "github_url" field.
func (pc *ProductCreate) SetGithubURL(s string) *ProductCreate {
	pc.mutation.SetGithubURL(s)
	return pc
}

// SetNillableGithubURL sets the "github_url" field if the given value is not nil.
func (pc *ProductCreate) SetNillableGithubURL(s *string) *ProductCreate {
	if s != nil {
		pc.SetGithubURL(*s)
	}
	return pc
}

// SetDevTime sets the "dev_time" field.
func (pc *ProductCreate) SetDevTime(t time.Time) *ProductCreate {
	pc.mutation.SetDevTime(t)
	return pc
}

// SetCreated sets the "created" field.
func (pc *ProductCreate) SetCreated(t time.Time) *ProductCreate {
	pc.mutation.SetCreated(t)
	return pc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCreated(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetCreated(*t)
	}
	return pc
}

// SetModified sets the "modified" field.
func (pc *ProductCreate) SetModified(t time.Time) *ProductCreate {
	pc.mutation.SetModified(t)
	return pc
}

// SetNillableModified sets the "modified" field if the given value is not nil.
func (pc *ProductCreate) SetNillableModified(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetModified(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProductCreate) SetID(u uint32) *ProductCreate {
	pc.mutation.SetID(u)
	return pc
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Product)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProductMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.Created(); !ok {
		v := product.DefaultCreated()
		pc.mutation.SetCreated(v)
	}
	if _, ok := pc.mutation.Modified(); !ok {
		v := product.DefaultModified()
		pc.mutation.SetModified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Product.user_id"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Product.name"`)}
	}
	if _, ok := pc.mutation.NameJa(); !ok {
		return &ValidationError{Name: "name_ja", err: errors.New(`ent: missing required field "Product.name_ja"`)}
	}
	if _, ok := pc.mutation.Detail(); !ok {
		return &ValidationError{Name: "detail", err: errors.New(`ent: missing required field "Product.detail"`)}
	}
	if _, ok := pc.mutation.DetailJa(); !ok {
		return &ValidationError{Name: "detail_ja", err: errors.New(`ent: missing required field "Product.detail_ja"`)}
	}
	if _, ok := pc.mutation.SiteURL(); !ok {
		return &ValidationError{Name: "site_url", err: errors.New(`ent: missing required field "Product.site_url"`)}
	}
	if _, ok := pc.mutation.DevTime(); !ok {
		return &ValidationError{Name: "dev_time", err: errors.New(`ent: missing required field "Product.dev_time"`)}
	}
	if _, ok := pc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "Product.created"`)}
	}
	if _, ok := pc.mutation.Modified(); !ok {
		return &ValidationError{Name: "modified", err: errors.New(`ent: missing required field "Product.modified"`)}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: product.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: product.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: product.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.NameJa(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldNameJa,
		})
		_node.NameJa = value
	}
	if value, ok := pc.mutation.Detail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDetail,
		})
		_node.Detail = value
	}
	if value, ok := pc.mutation.DetailJa(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDetailJa,
		})
		_node.DetailJa = value
	}
	if value, ok := pc.mutation.SiteURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldSiteURL,
		})
		_node.SiteURL = value
	}
	if value, ok := pc.mutation.GithubURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldGithubURL,
		})
		_node.GithubURL = value
	}
	if value, ok := pc.mutation.DevTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldDevTime,
		})
		_node.DevTime = value
	}
	if value, ok := pc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := pc.mutation.Modified(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldModified,
		})
		_node.Modified = value
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
