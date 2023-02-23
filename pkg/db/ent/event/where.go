// Code generated by ent, DO NOT EDIT.

package event

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// EventType applies equality check predicate on the "event_type" field. It's identical to EventTypeEQ.
func EventType(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventType), v))
	})
}

// Credits applies equality check predicate on the "credits" field. It's identical to CreditsEQ.
func Credits(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCredits), v))
	})
}

// CreditsPerUsd applies equality check predicate on the "credits_per_usd" field. It's identical to CreditsPerUsdEQ.
func CreditsPerUsd(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreditsPerUsd), v))
	})
}

// MaxConsecutive applies equality check predicate on the "max_consecutive" field. It's identical to MaxConsecutiveEQ.
func MaxConsecutive(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxConsecutive), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// EventTypeEQ applies the EQ predicate on the "event_type" field.
func EventTypeEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventType), v))
	})
}

// EventTypeNEQ applies the NEQ predicate on the "event_type" field.
func EventTypeNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventType), v))
	})
}

// EventTypeIn applies the In predicate on the "event_type" field.
func EventTypeIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEventType), v...))
	})
}

// EventTypeNotIn applies the NotIn predicate on the "event_type" field.
func EventTypeNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEventType), v...))
	})
}

// EventTypeGT applies the GT predicate on the "event_type" field.
func EventTypeGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventType), v))
	})
}

// EventTypeGTE applies the GTE predicate on the "event_type" field.
func EventTypeGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventType), v))
	})
}

// EventTypeLT applies the LT predicate on the "event_type" field.
func EventTypeLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventType), v))
	})
}

// EventTypeLTE applies the LTE predicate on the "event_type" field.
func EventTypeLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventType), v))
	})
}

// EventTypeContains applies the Contains predicate on the "event_type" field.
func EventTypeContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventType), v))
	})
}

// EventTypeHasPrefix applies the HasPrefix predicate on the "event_type" field.
func EventTypeHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventType), v))
	})
}

// EventTypeHasSuffix applies the HasSuffix predicate on the "event_type" field.
func EventTypeHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventType), v))
	})
}

// EventTypeIsNil applies the IsNil predicate on the "event_type" field.
func EventTypeIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEventType)))
	})
}

// EventTypeNotNil applies the NotNil predicate on the "event_type" field.
func EventTypeNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEventType)))
	})
}

// EventTypeEqualFold applies the EqualFold predicate on the "event_type" field.
func EventTypeEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventType), v))
	})
}

// EventTypeContainsFold applies the ContainsFold predicate on the "event_type" field.
func EventTypeContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventType), v))
	})
}

// CouponIdsIsNil applies the IsNil predicate on the "coupon_ids" field.
func CouponIdsIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCouponIds)))
	})
}

// CouponIdsNotNil applies the NotNil predicate on the "coupon_ids" field.
func CouponIdsNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCouponIds)))
	})
}

// CreditsEQ applies the EQ predicate on the "credits" field.
func CreditsEQ(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCredits), v))
	})
}

// CreditsNEQ applies the NEQ predicate on the "credits" field.
func CreditsNEQ(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCredits), v))
	})
}

// CreditsIn applies the In predicate on the "credits" field.
func CreditsIn(vs ...decimal.Decimal) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCredits), v...))
	})
}

// CreditsNotIn applies the NotIn predicate on the "credits" field.
func CreditsNotIn(vs ...decimal.Decimal) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCredits), v...))
	})
}

// CreditsGT applies the GT predicate on the "credits" field.
func CreditsGT(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCredits), v))
	})
}

// CreditsGTE applies the GTE predicate on the "credits" field.
func CreditsGTE(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCredits), v))
	})
}

// CreditsLT applies the LT predicate on the "credits" field.
func CreditsLT(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCredits), v))
	})
}

// CreditsLTE applies the LTE predicate on the "credits" field.
func CreditsLTE(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCredits), v))
	})
}

// CreditsIsNil applies the IsNil predicate on the "credits" field.
func CreditsIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCredits)))
	})
}

// CreditsNotNil applies the NotNil predicate on the "credits" field.
func CreditsNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCredits)))
	})
}

// CreditsPerUsdEQ applies the EQ predicate on the "credits_per_usd" field.
func CreditsPerUsdEQ(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreditsPerUsd), v))
	})
}

// CreditsPerUsdNEQ applies the NEQ predicate on the "credits_per_usd" field.
func CreditsPerUsdNEQ(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreditsPerUsd), v))
	})
}

// CreditsPerUsdIn applies the In predicate on the "credits_per_usd" field.
func CreditsPerUsdIn(vs ...decimal.Decimal) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreditsPerUsd), v...))
	})
}

// CreditsPerUsdNotIn applies the NotIn predicate on the "credits_per_usd" field.
func CreditsPerUsdNotIn(vs ...decimal.Decimal) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreditsPerUsd), v...))
	})
}

// CreditsPerUsdGT applies the GT predicate on the "credits_per_usd" field.
func CreditsPerUsdGT(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreditsPerUsd), v))
	})
}

// CreditsPerUsdGTE applies the GTE predicate on the "credits_per_usd" field.
func CreditsPerUsdGTE(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreditsPerUsd), v))
	})
}

// CreditsPerUsdLT applies the LT predicate on the "credits_per_usd" field.
func CreditsPerUsdLT(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreditsPerUsd), v))
	})
}

// CreditsPerUsdLTE applies the LTE predicate on the "credits_per_usd" field.
func CreditsPerUsdLTE(v decimal.Decimal) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreditsPerUsd), v))
	})
}

// CreditsPerUsdIsNil applies the IsNil predicate on the "credits_per_usd" field.
func CreditsPerUsdIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreditsPerUsd)))
	})
}

// CreditsPerUsdNotNil applies the NotNil predicate on the "credits_per_usd" field.
func CreditsPerUsdNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreditsPerUsd)))
	})
}

// MaxConsecutiveEQ applies the EQ predicate on the "max_consecutive" field.
func MaxConsecutiveEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxConsecutive), v))
	})
}

// MaxConsecutiveNEQ applies the NEQ predicate on the "max_consecutive" field.
func MaxConsecutiveNEQ(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxConsecutive), v))
	})
}

// MaxConsecutiveIn applies the In predicate on the "max_consecutive" field.
func MaxConsecutiveIn(vs ...uint32) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMaxConsecutive), v...))
	})
}

// MaxConsecutiveNotIn applies the NotIn predicate on the "max_consecutive" field.
func MaxConsecutiveNotIn(vs ...uint32) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMaxConsecutive), v...))
	})
}

// MaxConsecutiveGT applies the GT predicate on the "max_consecutive" field.
func MaxConsecutiveGT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxConsecutive), v))
	})
}

// MaxConsecutiveGTE applies the GTE predicate on the "max_consecutive" field.
func MaxConsecutiveGTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxConsecutive), v))
	})
}

// MaxConsecutiveLT applies the LT predicate on the "max_consecutive" field.
func MaxConsecutiveLT(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxConsecutive), v))
	})
}

// MaxConsecutiveLTE applies the LTE predicate on the "max_consecutive" field.
func MaxConsecutiveLTE(v uint32) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxConsecutive), v))
	})
}

// MaxConsecutiveIsNil applies the IsNil predicate on the "max_consecutive" field.
func MaxConsecutiveIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMaxConsecutive)))
	})
}

// MaxConsecutiveNotNil applies the NotNil predicate on the "max_consecutive" field.
func MaxConsecutiveNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMaxConsecutive)))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// GoodIDIsNil applies the IsNil predicate on the "good_id" field.
func GoodIDIsNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodID)))
	})
}

// GoodIDNotNil applies the NotNil predicate on the "good_id" field.
func GoodIDNotNil() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodID)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
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
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		p(s.Not())
	})
}
