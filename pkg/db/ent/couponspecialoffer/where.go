// Code generated by ent, DO NOT EDIT.

package couponspecialoffer

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v decimal.Decimal) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// ReleasedByUserID applies equality check predicate on the "released_by_user_id" field. It's identical to ReleasedByUserIDEQ.
func ReleasedByUserID(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleasedByUserID), v))
	})
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// DurationDays applies equality check predicate on the "duration_days" field. It's identical to DurationDaysEQ.
func DurationDays(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDurationDays), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v decimal.Decimal) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v decimal.Decimal) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...decimal.Decimal) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...decimal.Decimal) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v decimal.Decimal) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v decimal.Decimal) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v decimal.Decimal) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v decimal.Decimal) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAmount)))
	})
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAmount)))
	})
}

// ReleasedByUserIDEQ applies the EQ predicate on the "released_by_user_id" field.
func ReleasedByUserIDEQ(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDNEQ applies the NEQ predicate on the "released_by_user_id" field.
func ReleasedByUserIDNEQ(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDIn applies the In predicate on the "released_by_user_id" field.
func ReleasedByUserIDIn(vs ...uuid.UUID) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReleasedByUserID), v...))
	})
}

// ReleasedByUserIDNotIn applies the NotIn predicate on the "released_by_user_id" field.
func ReleasedByUserIDNotIn(vs ...uuid.UUID) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReleasedByUserID), v...))
	})
}

// ReleasedByUserIDGT applies the GT predicate on the "released_by_user_id" field.
func ReleasedByUserIDGT(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDGTE applies the GTE predicate on the "released_by_user_id" field.
func ReleasedByUserIDGTE(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDLT applies the LT predicate on the "released_by_user_id" field.
func ReleasedByUserIDLT(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReleasedByUserID), v))
	})
}

// ReleasedByUserIDLTE applies the LTE predicate on the "released_by_user_id" field.
func ReleasedByUserIDLTE(v uuid.UUID) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReleasedByUserID), v))
	})
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// StartAtIsNil applies the IsNil predicate on the "start_at" field.
func StartAtIsNil() predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartAt)))
	})
}

// StartAtNotNil applies the NotNil predicate on the "start_at" field.
func StartAtNotNil() predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartAt)))
	})
}

// DurationDaysEQ applies the EQ predicate on the "duration_days" field.
func DurationDaysEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDurationDays), v))
	})
}

// DurationDaysNEQ applies the NEQ predicate on the "duration_days" field.
func DurationDaysNEQ(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDurationDays), v))
	})
}

// DurationDaysIn applies the In predicate on the "duration_days" field.
func DurationDaysIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDurationDays), v...))
	})
}

// DurationDaysNotIn applies the NotIn predicate on the "duration_days" field.
func DurationDaysNotIn(vs ...uint32) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDurationDays), v...))
	})
}

// DurationDaysGT applies the GT predicate on the "duration_days" field.
func DurationDaysGT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDurationDays), v))
	})
}

// DurationDaysGTE applies the GTE predicate on the "duration_days" field.
func DurationDaysGTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDurationDays), v))
	})
}

// DurationDaysLT applies the LT predicate on the "duration_days" field.
func DurationDaysLT(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDurationDays), v))
	})
}

// DurationDaysLTE applies the LTE predicate on the "duration_days" field.
func DurationDaysLTE(v uint32) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDurationDays), v))
	})
}

// DurationDaysIsNil applies the IsNil predicate on the "duration_days" field.
func DurationDaysIsNil() predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDurationDays)))
	})
}

// DurationDaysNotNil applies the NotNil predicate on the "duration_days" field.
func DurationDaysNotNil() predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDurationDays)))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.CouponSpecialOffer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageIsNil applies the IsNil predicate on the "message" field.
func MessageIsNil() predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMessage)))
	})
}

// MessageNotNil applies the NotNil predicate on the "message" field.
func MessageNotNil() predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMessage)))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CouponSpecialOffer) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CouponSpecialOffer) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
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
func Not(p predicate.CouponSpecialOffer) predicate.CouponSpecialOffer {
	return predicate.CouponSpecialOffer(func(s *sql.Selector) {
		p(s.Not())
	})
}
