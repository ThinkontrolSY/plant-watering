// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"plant-watering/ent/predicate"
	"plant-watering/ent/waterlog"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WaterLogDelete is the builder for deleting a WaterLog entity.
type WaterLogDelete struct {
	config
	hooks    []Hook
	mutation *WaterLogMutation
}

// Where appends a list predicates to the WaterLogDelete builder.
func (wld *WaterLogDelete) Where(ps ...predicate.WaterLog) *WaterLogDelete {
	wld.mutation.Where(ps...)
	return wld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wld *WaterLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wld.sqlExec, wld.mutation, wld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wld *WaterLogDelete) ExecX(ctx context.Context) int {
	n, err := wld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wld *WaterLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(waterlog.Table, sqlgraph.NewFieldSpec(waterlog.FieldID, field.TypeInt))
	if ps := wld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wld.mutation.done = true
	return affected, err
}

// WaterLogDeleteOne is the builder for deleting a single WaterLog entity.
type WaterLogDeleteOne struct {
	wld *WaterLogDelete
}

// Where appends a list predicates to the WaterLogDelete builder.
func (wldo *WaterLogDeleteOne) Where(ps ...predicate.WaterLog) *WaterLogDeleteOne {
	wldo.wld.mutation.Where(ps...)
	return wldo
}

// Exec executes the deletion query.
func (wldo *WaterLogDeleteOne) Exec(ctx context.Context) error {
	n, err := wldo.wld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{waterlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wldo *WaterLogDeleteOne) ExecX(ctx context.Context) {
	if err := wldo.Exec(ctx); err != nil {
		panic(err)
	}
}
