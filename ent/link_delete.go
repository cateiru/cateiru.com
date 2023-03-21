// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cateiru/cateiru.com/ent/link"
	"github.com/cateiru/cateiru.com/ent/predicate"
)

// LinkDelete is the builder for deleting a Link entity.
type LinkDelete struct {
	config
	hooks    []Hook
	mutation *LinkMutation
}

// Where appends a list predicates to the LinkDelete builder.
func (ld *LinkDelete) Where(ps ...predicate.Link) *LinkDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LinkDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LinkMutation](ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LinkDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LinkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(link.Table, sqlgraph.NewFieldSpec(link.FieldID, field.TypeUint32))
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LinkDeleteOne is the builder for deleting a single Link entity.
type LinkDeleteOne struct {
	ld *LinkDelete
}

// Where appends a list predicates to the LinkDelete builder.
func (ldo *LinkDeleteOne) Where(ps ...predicate.Link) *LinkDeleteOne {
	ldo.ld.mutation.Where(ps...)
	return ldo
}

// Exec executes the deletion query.
func (ldo *LinkDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{link.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LinkDeleteOne) ExecX(ctx context.Context) {
	if err := ldo.Exec(ctx); err != nil {
		panic(err)
	}
}
