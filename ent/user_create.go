// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cateiru/cateiru.com/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetGivenName sets the "given_name" field.
func (uc *UserCreate) SetGivenName(s string) *UserCreate {
	uc.mutation.SetGivenName(s)
	return uc
}

// SetFamilyName sets the "family_name" field.
func (uc *UserCreate) SetFamilyName(s string) *UserCreate {
	uc.mutation.SetFamilyName(s)
	return uc
}

// SetGivenNameJa sets the "given_name_ja" field.
func (uc *UserCreate) SetGivenNameJa(s string) *UserCreate {
	uc.mutation.SetGivenNameJa(s)
	return uc
}

// SetFamilyNameJa sets the "family_name_ja" field.
func (uc *UserCreate) SetFamilyNameJa(s string) *UserCreate {
	uc.mutation.SetFamilyNameJa(s)
	return uc
}

// SetUserID sets the "user_id" field.
func (uc *UserCreate) SetUserID(s string) *UserCreate {
	uc.mutation.SetUserID(s)
	return uc
}

// SetMail sets the "mail" field.
func (uc *UserCreate) SetMail(s string) *UserCreate {
	uc.mutation.SetMail(s)
	return uc
}

// SetBirthDate sets the "birth_date" field.
func (uc *UserCreate) SetBirthDate(t time.Time) *UserCreate {
	uc.mutation.SetBirthDate(t)
	return uc
}

// SetLocation sets the "location" field.
func (uc *UserCreate) SetLocation(s string) *UserCreate {
	uc.mutation.SetLocation(s)
	return uc
}

// SetLocationJa sets the "location_ja" field.
func (uc *UserCreate) SetLocationJa(s string) *UserCreate {
	uc.mutation.SetLocationJa(s)
	return uc
}

// SetSSOToken sets the "sso_token" field.
func (uc *UserCreate) SetSSOToken(s string) *UserCreate {
	uc.mutation.SetSSOToken(s)
	return uc
}

// SetAvatarURL sets the "avatar_url" field.
func (uc *UserCreate) SetAvatarURL(s string) *UserCreate {
	uc.mutation.SetAvatarURL(s)
	return uc
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatarURL(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatarURL(*s)
	}
	return uc
}

// SetCreated sets the "created" field.
func (uc *UserCreate) SetCreated(t time.Time) *UserCreate {
	uc.mutation.SetCreated(t)
	return uc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreated(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreated(*t)
	}
	return uc
}

// SetModified sets the "modified" field.
func (uc *UserCreate) SetModified(t time.Time) *UserCreate {
	uc.mutation.SetModified(t)
	return uc
}

// SetNillableModified sets the "modified" field if the given value is not nil.
func (uc *UserCreate) SetNillableModified(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetModified(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uint32) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Created(); !ok {
		v := user.DefaultCreated()
		uc.mutation.SetCreated(v)
	}
	if _, ok := uc.mutation.Modified(); !ok {
		v := user.DefaultModified()
		uc.mutation.SetModified(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.GivenName(); !ok {
		return &ValidationError{Name: "given_name", err: errors.New(`ent: missing required field "User.given_name"`)}
	}
	if _, ok := uc.mutation.FamilyName(); !ok {
		return &ValidationError{Name: "family_name", err: errors.New(`ent: missing required field "User.family_name"`)}
	}
	if _, ok := uc.mutation.GivenNameJa(); !ok {
		return &ValidationError{Name: "given_name_ja", err: errors.New(`ent: missing required field "User.given_name_ja"`)}
	}
	if _, ok := uc.mutation.FamilyNameJa(); !ok {
		return &ValidationError{Name: "family_name_ja", err: errors.New(`ent: missing required field "User.family_name_ja"`)}
	}
	if _, ok := uc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "User.user_id"`)}
	}
	if _, ok := uc.mutation.Mail(); !ok {
		return &ValidationError{Name: "mail", err: errors.New(`ent: missing required field "User.mail"`)}
	}
	if _, ok := uc.mutation.BirthDate(); !ok {
		return &ValidationError{Name: "birth_date", err: errors.New(`ent: missing required field "User.birth_date"`)}
	}
	if _, ok := uc.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "User.location"`)}
	}
	if _, ok := uc.mutation.LocationJa(); !ok {
		return &ValidationError{Name: "location_ja", err: errors.New(`ent: missing required field "User.location_ja"`)}
	}
	if _, ok := uc.mutation.SSOToken(); !ok {
		return &ValidationError{Name: "sso_token", err: errors.New(`ent: missing required field "User.sso_token"`)}
	}
	if _, ok := uc.mutation.Created(); !ok {
		return &ValidationError{Name: "created", err: errors.New(`ent: missing required field "User.created"`)}
	}
	if _, ok := uc.mutation.Modified(); !ok {
		return &ValidationError{Name: "modified", err: errors.New(`ent: missing required field "User.modified"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
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

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.GivenName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGivenName,
		})
		_node.GivenName = value
	}
	if value, ok := uc.mutation.FamilyName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFamilyName,
		})
		_node.FamilyName = value
	}
	if value, ok := uc.mutation.GivenNameJa(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldGivenNameJa,
		})
		_node.GivenNameJa = value
	}
	if value, ok := uc.mutation.FamilyNameJa(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFamilyNameJa,
		})
		_node.FamilyNameJa = value
	}
	if value, ok := uc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := uc.mutation.Mail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldMail,
		})
		_node.Mail = value
	}
	if value, ok := uc.mutation.BirthDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldBirthDate,
		})
		_node.BirthDate = value
	}
	if value, ok := uc.mutation.Location(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLocation,
		})
		_node.Location = value
	}
	if value, ok := uc.mutation.LocationJa(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLocationJa,
		})
		_node.LocationJa = value
	}
	if value, ok := uc.mutation.SSOToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldSSOToken,
		})
		_node.SSOToken = value
	}
	if value, ok := uc.mutation.AvatarURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatarURL,
		})
		_node.AvatarURL = value
	}
	if value, ok := uc.mutation.Created(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreated,
		})
		_node.Created = value
	}
	if value, ok := uc.mutation.Modified(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldModified,
		})
		_node.Modified = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
