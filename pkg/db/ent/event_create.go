// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	entevent "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/event"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ec *EventCreate) SetCreatedAt(u uint32) *EventCreate {
	ec.mutation.SetCreatedAt(u)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableCreatedAt(u *uint32) *EventCreate {
	if u != nil {
		ec.SetCreatedAt(*u)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EventCreate) SetUpdatedAt(u uint32) *EventCreate {
	ec.mutation.SetUpdatedAt(u)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableUpdatedAt(u *uint32) *EventCreate {
	if u != nil {
		ec.SetUpdatedAt(*u)
	}
	return ec
}

// SetDeletedAt sets the "deleted_at" field.
func (ec *EventCreate) SetDeletedAt(u uint32) *EventCreate {
	ec.mutation.SetDeletedAt(u)
	return ec
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableDeletedAt(u *uint32) *EventCreate {
	if u != nil {
		ec.SetDeletedAt(*u)
	}
	return ec
}

// SetAppID sets the "app_id" field.
func (ec *EventCreate) SetAppID(u uuid.UUID) *EventCreate {
	ec.mutation.SetAppID(u)
	return ec
}

// SetEventType sets the "event_type" field.
func (ec *EventCreate) SetEventType(s string) *EventCreate {
	ec.mutation.SetEventType(s)
	return ec
}

// SetNillableEventType sets the "event_type" field if the given value is not nil.
func (ec *EventCreate) SetNillableEventType(s *string) *EventCreate {
	if s != nil {
		ec.SetEventType(*s)
	}
	return ec
}

// SetCoupons sets the "coupons" field.
func (ec *EventCreate) SetCoupons(e []event.Coupon) *EventCreate {
	ec.mutation.SetCoupons(e)
	return ec
}

// SetCredits sets the "credits" field.
func (ec *EventCreate) SetCredits(d decimal.Decimal) *EventCreate {
	ec.mutation.SetCredits(d)
	return ec
}

// SetNillableCredits sets the "credits" field if the given value is not nil.
func (ec *EventCreate) SetNillableCredits(d *decimal.Decimal) *EventCreate {
	if d != nil {
		ec.SetCredits(*d)
	}
	return ec
}

// SetCreditsPerUsd sets the "credits_per_usd" field.
func (ec *EventCreate) SetCreditsPerUsd(d decimal.Decimal) *EventCreate {
	ec.mutation.SetCreditsPerUsd(d)
	return ec
}

// SetNillableCreditsPerUsd sets the "credits_per_usd" field if the given value is not nil.
func (ec *EventCreate) SetNillableCreditsPerUsd(d *decimal.Decimal) *EventCreate {
	if d != nil {
		ec.SetCreditsPerUsd(*d)
	}
	return ec
}

// SetMaxConsecutive sets the "max_consecutive" field.
func (ec *EventCreate) SetMaxConsecutive(u uint32) *EventCreate {
	ec.mutation.SetMaxConsecutive(u)
	return ec
}

// SetNillableMaxConsecutive sets the "max_consecutive" field if the given value is not nil.
func (ec *EventCreate) SetNillableMaxConsecutive(u *uint32) *EventCreate {
	if u != nil {
		ec.SetMaxConsecutive(*u)
	}
	return ec
}

// SetGoodID sets the "good_id" field.
func (ec *EventCreate) SetGoodID(u uuid.UUID) *EventCreate {
	ec.mutation.SetGoodID(u)
	return ec
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (ec *EventCreate) SetNillableGoodID(u *uuid.UUID) *EventCreate {
	if u != nil {
		ec.SetGoodID(*u)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EventCreate) SetID(u uuid.UUID) *EventCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EventCreate) SetNillableID(u *uuid.UUID) *EventCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if err := ec.defaults(); err != nil {
		return nil, err
	}
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Event)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventCreate) defaults() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		if entevent.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized entevent.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := entevent.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		if entevent.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized entevent.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := entevent.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	if _, ok := ec.mutation.DeletedAt(); !ok {
		if entevent.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized entevent.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := entevent.DefaultDeletedAt()
		ec.mutation.SetDeletedAt(v)
	}
	if _, ok := ec.mutation.EventType(); !ok {
		v := entevent.DefaultEventType
		ec.mutation.SetEventType(v)
	}
	if _, ok := ec.mutation.Coupons(); !ok {
		v := entevent.DefaultCoupons
		ec.mutation.SetCoupons(v)
	}
	if _, ok := ec.mutation.Credits(); !ok {
		v := entevent.DefaultCredits
		ec.mutation.SetCredits(v)
	}
	if _, ok := ec.mutation.CreditsPerUsd(); !ok {
		v := entevent.DefaultCreditsPerUsd
		ec.mutation.SetCreditsPerUsd(v)
	}
	if _, ok := ec.mutation.MaxConsecutive(); !ok {
		v := entevent.DefaultMaxConsecutive
		ec.mutation.SetMaxConsecutive(v)
	}
	if _, ok := ec.mutation.GoodID(); !ok {
		if entevent.DefaultGoodID == nil {
			return fmt.Errorf("ent: uninitialized entevent.DefaultGoodID (forgotten import ent/runtime?)")
		}
		v := entevent.DefaultGoodID()
		ec.mutation.SetGoodID(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		if entevent.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized entevent.DefaultID (forgotten import ent/runtime?)")
		}
		v := entevent.DefaultID()
		ec.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Event.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Event.updated_at"`)}
	}
	if _, ok := ec.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Event.deleted_at"`)}
	}
	if _, ok := ec.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "Event.app_id"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: entevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: entevent.FieldID,
			},
		}
	)
	_spec.OnConflict = ec.conflict
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: entevent.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: entevent.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: entevent.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ec.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: entevent.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := ec.mutation.EventType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entevent.FieldEventType,
		})
		_node.EventType = value
	}
	if value, ok := ec.mutation.Coupons(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: entevent.FieldCoupons,
		})
		_node.Coupons = value
	}
	if value, ok := ec.mutation.Credits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: entevent.FieldCredits,
		})
		_node.Credits = value
	}
	if value, ok := ec.mutation.CreditsPerUsd(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: entevent.FieldCreditsPerUsd,
		})
		_node.CreditsPerUsd = value
	}
	if value, ok := ec.mutation.MaxConsecutive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: entevent.FieldMaxConsecutive,
		})
		_node.MaxConsecutive = value
	}
	if value, ok := ec.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: entevent.FieldGoodID,
		})
		_node.GoodID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Event.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ec *EventCreate) OnConflict(opts ...sql.ConflictOption) *EventUpsertOne {
	ec.conflict = opts
	return &EventUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ec *EventCreate) OnConflictColumns(columns ...string) *EventUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EventUpsertOne{
		create: ec,
	}
}

type (
	// EventUpsertOne is the builder for "upsert"-ing
	//  one Event node.
	EventUpsertOne struct {
		create *EventCreate
	}

	// EventUpsert is the "OnConflict" setter.
	EventUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *EventUpsert) SetCreatedAt(v uint32) *EventUpsert {
	u.Set(entevent.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *EventUpsert) UpdateCreatedAt() *EventUpsert {
	u.SetExcluded(entevent.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *EventUpsert) AddCreatedAt(v uint32) *EventUpsert {
	u.Add(entevent.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EventUpsert) SetUpdatedAt(v uint32) *EventUpsert {
	u.Set(entevent.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EventUpsert) UpdateUpdatedAt() *EventUpsert {
	u.SetExcluded(entevent.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *EventUpsert) AddUpdatedAt(v uint32) *EventUpsert {
	u.Add(entevent.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EventUpsert) SetDeletedAt(v uint32) *EventUpsert {
	u.Set(entevent.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EventUpsert) UpdateDeletedAt() *EventUpsert {
	u.SetExcluded(entevent.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *EventUpsert) AddDeletedAt(v uint32) *EventUpsert {
	u.Add(entevent.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *EventUpsert) SetAppID(v uuid.UUID) *EventUpsert {
	u.Set(entevent.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *EventUpsert) UpdateAppID() *EventUpsert {
	u.SetExcluded(entevent.FieldAppID)
	return u
}

// SetEventType sets the "event_type" field.
func (u *EventUpsert) SetEventType(v string) *EventUpsert {
	u.Set(entevent.FieldEventType, v)
	return u
}

// UpdateEventType sets the "event_type" field to the value that was provided on create.
func (u *EventUpsert) UpdateEventType() *EventUpsert {
	u.SetExcluded(entevent.FieldEventType)
	return u
}

// ClearEventType clears the value of the "event_type" field.
func (u *EventUpsert) ClearEventType() *EventUpsert {
	u.SetNull(entevent.FieldEventType)
	return u
}

// SetCoupons sets the "coupons" field.
func (u *EventUpsert) SetCoupons(v []event.Coupon) *EventUpsert {
	u.Set(entevent.FieldCoupons, v)
	return u
}

// UpdateCoupons sets the "coupons" field to the value that was provided on create.
func (u *EventUpsert) UpdateCoupons() *EventUpsert {
	u.SetExcluded(entevent.FieldCoupons)
	return u
}

// ClearCoupons clears the value of the "coupons" field.
func (u *EventUpsert) ClearCoupons() *EventUpsert {
	u.SetNull(entevent.FieldCoupons)
	return u
}

// SetCredits sets the "credits" field.
func (u *EventUpsert) SetCredits(v decimal.Decimal) *EventUpsert {
	u.Set(entevent.FieldCredits, v)
	return u
}

// UpdateCredits sets the "credits" field to the value that was provided on create.
func (u *EventUpsert) UpdateCredits() *EventUpsert {
	u.SetExcluded(entevent.FieldCredits)
	return u
}

// ClearCredits clears the value of the "credits" field.
func (u *EventUpsert) ClearCredits() *EventUpsert {
	u.SetNull(entevent.FieldCredits)
	return u
}

// SetCreditsPerUsd sets the "credits_per_usd" field.
func (u *EventUpsert) SetCreditsPerUsd(v decimal.Decimal) *EventUpsert {
	u.Set(entevent.FieldCreditsPerUsd, v)
	return u
}

// UpdateCreditsPerUsd sets the "credits_per_usd" field to the value that was provided on create.
func (u *EventUpsert) UpdateCreditsPerUsd() *EventUpsert {
	u.SetExcluded(entevent.FieldCreditsPerUsd)
	return u
}

// ClearCreditsPerUsd clears the value of the "credits_per_usd" field.
func (u *EventUpsert) ClearCreditsPerUsd() *EventUpsert {
	u.SetNull(entevent.FieldCreditsPerUsd)
	return u
}

// SetMaxConsecutive sets the "max_consecutive" field.
func (u *EventUpsert) SetMaxConsecutive(v uint32) *EventUpsert {
	u.Set(entevent.FieldMaxConsecutive, v)
	return u
}

// UpdateMaxConsecutive sets the "max_consecutive" field to the value that was provided on create.
func (u *EventUpsert) UpdateMaxConsecutive() *EventUpsert {
	u.SetExcluded(entevent.FieldMaxConsecutive)
	return u
}

// AddMaxConsecutive adds v to the "max_consecutive" field.
func (u *EventUpsert) AddMaxConsecutive(v uint32) *EventUpsert {
	u.Add(entevent.FieldMaxConsecutive, v)
	return u
}

// ClearMaxConsecutive clears the value of the "max_consecutive" field.
func (u *EventUpsert) ClearMaxConsecutive() *EventUpsert {
	u.SetNull(entevent.FieldMaxConsecutive)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *EventUpsert) SetGoodID(v uuid.UUID) *EventUpsert {
	u.Set(entevent.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *EventUpsert) UpdateGoodID() *EventUpsert {
	u.SetExcluded(entevent.FieldGoodID)
	return u
}

// ClearGoodID clears the value of the "good_id" field.
func (u *EventUpsert) ClearGoodID() *EventUpsert {
	u.SetNull(entevent.FieldGoodID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(entevent.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *EventUpsertOne) UpdateNewValues() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(entevent.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Event.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *EventUpsertOne) Ignore() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventUpsertOne) DoNothing() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventCreate.OnConflict
// documentation for more info.
func (u *EventUpsertOne) Update(set func(*EventUpsert)) *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *EventUpsertOne) SetCreatedAt(v uint32) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *EventUpsertOne) AddCreatedAt(v uint32) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateCreatedAt() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EventUpsertOne) SetUpdatedAt(v uint32) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *EventUpsertOne) AddUpdatedAt(v uint32) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateUpdatedAt() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EventUpsertOne) SetDeletedAt(v uint32) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *EventUpsertOne) AddDeletedAt(v uint32) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateDeletedAt() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *EventUpsertOne) SetAppID(v uuid.UUID) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateAppID() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateAppID()
	})
}

// SetEventType sets the "event_type" field.
func (u *EventUpsertOne) SetEventType(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetEventType(v)
	})
}

// UpdateEventType sets the "event_type" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateEventType() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateEventType()
	})
}

// ClearEventType clears the value of the "event_type" field.
func (u *EventUpsertOne) ClearEventType() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearEventType()
	})
}

// SetCoupons sets the "coupons" field.
func (u *EventUpsertOne) SetCoupons(v []event.Coupon) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetCoupons(v)
	})
}

// UpdateCoupons sets the "coupons" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateCoupons() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateCoupons()
	})
}

// ClearCoupons clears the value of the "coupons" field.
func (u *EventUpsertOne) ClearCoupons() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearCoupons()
	})
}

// SetCredits sets the "credits" field.
func (u *EventUpsertOne) SetCredits(v decimal.Decimal) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetCredits(v)
	})
}

// UpdateCredits sets the "credits" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateCredits() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateCredits()
	})
}

// ClearCredits clears the value of the "credits" field.
func (u *EventUpsertOne) ClearCredits() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearCredits()
	})
}

// SetCreditsPerUsd sets the "credits_per_usd" field.
func (u *EventUpsertOne) SetCreditsPerUsd(v decimal.Decimal) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetCreditsPerUsd(v)
	})
}

// UpdateCreditsPerUsd sets the "credits_per_usd" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateCreditsPerUsd() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateCreditsPerUsd()
	})
}

// ClearCreditsPerUsd clears the value of the "credits_per_usd" field.
func (u *EventUpsertOne) ClearCreditsPerUsd() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearCreditsPerUsd()
	})
}

// SetMaxConsecutive sets the "max_consecutive" field.
func (u *EventUpsertOne) SetMaxConsecutive(v uint32) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetMaxConsecutive(v)
	})
}

// AddMaxConsecutive adds v to the "max_consecutive" field.
func (u *EventUpsertOne) AddMaxConsecutive(v uint32) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.AddMaxConsecutive(v)
	})
}

// UpdateMaxConsecutive sets the "max_consecutive" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateMaxConsecutive() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateMaxConsecutive()
	})
}

// ClearMaxConsecutive clears the value of the "max_consecutive" field.
func (u *EventUpsertOne) ClearMaxConsecutive() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearMaxConsecutive()
	})
}

// SetGoodID sets the "good_id" field.
func (u *EventUpsertOne) SetGoodID(v uuid.UUID) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateGoodID() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *EventUpsertOne) ClearGoodID() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearGoodID()
	})
}

// Exec executes the query.
func (u *EventUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EventUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: EventUpsertOne.ID is not supported by MySQL driver. Use EventUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EventUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	builders []*EventCreate
	conflict []sql.ConflictOption
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Event.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ecb *EventCreateBulk) OnConflict(opts ...sql.ConflictOption) *EventUpsertBulk {
	ecb.conflict = opts
	return &EventUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ecb *EventCreateBulk) OnConflictColumns(columns ...string) *EventUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EventUpsertBulk{
		create: ecb,
	}
}

// EventUpsertBulk is the builder for "upsert"-ing
// a bulk of Event nodes.
type EventUpsertBulk struct {
	create *EventCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(entevent.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *EventUpsertBulk) UpdateNewValues() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(entevent.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *EventUpsertBulk) Ignore() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventUpsertBulk) DoNothing() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventCreateBulk.OnConflict
// documentation for more info.
func (u *EventUpsertBulk) Update(set func(*EventUpsert)) *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *EventUpsertBulk) SetCreatedAt(v uint32) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *EventUpsertBulk) AddCreatedAt(v uint32) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateCreatedAt() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EventUpsertBulk) SetUpdatedAt(v uint32) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *EventUpsertBulk) AddUpdatedAt(v uint32) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateUpdatedAt() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EventUpsertBulk) SetDeletedAt(v uint32) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *EventUpsertBulk) AddDeletedAt(v uint32) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateDeletedAt() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *EventUpsertBulk) SetAppID(v uuid.UUID) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateAppID() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateAppID()
	})
}

// SetEventType sets the "event_type" field.
func (u *EventUpsertBulk) SetEventType(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetEventType(v)
	})
}

// UpdateEventType sets the "event_type" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateEventType() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateEventType()
	})
}

// ClearEventType clears the value of the "event_type" field.
func (u *EventUpsertBulk) ClearEventType() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearEventType()
	})
}

// SetCoupons sets the "coupons" field.
func (u *EventUpsertBulk) SetCoupons(v []event.Coupon) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetCoupons(v)
	})
}

// UpdateCoupons sets the "coupons" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateCoupons() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateCoupons()
	})
}

// ClearCoupons clears the value of the "coupons" field.
func (u *EventUpsertBulk) ClearCoupons() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearCoupons()
	})
}

// SetCredits sets the "credits" field.
func (u *EventUpsertBulk) SetCredits(v decimal.Decimal) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetCredits(v)
	})
}

// UpdateCredits sets the "credits" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateCredits() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateCredits()
	})
}

// ClearCredits clears the value of the "credits" field.
func (u *EventUpsertBulk) ClearCredits() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearCredits()
	})
}

// SetCreditsPerUsd sets the "credits_per_usd" field.
func (u *EventUpsertBulk) SetCreditsPerUsd(v decimal.Decimal) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetCreditsPerUsd(v)
	})
}

// UpdateCreditsPerUsd sets the "credits_per_usd" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateCreditsPerUsd() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateCreditsPerUsd()
	})
}

// ClearCreditsPerUsd clears the value of the "credits_per_usd" field.
func (u *EventUpsertBulk) ClearCreditsPerUsd() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearCreditsPerUsd()
	})
}

// SetMaxConsecutive sets the "max_consecutive" field.
func (u *EventUpsertBulk) SetMaxConsecutive(v uint32) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetMaxConsecutive(v)
	})
}

// AddMaxConsecutive adds v to the "max_consecutive" field.
func (u *EventUpsertBulk) AddMaxConsecutive(v uint32) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.AddMaxConsecutive(v)
	})
}

// UpdateMaxConsecutive sets the "max_consecutive" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateMaxConsecutive() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateMaxConsecutive()
	})
}

// ClearMaxConsecutive clears the value of the "max_consecutive" field.
func (u *EventUpsertBulk) ClearMaxConsecutive() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearMaxConsecutive()
	})
}

// SetGoodID sets the "good_id" field.
func (u *EventUpsertBulk) SetGoodID(v uuid.UUID) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateGoodID() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *EventUpsertBulk) ClearGoodID() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearGoodID()
	})
}

// Exec executes the query.
func (u *EventUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EventCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
