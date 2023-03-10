// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponfixamount"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CouponFixAmountUpdate is the builder for updating CouponFixAmount entities.
type CouponFixAmountUpdate struct {
	config
	hooks     []Hook
	mutation  *CouponFixAmountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CouponFixAmountUpdate builder.
func (cfau *CouponFixAmountUpdate) Where(ps ...predicate.CouponFixAmount) *CouponFixAmountUpdate {
	cfau.mutation.Where(ps...)
	return cfau
}

// SetCreatedAt sets the "created_at" field.
func (cfau *CouponFixAmountUpdate) SetCreatedAt(u uint32) *CouponFixAmountUpdate {
	cfau.mutation.ResetCreatedAt()
	cfau.mutation.SetCreatedAt(u)
	return cfau
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cfau *CouponFixAmountUpdate) SetNillableCreatedAt(u *uint32) *CouponFixAmountUpdate {
	if u != nil {
		cfau.SetCreatedAt(*u)
	}
	return cfau
}

// AddCreatedAt adds u to the "created_at" field.
func (cfau *CouponFixAmountUpdate) AddCreatedAt(u int32) *CouponFixAmountUpdate {
	cfau.mutation.AddCreatedAt(u)
	return cfau
}

// SetUpdatedAt sets the "updated_at" field.
func (cfau *CouponFixAmountUpdate) SetUpdatedAt(u uint32) *CouponFixAmountUpdate {
	cfau.mutation.ResetUpdatedAt()
	cfau.mutation.SetUpdatedAt(u)
	return cfau
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cfau *CouponFixAmountUpdate) AddUpdatedAt(u int32) *CouponFixAmountUpdate {
	cfau.mutation.AddUpdatedAt(u)
	return cfau
}

// SetDeletedAt sets the "deleted_at" field.
func (cfau *CouponFixAmountUpdate) SetDeletedAt(u uint32) *CouponFixAmountUpdate {
	cfau.mutation.ResetDeletedAt()
	cfau.mutation.SetDeletedAt(u)
	return cfau
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cfau *CouponFixAmountUpdate) SetNillableDeletedAt(u *uint32) *CouponFixAmountUpdate {
	if u != nil {
		cfau.SetDeletedAt(*u)
	}
	return cfau
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cfau *CouponFixAmountUpdate) AddDeletedAt(u int32) *CouponFixAmountUpdate {
	cfau.mutation.AddDeletedAt(u)
	return cfau
}

// SetAppID sets the "app_id" field.
func (cfau *CouponFixAmountUpdate) SetAppID(u uuid.UUID) *CouponFixAmountUpdate {
	cfau.mutation.SetAppID(u)
	return cfau
}

// SetDenomination sets the "denomination" field.
func (cfau *CouponFixAmountUpdate) SetDenomination(d decimal.Decimal) *CouponFixAmountUpdate {
	cfau.mutation.SetDenomination(d)
	return cfau
}

// SetNillableDenomination sets the "denomination" field if the given value is not nil.
func (cfau *CouponFixAmountUpdate) SetNillableDenomination(d *decimal.Decimal) *CouponFixAmountUpdate {
	if d != nil {
		cfau.SetDenomination(*d)
	}
	return cfau
}

// ClearDenomination clears the value of the "denomination" field.
func (cfau *CouponFixAmountUpdate) ClearDenomination() *CouponFixAmountUpdate {
	cfau.mutation.ClearDenomination()
	return cfau
}

// SetCirculation sets the "circulation" field.
func (cfau *CouponFixAmountUpdate) SetCirculation(d decimal.Decimal) *CouponFixAmountUpdate {
	cfau.mutation.SetCirculation(d)
	return cfau
}

// SetNillableCirculation sets the "circulation" field if the given value is not nil.
func (cfau *CouponFixAmountUpdate) SetNillableCirculation(d *decimal.Decimal) *CouponFixAmountUpdate {
	if d != nil {
		cfau.SetCirculation(*d)
	}
	return cfau
}

// ClearCirculation clears the value of the "circulation" field.
func (cfau *CouponFixAmountUpdate) ClearCirculation() *CouponFixAmountUpdate {
	cfau.mutation.ClearCirculation()
	return cfau
}

// SetReleasedByUserID sets the "released_by_user_id" field.
func (cfau *CouponFixAmountUpdate) SetReleasedByUserID(u uuid.UUID) *CouponFixAmountUpdate {
	cfau.mutation.SetReleasedByUserID(u)
	return cfau
}

// SetStartAt sets the "start_at" field.
func (cfau *CouponFixAmountUpdate) SetStartAt(u uint32) *CouponFixAmountUpdate {
	cfau.mutation.ResetStartAt()
	cfau.mutation.SetStartAt(u)
	return cfau
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cfau *CouponFixAmountUpdate) SetNillableStartAt(u *uint32) *CouponFixAmountUpdate {
	if u != nil {
		cfau.SetStartAt(*u)
	}
	return cfau
}

// AddStartAt adds u to the "start_at" field.
func (cfau *CouponFixAmountUpdate) AddStartAt(u int32) *CouponFixAmountUpdate {
	cfau.mutation.AddStartAt(u)
	return cfau
}

// ClearStartAt clears the value of the "start_at" field.
func (cfau *CouponFixAmountUpdate) ClearStartAt() *CouponFixAmountUpdate {
	cfau.mutation.ClearStartAt()
	return cfau
}

// SetDurationDays sets the "duration_days" field.
func (cfau *CouponFixAmountUpdate) SetDurationDays(u uint32) *CouponFixAmountUpdate {
	cfau.mutation.ResetDurationDays()
	cfau.mutation.SetDurationDays(u)
	return cfau
}

// SetNillableDurationDays sets the "duration_days" field if the given value is not nil.
func (cfau *CouponFixAmountUpdate) SetNillableDurationDays(u *uint32) *CouponFixAmountUpdate {
	if u != nil {
		cfau.SetDurationDays(*u)
	}
	return cfau
}

// AddDurationDays adds u to the "duration_days" field.
func (cfau *CouponFixAmountUpdate) AddDurationDays(u int32) *CouponFixAmountUpdate {
	cfau.mutation.AddDurationDays(u)
	return cfau
}

// ClearDurationDays clears the value of the "duration_days" field.
func (cfau *CouponFixAmountUpdate) ClearDurationDays() *CouponFixAmountUpdate {
	cfau.mutation.ClearDurationDays()
	return cfau
}

// SetMessage sets the "message" field.
func (cfau *CouponFixAmountUpdate) SetMessage(s string) *CouponFixAmountUpdate {
	cfau.mutation.SetMessage(s)
	return cfau
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cfau *CouponFixAmountUpdate) SetNillableMessage(s *string) *CouponFixAmountUpdate {
	if s != nil {
		cfau.SetMessage(*s)
	}
	return cfau
}

// ClearMessage clears the value of the "message" field.
func (cfau *CouponFixAmountUpdate) ClearMessage() *CouponFixAmountUpdate {
	cfau.mutation.ClearMessage()
	return cfau
}

// SetName sets the "name" field.
func (cfau *CouponFixAmountUpdate) SetName(s string) *CouponFixAmountUpdate {
	cfau.mutation.SetName(s)
	return cfau
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cfau *CouponFixAmountUpdate) SetNillableName(s *string) *CouponFixAmountUpdate {
	if s != nil {
		cfau.SetName(*s)
	}
	return cfau
}

// ClearName clears the value of the "name" field.
func (cfau *CouponFixAmountUpdate) ClearName() *CouponFixAmountUpdate {
	cfau.mutation.ClearName()
	return cfau
}

// SetAllocated sets the "allocated" field.
func (cfau *CouponFixAmountUpdate) SetAllocated(u uint32) *CouponFixAmountUpdate {
	cfau.mutation.ResetAllocated()
	cfau.mutation.SetAllocated(u)
	return cfau
}

// SetNillableAllocated sets the "allocated" field if the given value is not nil.
func (cfau *CouponFixAmountUpdate) SetNillableAllocated(u *uint32) *CouponFixAmountUpdate {
	if u != nil {
		cfau.SetAllocated(*u)
	}
	return cfau
}

// AddAllocated adds u to the "allocated" field.
func (cfau *CouponFixAmountUpdate) AddAllocated(u int32) *CouponFixAmountUpdate {
	cfau.mutation.AddAllocated(u)
	return cfau
}

// ClearAllocated clears the value of the "allocated" field.
func (cfau *CouponFixAmountUpdate) ClearAllocated() *CouponFixAmountUpdate {
	cfau.mutation.ClearAllocated()
	return cfau
}

// Mutation returns the CouponFixAmountMutation object of the builder.
func (cfau *CouponFixAmountUpdate) Mutation() *CouponFixAmountMutation {
	return cfau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cfau *CouponFixAmountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cfau.defaults(); err != nil {
		return 0, err
	}
	if len(cfau.hooks) == 0 {
		affected, err = cfau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponFixAmountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cfau.mutation = mutation
			affected, err = cfau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cfau.hooks) - 1; i >= 0; i-- {
			if cfau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cfau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cfau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cfau *CouponFixAmountUpdate) SaveX(ctx context.Context) int {
	affected, err := cfau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cfau *CouponFixAmountUpdate) Exec(ctx context.Context) error {
	_, err := cfau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfau *CouponFixAmountUpdate) ExecX(ctx context.Context) {
	if err := cfau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cfau *CouponFixAmountUpdate) defaults() error {
	if _, ok := cfau.mutation.UpdatedAt(); !ok {
		if couponfixamount.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized couponfixamount.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := couponfixamount.UpdateDefaultUpdatedAt()
		cfau.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cfau *CouponFixAmountUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CouponFixAmountUpdate {
	cfau.modifiers = append(cfau.modifiers, modifiers...)
	return cfau
}

func (cfau *CouponFixAmountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponfixamount.Table,
			Columns: couponfixamount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponfixamount.FieldID,
			},
		},
	}
	if ps := cfau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cfau.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldCreatedAt,
		})
	}
	if value, ok := cfau.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldCreatedAt,
		})
	}
	if value, ok := cfau.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldUpdatedAt,
		})
	}
	if value, ok := cfau.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldUpdatedAt,
		})
	}
	if value, ok := cfau.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldDeletedAt,
		})
	}
	if value, ok := cfau.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldDeletedAt,
		})
	}
	if value, ok := cfau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponfixamount.FieldAppID,
		})
	}
	if value, ok := cfau.mutation.Denomination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: couponfixamount.FieldDenomination,
		})
	}
	if cfau.mutation.DenominationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: couponfixamount.FieldDenomination,
		})
	}
	if value, ok := cfau.mutation.Circulation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: couponfixamount.FieldCirculation,
		})
	}
	if cfau.mutation.CirculationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: couponfixamount.FieldCirculation,
		})
	}
	if value, ok := cfau.mutation.ReleasedByUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponfixamount.FieldReleasedByUserID,
		})
	}
	if value, ok := cfau.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldStartAt,
		})
	}
	if value, ok := cfau.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldStartAt,
		})
	}
	if cfau.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponfixamount.FieldStartAt,
		})
	}
	if value, ok := cfau.mutation.DurationDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldDurationDays,
		})
	}
	if value, ok := cfau.mutation.AddedDurationDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldDurationDays,
		})
	}
	if cfau.mutation.DurationDaysCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponfixamount.FieldDurationDays,
		})
	}
	if value, ok := cfau.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponfixamount.FieldMessage,
		})
	}
	if cfau.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: couponfixamount.FieldMessage,
		})
	}
	if value, ok := cfau.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponfixamount.FieldName,
		})
	}
	if cfau.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: couponfixamount.FieldName,
		})
	}
	if value, ok := cfau.mutation.Allocated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldAllocated,
		})
	}
	if value, ok := cfau.mutation.AddedAllocated(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldAllocated,
		})
	}
	if cfau.mutation.AllocatedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponfixamount.FieldAllocated,
		})
	}
	_spec.Modifiers = cfau.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cfau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponfixamount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CouponFixAmountUpdateOne is the builder for updating a single CouponFixAmount entity.
type CouponFixAmountUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CouponFixAmountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cfauo *CouponFixAmountUpdateOne) SetCreatedAt(u uint32) *CouponFixAmountUpdateOne {
	cfauo.mutation.ResetCreatedAt()
	cfauo.mutation.SetCreatedAt(u)
	return cfauo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cfauo *CouponFixAmountUpdateOne) SetNillableCreatedAt(u *uint32) *CouponFixAmountUpdateOne {
	if u != nil {
		cfauo.SetCreatedAt(*u)
	}
	return cfauo
}

// AddCreatedAt adds u to the "created_at" field.
func (cfauo *CouponFixAmountUpdateOne) AddCreatedAt(u int32) *CouponFixAmountUpdateOne {
	cfauo.mutation.AddCreatedAt(u)
	return cfauo
}

// SetUpdatedAt sets the "updated_at" field.
func (cfauo *CouponFixAmountUpdateOne) SetUpdatedAt(u uint32) *CouponFixAmountUpdateOne {
	cfauo.mutation.ResetUpdatedAt()
	cfauo.mutation.SetUpdatedAt(u)
	return cfauo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cfauo *CouponFixAmountUpdateOne) AddUpdatedAt(u int32) *CouponFixAmountUpdateOne {
	cfauo.mutation.AddUpdatedAt(u)
	return cfauo
}

// SetDeletedAt sets the "deleted_at" field.
func (cfauo *CouponFixAmountUpdateOne) SetDeletedAt(u uint32) *CouponFixAmountUpdateOne {
	cfauo.mutation.ResetDeletedAt()
	cfauo.mutation.SetDeletedAt(u)
	return cfauo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cfauo *CouponFixAmountUpdateOne) SetNillableDeletedAt(u *uint32) *CouponFixAmountUpdateOne {
	if u != nil {
		cfauo.SetDeletedAt(*u)
	}
	return cfauo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cfauo *CouponFixAmountUpdateOne) AddDeletedAt(u int32) *CouponFixAmountUpdateOne {
	cfauo.mutation.AddDeletedAt(u)
	return cfauo
}

// SetAppID sets the "app_id" field.
func (cfauo *CouponFixAmountUpdateOne) SetAppID(u uuid.UUID) *CouponFixAmountUpdateOne {
	cfauo.mutation.SetAppID(u)
	return cfauo
}

// SetDenomination sets the "denomination" field.
func (cfauo *CouponFixAmountUpdateOne) SetDenomination(d decimal.Decimal) *CouponFixAmountUpdateOne {
	cfauo.mutation.SetDenomination(d)
	return cfauo
}

// SetNillableDenomination sets the "denomination" field if the given value is not nil.
func (cfauo *CouponFixAmountUpdateOne) SetNillableDenomination(d *decimal.Decimal) *CouponFixAmountUpdateOne {
	if d != nil {
		cfauo.SetDenomination(*d)
	}
	return cfauo
}

// ClearDenomination clears the value of the "denomination" field.
func (cfauo *CouponFixAmountUpdateOne) ClearDenomination() *CouponFixAmountUpdateOne {
	cfauo.mutation.ClearDenomination()
	return cfauo
}

// SetCirculation sets the "circulation" field.
func (cfauo *CouponFixAmountUpdateOne) SetCirculation(d decimal.Decimal) *CouponFixAmountUpdateOne {
	cfauo.mutation.SetCirculation(d)
	return cfauo
}

// SetNillableCirculation sets the "circulation" field if the given value is not nil.
func (cfauo *CouponFixAmountUpdateOne) SetNillableCirculation(d *decimal.Decimal) *CouponFixAmountUpdateOne {
	if d != nil {
		cfauo.SetCirculation(*d)
	}
	return cfauo
}

// ClearCirculation clears the value of the "circulation" field.
func (cfauo *CouponFixAmountUpdateOne) ClearCirculation() *CouponFixAmountUpdateOne {
	cfauo.mutation.ClearCirculation()
	return cfauo
}

// SetReleasedByUserID sets the "released_by_user_id" field.
func (cfauo *CouponFixAmountUpdateOne) SetReleasedByUserID(u uuid.UUID) *CouponFixAmountUpdateOne {
	cfauo.mutation.SetReleasedByUserID(u)
	return cfauo
}

// SetStartAt sets the "start_at" field.
func (cfauo *CouponFixAmountUpdateOne) SetStartAt(u uint32) *CouponFixAmountUpdateOne {
	cfauo.mutation.ResetStartAt()
	cfauo.mutation.SetStartAt(u)
	return cfauo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cfauo *CouponFixAmountUpdateOne) SetNillableStartAt(u *uint32) *CouponFixAmountUpdateOne {
	if u != nil {
		cfauo.SetStartAt(*u)
	}
	return cfauo
}

// AddStartAt adds u to the "start_at" field.
func (cfauo *CouponFixAmountUpdateOne) AddStartAt(u int32) *CouponFixAmountUpdateOne {
	cfauo.mutation.AddStartAt(u)
	return cfauo
}

// ClearStartAt clears the value of the "start_at" field.
func (cfauo *CouponFixAmountUpdateOne) ClearStartAt() *CouponFixAmountUpdateOne {
	cfauo.mutation.ClearStartAt()
	return cfauo
}

// SetDurationDays sets the "duration_days" field.
func (cfauo *CouponFixAmountUpdateOne) SetDurationDays(u uint32) *CouponFixAmountUpdateOne {
	cfauo.mutation.ResetDurationDays()
	cfauo.mutation.SetDurationDays(u)
	return cfauo
}

// SetNillableDurationDays sets the "duration_days" field if the given value is not nil.
func (cfauo *CouponFixAmountUpdateOne) SetNillableDurationDays(u *uint32) *CouponFixAmountUpdateOne {
	if u != nil {
		cfauo.SetDurationDays(*u)
	}
	return cfauo
}

// AddDurationDays adds u to the "duration_days" field.
func (cfauo *CouponFixAmountUpdateOne) AddDurationDays(u int32) *CouponFixAmountUpdateOne {
	cfauo.mutation.AddDurationDays(u)
	return cfauo
}

// ClearDurationDays clears the value of the "duration_days" field.
func (cfauo *CouponFixAmountUpdateOne) ClearDurationDays() *CouponFixAmountUpdateOne {
	cfauo.mutation.ClearDurationDays()
	return cfauo
}

// SetMessage sets the "message" field.
func (cfauo *CouponFixAmountUpdateOne) SetMessage(s string) *CouponFixAmountUpdateOne {
	cfauo.mutation.SetMessage(s)
	return cfauo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cfauo *CouponFixAmountUpdateOne) SetNillableMessage(s *string) *CouponFixAmountUpdateOne {
	if s != nil {
		cfauo.SetMessage(*s)
	}
	return cfauo
}

// ClearMessage clears the value of the "message" field.
func (cfauo *CouponFixAmountUpdateOne) ClearMessage() *CouponFixAmountUpdateOne {
	cfauo.mutation.ClearMessage()
	return cfauo
}

// SetName sets the "name" field.
func (cfauo *CouponFixAmountUpdateOne) SetName(s string) *CouponFixAmountUpdateOne {
	cfauo.mutation.SetName(s)
	return cfauo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cfauo *CouponFixAmountUpdateOne) SetNillableName(s *string) *CouponFixAmountUpdateOne {
	if s != nil {
		cfauo.SetName(*s)
	}
	return cfauo
}

// ClearName clears the value of the "name" field.
func (cfauo *CouponFixAmountUpdateOne) ClearName() *CouponFixAmountUpdateOne {
	cfauo.mutation.ClearName()
	return cfauo
}

// SetAllocated sets the "allocated" field.
func (cfauo *CouponFixAmountUpdateOne) SetAllocated(u uint32) *CouponFixAmountUpdateOne {
	cfauo.mutation.ResetAllocated()
	cfauo.mutation.SetAllocated(u)
	return cfauo
}

// SetNillableAllocated sets the "allocated" field if the given value is not nil.
func (cfauo *CouponFixAmountUpdateOne) SetNillableAllocated(u *uint32) *CouponFixAmountUpdateOne {
	if u != nil {
		cfauo.SetAllocated(*u)
	}
	return cfauo
}

// AddAllocated adds u to the "allocated" field.
func (cfauo *CouponFixAmountUpdateOne) AddAllocated(u int32) *CouponFixAmountUpdateOne {
	cfauo.mutation.AddAllocated(u)
	return cfauo
}

// ClearAllocated clears the value of the "allocated" field.
func (cfauo *CouponFixAmountUpdateOne) ClearAllocated() *CouponFixAmountUpdateOne {
	cfauo.mutation.ClearAllocated()
	return cfauo
}

// Mutation returns the CouponFixAmountMutation object of the builder.
func (cfauo *CouponFixAmountUpdateOne) Mutation() *CouponFixAmountMutation {
	return cfauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cfauo *CouponFixAmountUpdateOne) Select(field string, fields ...string) *CouponFixAmountUpdateOne {
	cfauo.fields = append([]string{field}, fields...)
	return cfauo
}

// Save executes the query and returns the updated CouponFixAmount entity.
func (cfauo *CouponFixAmountUpdateOne) Save(ctx context.Context) (*CouponFixAmount, error) {
	var (
		err  error
		node *CouponFixAmount
	)
	if err := cfauo.defaults(); err != nil {
		return nil, err
	}
	if len(cfauo.hooks) == 0 {
		node, err = cfauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CouponFixAmountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cfauo.mutation = mutation
			node, err = cfauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cfauo.hooks) - 1; i >= 0; i-- {
			if cfauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cfauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cfauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CouponFixAmount)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CouponFixAmountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cfauo *CouponFixAmountUpdateOne) SaveX(ctx context.Context) *CouponFixAmount {
	node, err := cfauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cfauo *CouponFixAmountUpdateOne) Exec(ctx context.Context) error {
	_, err := cfauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfauo *CouponFixAmountUpdateOne) ExecX(ctx context.Context) {
	if err := cfauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cfauo *CouponFixAmountUpdateOne) defaults() error {
	if _, ok := cfauo.mutation.UpdatedAt(); !ok {
		if couponfixamount.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized couponfixamount.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := couponfixamount.UpdateDefaultUpdatedAt()
		cfauo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cfauo *CouponFixAmountUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CouponFixAmountUpdateOne {
	cfauo.modifiers = append(cfauo.modifiers, modifiers...)
	return cfauo
}

func (cfauo *CouponFixAmountUpdateOne) sqlSave(ctx context.Context) (_node *CouponFixAmount, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   couponfixamount.Table,
			Columns: couponfixamount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: couponfixamount.FieldID,
			},
		},
	}
	id, ok := cfauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CouponFixAmount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cfauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, couponfixamount.FieldID)
		for _, f := range fields {
			if !couponfixamount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != couponfixamount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cfauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cfauo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldCreatedAt,
		})
	}
	if value, ok := cfauo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldCreatedAt,
		})
	}
	if value, ok := cfauo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldUpdatedAt,
		})
	}
	if value, ok := cfauo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldUpdatedAt,
		})
	}
	if value, ok := cfauo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldDeletedAt,
		})
	}
	if value, ok := cfauo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldDeletedAt,
		})
	}
	if value, ok := cfauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponfixamount.FieldAppID,
		})
	}
	if value, ok := cfauo.mutation.Denomination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: couponfixamount.FieldDenomination,
		})
	}
	if cfauo.mutation.DenominationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: couponfixamount.FieldDenomination,
		})
	}
	if value, ok := cfauo.mutation.Circulation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: couponfixamount.FieldCirculation,
		})
	}
	if cfauo.mutation.CirculationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: couponfixamount.FieldCirculation,
		})
	}
	if value, ok := cfauo.mutation.ReleasedByUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: couponfixamount.FieldReleasedByUserID,
		})
	}
	if value, ok := cfauo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldStartAt,
		})
	}
	if value, ok := cfauo.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldStartAt,
		})
	}
	if cfauo.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponfixamount.FieldStartAt,
		})
	}
	if value, ok := cfauo.mutation.DurationDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldDurationDays,
		})
	}
	if value, ok := cfauo.mutation.AddedDurationDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldDurationDays,
		})
	}
	if cfauo.mutation.DurationDaysCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponfixamount.FieldDurationDays,
		})
	}
	if value, ok := cfauo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponfixamount.FieldMessage,
		})
	}
	if cfauo.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: couponfixamount.FieldMessage,
		})
	}
	if value, ok := cfauo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: couponfixamount.FieldName,
		})
	}
	if cfauo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: couponfixamount.FieldName,
		})
	}
	if value, ok := cfauo.mutation.Allocated(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldAllocated,
		})
	}
	if value, ok := cfauo.mutation.AddedAllocated(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: couponfixamount.FieldAllocated,
		})
	}
	if cfauo.mutation.AllocatedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: couponfixamount.FieldAllocated,
		})
	}
	_spec.Modifiers = cfauo.modifiers
	_node = &CouponFixAmount{config: cfauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cfauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{couponfixamount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
