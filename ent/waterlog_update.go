// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"plant-watering/ent/predicate"
	"plant-watering/ent/waterlog"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WaterLogUpdate is the builder for updating WaterLog entities.
type WaterLogUpdate struct {
	config
	hooks    []Hook
	mutation *WaterLogMutation
}

// Where appends a list predicates to the WaterLogUpdate builder.
func (wlu *WaterLogUpdate) Where(ps ...predicate.WaterLog) *WaterLogUpdate {
	wlu.mutation.Where(ps...)
	return wlu
}

// SetSeconds sets the "seconds" field.
func (wlu *WaterLogUpdate) SetSeconds(i int) *WaterLogUpdate {
	wlu.mutation.ResetSeconds()
	wlu.mutation.SetSeconds(i)
	return wlu
}

// SetNillableSeconds sets the "seconds" field if the given value is not nil.
func (wlu *WaterLogUpdate) SetNillableSeconds(i *int) *WaterLogUpdate {
	if i != nil {
		wlu.SetSeconds(*i)
	}
	return wlu
}

// AddSeconds adds i to the "seconds" field.
func (wlu *WaterLogUpdate) AddSeconds(i int) *WaterLogUpdate {
	wlu.mutation.AddSeconds(i)
	return wlu
}

// SetChannel sets the "channel" field.
func (wlu *WaterLogUpdate) SetChannel(s string) *WaterLogUpdate {
	wlu.mutation.SetChannel(s)
	return wlu
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (wlu *WaterLogUpdate) SetNillableChannel(s *string) *WaterLogUpdate {
	if s != nil {
		wlu.SetChannel(*s)
	}
	return wlu
}

// SetManual sets the "manual" field.
func (wlu *WaterLogUpdate) SetManual(b bool) *WaterLogUpdate {
	wlu.mutation.SetManual(b)
	return wlu
}

// SetNillableManual sets the "manual" field if the given value is not nil.
func (wlu *WaterLogUpdate) SetNillableManual(b *bool) *WaterLogUpdate {
	if b != nil {
		wlu.SetManual(*b)
	}
	return wlu
}

// SetTime sets the "time" field.
func (wlu *WaterLogUpdate) SetTime(t time.Time) *WaterLogUpdate {
	wlu.mutation.SetTime(t)
	return wlu
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (wlu *WaterLogUpdate) SetNillableTime(t *time.Time) *WaterLogUpdate {
	if t != nil {
		wlu.SetTime(*t)
	}
	return wlu
}

// Mutation returns the WaterLogMutation object of the builder.
func (wlu *WaterLogUpdate) Mutation() *WaterLogMutation {
	return wlu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wlu *WaterLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wlu.sqlSave, wlu.mutation, wlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wlu *WaterLogUpdate) SaveX(ctx context.Context) int {
	affected, err := wlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wlu *WaterLogUpdate) Exec(ctx context.Context) error {
	_, err := wlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wlu *WaterLogUpdate) ExecX(ctx context.Context) {
	if err := wlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wlu *WaterLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(waterlog.Table, waterlog.Columns, sqlgraph.NewFieldSpec(waterlog.FieldID, field.TypeInt))
	if ps := wlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wlu.mutation.Seconds(); ok {
		_spec.SetField(waterlog.FieldSeconds, field.TypeInt, value)
	}
	if value, ok := wlu.mutation.AddedSeconds(); ok {
		_spec.AddField(waterlog.FieldSeconds, field.TypeInt, value)
	}
	if value, ok := wlu.mutation.Channel(); ok {
		_spec.SetField(waterlog.FieldChannel, field.TypeString, value)
	}
	if value, ok := wlu.mutation.Manual(); ok {
		_spec.SetField(waterlog.FieldManual, field.TypeBool, value)
	}
	if value, ok := wlu.mutation.Time(); ok {
		_spec.SetField(waterlog.FieldTime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{waterlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wlu.mutation.done = true
	return n, nil
}

// WaterLogUpdateOne is the builder for updating a single WaterLog entity.
type WaterLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WaterLogMutation
}

// SetSeconds sets the "seconds" field.
func (wluo *WaterLogUpdateOne) SetSeconds(i int) *WaterLogUpdateOne {
	wluo.mutation.ResetSeconds()
	wluo.mutation.SetSeconds(i)
	return wluo
}

// SetNillableSeconds sets the "seconds" field if the given value is not nil.
func (wluo *WaterLogUpdateOne) SetNillableSeconds(i *int) *WaterLogUpdateOne {
	if i != nil {
		wluo.SetSeconds(*i)
	}
	return wluo
}

// AddSeconds adds i to the "seconds" field.
func (wluo *WaterLogUpdateOne) AddSeconds(i int) *WaterLogUpdateOne {
	wluo.mutation.AddSeconds(i)
	return wluo
}

// SetChannel sets the "channel" field.
func (wluo *WaterLogUpdateOne) SetChannel(s string) *WaterLogUpdateOne {
	wluo.mutation.SetChannel(s)
	return wluo
}

// SetNillableChannel sets the "channel" field if the given value is not nil.
func (wluo *WaterLogUpdateOne) SetNillableChannel(s *string) *WaterLogUpdateOne {
	if s != nil {
		wluo.SetChannel(*s)
	}
	return wluo
}

// SetManual sets the "manual" field.
func (wluo *WaterLogUpdateOne) SetManual(b bool) *WaterLogUpdateOne {
	wluo.mutation.SetManual(b)
	return wluo
}

// SetNillableManual sets the "manual" field if the given value is not nil.
func (wluo *WaterLogUpdateOne) SetNillableManual(b *bool) *WaterLogUpdateOne {
	if b != nil {
		wluo.SetManual(*b)
	}
	return wluo
}

// SetTime sets the "time" field.
func (wluo *WaterLogUpdateOne) SetTime(t time.Time) *WaterLogUpdateOne {
	wluo.mutation.SetTime(t)
	return wluo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (wluo *WaterLogUpdateOne) SetNillableTime(t *time.Time) *WaterLogUpdateOne {
	if t != nil {
		wluo.SetTime(*t)
	}
	return wluo
}

// Mutation returns the WaterLogMutation object of the builder.
func (wluo *WaterLogUpdateOne) Mutation() *WaterLogMutation {
	return wluo.mutation
}

// Where appends a list predicates to the WaterLogUpdate builder.
func (wluo *WaterLogUpdateOne) Where(ps ...predicate.WaterLog) *WaterLogUpdateOne {
	wluo.mutation.Where(ps...)
	return wluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wluo *WaterLogUpdateOne) Select(field string, fields ...string) *WaterLogUpdateOne {
	wluo.fields = append([]string{field}, fields...)
	return wluo
}

// Save executes the query and returns the updated WaterLog entity.
func (wluo *WaterLogUpdateOne) Save(ctx context.Context) (*WaterLog, error) {
	return withHooks(ctx, wluo.sqlSave, wluo.mutation, wluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wluo *WaterLogUpdateOne) SaveX(ctx context.Context) *WaterLog {
	node, err := wluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wluo *WaterLogUpdateOne) Exec(ctx context.Context) error {
	_, err := wluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wluo *WaterLogUpdateOne) ExecX(ctx context.Context) {
	if err := wluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wluo *WaterLogUpdateOne) sqlSave(ctx context.Context) (_node *WaterLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(waterlog.Table, waterlog.Columns, sqlgraph.NewFieldSpec(waterlog.FieldID, field.TypeInt))
	id, ok := wluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WaterLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, waterlog.FieldID)
		for _, f := range fields {
			if !waterlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != waterlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wluo.mutation.Seconds(); ok {
		_spec.SetField(waterlog.FieldSeconds, field.TypeInt, value)
	}
	if value, ok := wluo.mutation.AddedSeconds(); ok {
		_spec.AddField(waterlog.FieldSeconds, field.TypeInt, value)
	}
	if value, ok := wluo.mutation.Channel(); ok {
		_spec.SetField(waterlog.FieldChannel, field.TypeString, value)
	}
	if value, ok := wluo.mutation.Manual(); ok {
		_spec.SetField(waterlog.FieldManual, field.TypeBool, value)
	}
	if value, ok := wluo.mutation.Time(); ok {
		_spec.SetField(waterlog.FieldTime, field.TypeTime, value)
	}
	_node = &WaterLog{config: wluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{waterlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wluo.mutation.done = true
	return _node, nil
}
