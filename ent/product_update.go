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
	"github.com/cateiru/cateiru.com/ent/predicate"
	"github.com/cateiru/cateiru.com/ent/product"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *ProductUpdate) SetUserID(u uint32) *ProductUpdate {
	pu.mutation.ResetUserID()
	pu.mutation.SetUserID(u)
	return pu
}

// AddUserID adds u to the "user_id" field.
func (pu *ProductUpdate) AddUserID(u int32) *ProductUpdate {
	pu.mutation.AddUserID(u)
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNameJa sets the "name_ja" field.
func (pu *ProductUpdate) SetNameJa(s string) *ProductUpdate {
	pu.mutation.SetNameJa(s)
	return pu
}

// SetDetail sets the "detail" field.
func (pu *ProductUpdate) SetDetail(s string) *ProductUpdate {
	pu.mutation.SetDetail(s)
	return pu
}

// SetDetailJa sets the "detail_ja" field.
func (pu *ProductUpdate) SetDetailJa(s string) *ProductUpdate {
	pu.mutation.SetDetailJa(s)
	return pu
}

// SetSiteURL sets the "site_url" field.
func (pu *ProductUpdate) SetSiteURL(s string) *ProductUpdate {
	pu.mutation.SetSiteURL(s)
	return pu
}

// SetGithubURL sets the "github_url" field.
func (pu *ProductUpdate) SetGithubURL(s string) *ProductUpdate {
	pu.mutation.SetGithubURL(s)
	return pu
}

// SetNillableGithubURL sets the "github_url" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableGithubURL(s *string) *ProductUpdate {
	if s != nil {
		pu.SetGithubURL(*s)
	}
	return pu
}

// ClearGithubURL clears the value of the "github_url" field.
func (pu *ProductUpdate) ClearGithubURL() *ProductUpdate {
	pu.mutation.ClearGithubURL()
	return pu
}

// SetDevTime sets the "dev_time" field.
func (pu *ProductUpdate) SetDevTime(t time.Time) *ProductUpdate {
	pu.mutation.SetDevTime(t)
	return pu
}

// SetCreated sets the "created" field.
func (pu *ProductUpdate) SetCreated(t time.Time) *ProductUpdate {
	pu.mutation.SetCreated(t)
	return pu
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableCreated(t *time.Time) *ProductUpdate {
	if t != nil {
		pu.SetCreated(*t)
	}
	return pu
}

// SetModified sets the "modified" field.
func (pu *ProductUpdate) SetModified(t time.Time) *ProductUpdate {
	pu.mutation.SetModified(t)
	return pu
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProductUpdate) defaults() {
	if _, ok := pu.mutation.Modified(); !ok {
		v := product.UpdateDefaultModified()
		pu.mutation.SetModified(v)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: product.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: product.FieldUserID,
		})
	}
	if value, ok := pu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: product.FieldUserID,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
	}
	if value, ok := pu.mutation.NameJa(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldNameJa,
		})
	}
	if value, ok := pu.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDetail,
		})
	}
	if value, ok := pu.mutation.DetailJa(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDetailJa,
		})
	}
	if value, ok := pu.mutation.SiteURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldSiteURL,
		})
	}
	if value, ok := pu.mutation.GithubURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldGithubURL,
		})
	}
	if pu.mutation.GithubURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: product.FieldGithubURL,
		})
	}
	if value, ok := pu.mutation.DevTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldDevTime,
		})
	}
	if value, ok := pu.mutation.Created(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldCreated,
		})
	}
	if value, ok := pu.mutation.Modified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldModified,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductMutation
}

// SetUserID sets the "user_id" field.
func (puo *ProductUpdateOne) SetUserID(u uint32) *ProductUpdateOne {
	puo.mutation.ResetUserID()
	puo.mutation.SetUserID(u)
	return puo
}

// AddUserID adds u to the "user_id" field.
func (puo *ProductUpdateOne) AddUserID(u int32) *ProductUpdateOne {
	puo.mutation.AddUserID(u)
	return puo
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNameJa sets the "name_ja" field.
func (puo *ProductUpdateOne) SetNameJa(s string) *ProductUpdateOne {
	puo.mutation.SetNameJa(s)
	return puo
}

// SetDetail sets the "detail" field.
func (puo *ProductUpdateOne) SetDetail(s string) *ProductUpdateOne {
	puo.mutation.SetDetail(s)
	return puo
}

// SetDetailJa sets the "detail_ja" field.
func (puo *ProductUpdateOne) SetDetailJa(s string) *ProductUpdateOne {
	puo.mutation.SetDetailJa(s)
	return puo
}

// SetSiteURL sets the "site_url" field.
func (puo *ProductUpdateOne) SetSiteURL(s string) *ProductUpdateOne {
	puo.mutation.SetSiteURL(s)
	return puo
}

// SetGithubURL sets the "github_url" field.
func (puo *ProductUpdateOne) SetGithubURL(s string) *ProductUpdateOne {
	puo.mutation.SetGithubURL(s)
	return puo
}

// SetNillableGithubURL sets the "github_url" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableGithubURL(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetGithubURL(*s)
	}
	return puo
}

// ClearGithubURL clears the value of the "github_url" field.
func (puo *ProductUpdateOne) ClearGithubURL() *ProductUpdateOne {
	puo.mutation.ClearGithubURL()
	return puo
}

// SetDevTime sets the "dev_time" field.
func (puo *ProductUpdateOne) SetDevTime(t time.Time) *ProductUpdateOne {
	puo.mutation.SetDevTime(t)
	return puo
}

// SetCreated sets the "created" field.
func (puo *ProductUpdateOne) SetCreated(t time.Time) *ProductUpdateOne {
	puo.mutation.SetCreated(t)
	return puo
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableCreated(t *time.Time) *ProductUpdateOne {
	if t != nil {
		puo.SetCreated(*t)
	}
	return puo
}

// SetModified sets the "modified" field.
func (puo *ProductUpdateOne) SetModified(t time.Time) *ProductUpdateOne {
	puo.mutation.SetModified(t)
	return puo
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductUpdateOne) Select(field string, fields ...string) *ProductUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProductUpdateOne) defaults() {
	if _, ok := puo.mutation.Modified(); !ok {
		v := product.UpdateDefaultModified()
		puo.mutation.SetModified(v)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: product.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Product.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for _, f := range fields {
			if !product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: product.FieldUserID,
		})
	}
	if value, ok := puo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: product.FieldUserID,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
	}
	if value, ok := puo.mutation.NameJa(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldNameJa,
		})
	}
	if value, ok := puo.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDetail,
		})
	}
	if value, ok := puo.mutation.DetailJa(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDetailJa,
		})
	}
	if value, ok := puo.mutation.SiteURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldSiteURL,
		})
	}
	if value, ok := puo.mutation.GithubURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldGithubURL,
		})
	}
	if puo.mutation.GithubURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: product.FieldGithubURL,
		})
	}
	if value, ok := puo.mutation.DevTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldDevTime,
		})
	}
	if value, ok := puo.mutation.Created(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldCreated,
		})
	}
	if value, ok := puo.mutation.Modified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: product.FieldModified,
		})
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}