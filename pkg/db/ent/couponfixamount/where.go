// Code generated by ent, DO NOT EDIT.

package couponfixamount

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// Denomination applies equality check predicate on the "denomination" field. It's identical to DenominationEQ.
func Denomination(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDenomination), v))
	})
}

// Circulation applies equality check predicate on the "circulation" field. It's identical to CirculationEQ.
func Circulation(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCirculation), v))
	})
}

// ReleasedByUserID applies equality check predicate on the "released_by_user_id" field. It's identical to ReleasedByUserIDEQ.
func ReleasedByUserID(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleasedByUserID), v))
	})
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// DurationDays applies equality check predicate on the "duration_days" field. It's identical to DurationDaysEQ.
func DurationDays(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDurationDays), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Allocated applies equality check predicate on the "allocated" field. It's identical to AllocatedEQ.
func Allocated(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllocated), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// DenominationEQ applies the EQ predicate on the "denomination" field.
func DenominationEQ(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDenomination), v))
	})
}

// DenominationNEQ applies the NEQ predicate on the "denomination" field.
func DenominationNEQ(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDenomination), v))
	})
}

// DenominationIn applies the In predicate on the "denomination" field.
func DenominationIn(vs ...decimal.Decimal) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDenomination), v...))
	})
}

// DenominationNotIn applies the NotIn predicate on the "denomination" field.
func DenominationNotIn(vs ...decimal.Decimal) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDenomination), v...))
	})
}

// DenominationGT applies the GT predicate on the "denomination" field.
func DenominationGT(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDenomination), v))
	})
}

// DenominationGTE applies the GTE predicate on the "denomination" field.
func DenominationGTE(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDenomination), v))
	})
}

// DenominationLT applies the LT predicate on the "denomination" field.
func DenominationLT(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDenomination), v))
	})
}

// DenominationLTE applies the LTE predicate on the "denomination" field.
func DenominationLTE(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDenomination), v))
	})
}

// DenominationIsNil applies the IsNil predicate on the "denomination" field.
func DenominationIsNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDenomination)))
	})
}

// DenominationNotNil applies the NotNil predicate on the "denomination" field.
func DenominationNotNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDenomination)))
	})
}

// CirculationEQ applies the EQ predicate on the "circulation" field.
func CirculationEQ(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCirculation), v))
	})
}

// CirculationNEQ applies the NEQ predicate on the "circulation" field.
func CirculationNEQ(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCirculation), v))
	})
}

// CirculationIn applies the In predicate on the "circulation" field.
func CirculationIn(vs ...decimal.Decimal) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCirculation), v...))
	})
}

// CirculationNotIn applies the NotIn predicate on the "circulation" field.
func CirculationNotIn(vs ...decimal.Decimal) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCirculation), v...))
	})
}

// CirculationGT applies the GT predicate on the "circulation" field.
func CirculationGT(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCirculation), v))
	})
}

// CirculationGTE applies the GTE predicate on the "circulation" field.
func CirculationGTE(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCirculation), v))
	})
}

// CirculationLT applies the LT predicate on the "circulation" field.
func CirculationLT(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCirculation), v))
	})
}

// CirculationLTE applies the LTE predicate on the "circulation" field.
func CirculationLTE(v decimal.Decimal) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCirculation), v))
	})
}

// CirculationIsNil applies the IsNil predicate on the "circulation" field.
func CirculationIsNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCirculation)))
	})
}

// CirculationNotNil applies the NotNil predicate on the "circulation" field.
func CirculationNotNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCirculation)))
	})
}

// ReleasedByUserIDEQ applies the EQ predicate on the "released_by_user_id" field.
func ReleasedByUserIDEQ(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDNEQ applies the NEQ predicate on the "released_by_user_id" field.
func ReleasedByUserIDNEQ(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDIn applies the In predicate on the "released_by_user_id" field.
func ReleasedByUserIDIn(vs ...uuid.UUID) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReleasedByUserID), v...))
	})
}

// ReleasedByUserIDNotIn applies the NotIn predicate on the "released_by_user_id" field.
func ReleasedByUserIDNotIn(vs ...uuid.UUID) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReleasedByUserID), v...))
	})
}

// ReleasedByUserIDGT applies the GT predicate on the "released_by_user_id" field.
func ReleasedByUserIDGT(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDGTE applies the GTE predicate on the "released_by_user_id" field.
func ReleasedByUserIDGTE(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDLT applies the LT predicate on the "released_by_user_id" field.
func ReleasedByUserIDLT(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDLTE applies the LTE predicate on the "released_by_user_id" field.
func ReleasedByUserIDLTE(v uuid.UUID) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReleasedByUserID), v))
	})
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// StartAtIsNil applies the IsNil predicate on the "start_at" field.
func StartAtIsNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartAt)))
	})
}

// StartAtNotNil applies the NotNil predicate on the "start_at" field.
func StartAtNotNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartAt)))
	})
}

// DurationDaysEQ applies the EQ predicate on the "duration_days" field.
func DurationDaysEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDurationDays), v))
	})
}

// DurationDaysNEQ applies the NEQ predicate on the "duration_days" field.
func DurationDaysNEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDurationDays), v))
	})
}

// DurationDaysIn applies the In predicate on the "duration_days" field.
func DurationDaysIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDurationDays), v...))
	})
}

// DurationDaysNotIn applies the NotIn predicate on the "duration_days" field.
func DurationDaysNotIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDurationDays), v...))
	})
}

// DurationDaysGT applies the GT predicate on the "duration_days" field.
func DurationDaysGT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDurationDays), v))
	})
}

// DurationDaysGTE applies the GTE predicate on the "duration_days" field.
func DurationDaysGTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDurationDays), v))
	})
}

// DurationDaysLT applies the LT predicate on the "duration_days" field.
func DurationDaysLT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDurationDays), v))
	})
}

// DurationDaysLTE applies the LTE predicate on the "duration_days" field.
func DurationDaysLTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDurationDays), v))
	})
}

// DurationDaysIsNil applies the IsNil predicate on the "duration_days" field.
func DurationDaysIsNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDurationDays)))
	})
}

// DurationDaysNotNil applies the NotNil predicate on the "duration_days" field.
func DurationDaysNotNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDurationDays)))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageIsNil applies the IsNil predicate on the "message" field.
func MessageIsNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMessage)))
	})
}

// MessageNotNil applies the NotNil predicate on the "message" field.
func MessageNotNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMessage)))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// AllocatedEQ applies the EQ predicate on the "allocated" field.
func AllocatedEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllocated), v))
	})
}

// AllocatedNEQ applies the NEQ predicate on the "allocated" field.
func AllocatedNEQ(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAllocated), v))
	})
}

// AllocatedIn applies the In predicate on the "allocated" field.
func AllocatedIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAllocated), v...))
	})
}

// AllocatedNotIn applies the NotIn predicate on the "allocated" field.
func AllocatedNotIn(vs ...uint32) predicate.CouponFixAmount {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAllocated), v...))
	})
}

// AllocatedGT applies the GT predicate on the "allocated" field.
func AllocatedGT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAllocated), v))
	})
}

// AllocatedGTE applies the GTE predicate on the "allocated" field.
func AllocatedGTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAllocated), v))
	})
}

// AllocatedLT applies the LT predicate on the "allocated" field.
func AllocatedLT(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAllocated), v))
	})
}

// AllocatedLTE applies the LTE predicate on the "allocated" field.
func AllocatedLTE(v uint32) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAllocated), v))
	})
}

// AllocatedIsNil applies the IsNil predicate on the "allocated" field.
func AllocatedIsNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAllocated)))
	})
}

// AllocatedNotNil applies the NotNil predicate on the "allocated" field.
func AllocatedNotNil() predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAllocated)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CouponFixAmount) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CouponFixAmount) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CouponFixAmount) predicate.CouponFixAmount {
	return predicate.CouponFixAmount(func(s *sql.Selector) {
		p(s.Not())
	})
}
