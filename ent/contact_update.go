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
	"github.com/cateiru/cateir.com/ent/contact"
	"github.com/cateiru/cateir.com/ent/predicate"
)

// ContactUpdate is the builder for updating Contact entities.
type ContactUpdate struct {
	config
	hooks    []Hook
	mutation *ContactMutation
}

// Where appends a list predicates to the ContactUpdate builder.
func (cu *ContactUpdate) Where(ps ...predicate.Contact) *ContactUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetToUserID sets the "to_user_id" field.
func (cu *ContactUpdate) SetToUserID(u uint32) *ContactUpdate {
	cu.mutation.ResetToUserID()
	cu.mutation.SetToUserID(u)
	return cu
}

// AddToUserID adds u to the "to_user_id" field.
func (cu *ContactUpdate) AddToUserID(u int32) *ContactUpdate {
	cu.mutation.AddToUserID(u)
	return cu
}

// SetTitle sets the "title" field.
func (cu *ContactUpdate) SetTitle(s string) *ContactUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetDetail sets the "detail" field.
func (cu *ContactUpdate) SetDetail(s string) *ContactUpdate {
	cu.mutation.SetDetail(s)
	return cu
}

// SetMail sets the "mail" field.
func (cu *ContactUpdate) SetMail(s string) *ContactUpdate {
	cu.mutation.SetMail(s)
	return cu
}

// SetCategory sets the "category" field.
func (cu *ContactUpdate) SetCategory(s string) *ContactUpdate {
	cu.mutation.SetCategory(s)
	return cu
}

// SetIP sets the "ip" field.
func (cu *ContactUpdate) SetIP(s string) *ContactUpdate {
	cu.mutation.SetIP(s)
	return cu
}

// SetDeviceName sets the "device_name" field.
func (cu *ContactUpdate) SetDeviceName(s string) *ContactUpdate {
	cu.mutation.SetDeviceName(s)
	return cu
}

// SetOs sets the "os" field.
func (cu *ContactUpdate) SetOs(s string) *ContactUpdate {
	cu.mutation.SetOs(s)
	return cu
}

// SetBrowserName sets the "browser_name" field.
func (cu *ContactUpdate) SetBrowserName(s string) *ContactUpdate {
	cu.mutation.SetBrowserName(s)
	return cu
}

// SetCreated sets the "created" field.
func (cu *ContactUpdate) SetCreated(t time.Time) *ContactUpdate {
	cu.mutation.SetCreated(t)
	return cu
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (cu *ContactUpdate) SetNillableCreated(t *time.Time) *ContactUpdate {
	if t != nil {
		cu.SetCreated(*t)
	}
	return cu
}

// SetModified sets the "modified" field.
func (cu *ContactUpdate) SetModified(t time.Time) *ContactUpdate {
	cu.mutation.SetModified(t)
	return cu
}

// Mutation returns the ContactMutation object of the builder.
func (cu *ContactUpdate) Mutation() *ContactMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContactUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContactUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContactUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContactUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ContactUpdate) defaults() {
	if _, ok := cu.mutation.Modified(); !ok {
		v := contact.UpdateDefaultModified()
		cu.mutation.SetModified(v)
	}
}

func (cu *ContactUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contact.Table,
			Columns: contact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: contact.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.ToUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contact.FieldToUserID,
		})
	}
	if value, ok := cu.mutation.AddedToUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contact.FieldToUserID,
		})
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldTitle,
		})
	}
	if value, ok := cu.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldDetail,
		})
	}
	if value, ok := cu.mutation.Mail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldMail,
		})
	}
	if value, ok := cu.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldCategory,
		})
	}
	if value, ok := cu.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldIP,
		})
	}
	if value, ok := cu.mutation.DeviceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldDeviceName,
		})
	}
	if value, ok := cu.mutation.Os(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldOs,
		})
	}
	if value, ok := cu.mutation.BrowserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldBrowserName,
		})
	}
	if value, ok := cu.mutation.Created(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contact.FieldCreated,
		})
	}
	if value, ok := cu.mutation.Modified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contact.FieldModified,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ContactUpdateOne is the builder for updating a single Contact entity.
type ContactUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContactMutation
}

// SetToUserID sets the "to_user_id" field.
func (cuo *ContactUpdateOne) SetToUserID(u uint32) *ContactUpdateOne {
	cuo.mutation.ResetToUserID()
	cuo.mutation.SetToUserID(u)
	return cuo
}

// AddToUserID adds u to the "to_user_id" field.
func (cuo *ContactUpdateOne) AddToUserID(u int32) *ContactUpdateOne {
	cuo.mutation.AddToUserID(u)
	return cuo
}

// SetTitle sets the "title" field.
func (cuo *ContactUpdateOne) SetTitle(s string) *ContactUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetDetail sets the "detail" field.
func (cuo *ContactUpdateOne) SetDetail(s string) *ContactUpdateOne {
	cuo.mutation.SetDetail(s)
	return cuo
}

// SetMail sets the "mail" field.
func (cuo *ContactUpdateOne) SetMail(s string) *ContactUpdateOne {
	cuo.mutation.SetMail(s)
	return cuo
}

// SetCategory sets the "category" field.
func (cuo *ContactUpdateOne) SetCategory(s string) *ContactUpdateOne {
	cuo.mutation.SetCategory(s)
	return cuo
}

// SetIP sets the "ip" field.
func (cuo *ContactUpdateOne) SetIP(s string) *ContactUpdateOne {
	cuo.mutation.SetIP(s)
	return cuo
}

// SetDeviceName sets the "device_name" field.
func (cuo *ContactUpdateOne) SetDeviceName(s string) *ContactUpdateOne {
	cuo.mutation.SetDeviceName(s)
	return cuo
}

// SetOs sets the "os" field.
func (cuo *ContactUpdateOne) SetOs(s string) *ContactUpdateOne {
	cuo.mutation.SetOs(s)
	return cuo
}

// SetBrowserName sets the "browser_name" field.
func (cuo *ContactUpdateOne) SetBrowserName(s string) *ContactUpdateOne {
	cuo.mutation.SetBrowserName(s)
	return cuo
}

// SetCreated sets the "created" field.
func (cuo *ContactUpdateOne) SetCreated(t time.Time) *ContactUpdateOne {
	cuo.mutation.SetCreated(t)
	return cuo
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (cuo *ContactUpdateOne) SetNillableCreated(t *time.Time) *ContactUpdateOne {
	if t != nil {
		cuo.SetCreated(*t)
	}
	return cuo
}

// SetModified sets the "modified" field.
func (cuo *ContactUpdateOne) SetModified(t time.Time) *ContactUpdateOne {
	cuo.mutation.SetModified(t)
	return cuo
}

// Mutation returns the ContactMutation object of the builder.
func (cuo *ContactUpdateOne) Mutation() *ContactMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContactUpdateOne) Select(field string, fields ...string) *ContactUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Contact entity.
func (cuo *ContactUpdateOne) Save(ctx context.Context) (*Contact, error) {
	var (
		err  error
		node *Contact
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContactMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Contact)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ContactMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContactUpdateOne) SaveX(ctx context.Context) *Contact {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContactUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContactUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ContactUpdateOne) defaults() {
	if _, ok := cuo.mutation.Modified(); !ok {
		v := contact.UpdateDefaultModified()
		cuo.mutation.SetModified(v)
	}
}

func (cuo *ContactUpdateOne) sqlSave(ctx context.Context) (_node *Contact, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contact.Table,
			Columns: contact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: contact.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Contact.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contact.FieldID)
		for _, f := range fields {
			if !contact.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.ToUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contact.FieldToUserID,
		})
	}
	if value, ok := cuo.mutation.AddedToUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contact.FieldToUserID,
		})
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldTitle,
		})
	}
	if value, ok := cuo.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldDetail,
		})
	}
	if value, ok := cuo.mutation.Mail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldMail,
		})
	}
	if value, ok := cuo.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldCategory,
		})
	}
	if value, ok := cuo.mutation.IP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldIP,
		})
	}
	if value, ok := cuo.mutation.DeviceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldDeviceName,
		})
	}
	if value, ok := cuo.mutation.Os(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldOs,
		})
	}
	if value, ok := cuo.mutation.BrowserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contact.FieldBrowserName,
		})
	}
	if value, ok := cuo.mutation.Created(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contact.FieldCreated,
		})
	}
	if value, ok := cuo.mutation.Modified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contact.FieldModified,
		})
	}
	_node = &Contact{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contact.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
