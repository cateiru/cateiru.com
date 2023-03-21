// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cateiru/cateiru.com/ent/biography"
	"github.com/cateiru/cateiru.com/ent/predicate"
)

// BiographyDelete is the builder for deleting a Biography entity.
type BiographyDelete struct {
	config
	hooks    []Hook
	mutation *BiographyMutation
}

// Where appends a list predicates to the BiographyDelete builder.
func (bd *BiographyDelete) Where(ps ...predicate.Biography) *BiographyDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BiographyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, BiographyMutation](ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BiographyDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BiographyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(biography.Table, sqlgraph.NewFieldSpec(biography.FieldID, field.TypeUint32))
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BiographyDeleteOne is the builder for deleting a single Biography entity.
type BiographyDeleteOne struct {
	bd *BiographyDelete
}

// Where appends a list predicates to the BiographyDelete builder.
func (bdo *BiographyDeleteOne) Where(ps ...predicate.Biography) *BiographyDeleteOne {
	bdo.bd.mutation.Where(ps...)
	return bdo
}

// Exec executes the deletion query.
func (bdo *BiographyDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{biography.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BiographyDeleteOne) ExecX(ctx context.Context) {
	if err := bdo.Exec(ctx); err != nil {
		panic(err)
	}
}
