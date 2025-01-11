// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kratos-admin/app/admin/service/internal/data/ent/adminloginlog"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminLoginLogDelete is the builder for deleting a AdminLoginLog entity.
type AdminLoginLogDelete struct {
	config
	hooks    []Hook
	mutation *AdminLoginLogMutation
}

// Where appends a list predicates to the AdminLoginLogDelete builder.
func (alld *AdminLoginLogDelete) Where(ps ...predicate.AdminLoginLog) *AdminLoginLogDelete {
	alld.mutation.Where(ps...)
	return alld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (alld *AdminLoginLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, alld.sqlExec, alld.mutation, alld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (alld *AdminLoginLogDelete) ExecX(ctx context.Context) int {
	n, err := alld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (alld *AdminLoginLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(adminloginlog.Table, sqlgraph.NewFieldSpec(adminloginlog.FieldID, field.TypeUint32))
	if ps := alld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, alld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	alld.mutation.done = true
	return affected, err
}

// AdminLoginLogDeleteOne is the builder for deleting a single AdminLoginLog entity.
type AdminLoginLogDeleteOne struct {
	alld *AdminLoginLogDelete
}

// Where appends a list predicates to the AdminLoginLogDelete builder.
func (alldo *AdminLoginLogDeleteOne) Where(ps ...predicate.AdminLoginLog) *AdminLoginLogDeleteOne {
	alldo.alld.mutation.Where(ps...)
	return alldo
}

// Exec executes the deletion query.
func (alldo *AdminLoginLogDeleteOne) Exec(ctx context.Context) error {
	n, err := alldo.alld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{adminloginlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (alldo *AdminLoginLogDeleteOne) ExecX(ctx context.Context) {
	if err := alldo.Exec(ctx); err != nil {
		panic(err)
	}
}
