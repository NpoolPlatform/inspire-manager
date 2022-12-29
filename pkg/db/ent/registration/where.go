// Code generated by ent, DO NOT EDIT.

package registration

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// InviterID applies equality check predicate on the "inviter_id" field. It's identical to InviterIDEQ.
func InviterID(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviterID), v))
	})
}

// InviteeID applies equality check predicate on the "invitee_id" field. It's identical to InviteeIDEQ.
func InviteeID(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviteeID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// InviterIDEQ applies the EQ predicate on the "inviter_id" field.
func InviterIDEQ(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviterID), v))
	})
}

// InviterIDNEQ applies the NEQ predicate on the "inviter_id" field.
func InviterIDNEQ(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInviterID), v))
	})
}

// InviterIDIn applies the In predicate on the "inviter_id" field.
func InviterIDIn(vs ...uuid.UUID) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInviterID), v...))
	})
}

// InviterIDNotIn applies the NotIn predicate on the "inviter_id" field.
func InviterIDNotIn(vs ...uuid.UUID) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInviterID), v...))
	})
}

// InviterIDGT applies the GT predicate on the "inviter_id" field.
func InviterIDGT(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInviterID), v))
	})
}

// InviterIDGTE applies the GTE predicate on the "inviter_id" field.
func InviterIDGTE(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInviterID), v))
	})
}

// InviterIDLT applies the LT predicate on the "inviter_id" field.
func InviterIDLT(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInviterID), v))
	})
}

// InviterIDLTE applies the LTE predicate on the "inviter_id" field.
func InviterIDLTE(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInviterID), v))
	})
}

// InviteeIDEQ applies the EQ predicate on the "invitee_id" field.
func InviteeIDEQ(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInviteeID), v))
	})
}

// InviteeIDNEQ applies the NEQ predicate on the "invitee_id" field.
func InviteeIDNEQ(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInviteeID), v))
	})
}

// InviteeIDIn applies the In predicate on the "invitee_id" field.
func InviteeIDIn(vs ...uuid.UUID) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInviteeID), v...))
	})
}

// InviteeIDNotIn applies the NotIn predicate on the "invitee_id" field.
func InviteeIDNotIn(vs ...uuid.UUID) predicate.Registration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInviteeID), v...))
	})
}

// InviteeIDGT applies the GT predicate on the "invitee_id" field.
func InviteeIDGT(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInviteeID), v))
	})
}

// InviteeIDGTE applies the GTE predicate on the "invitee_id" field.
func InviteeIDGTE(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInviteeID), v))
	})
}

// InviteeIDLT applies the LT predicate on the "invitee_id" field.
func InviteeIDLT(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInviteeID), v))
	})
}

// InviteeIDLTE applies the LTE predicate on the "invitee_id" field.
func InviteeIDLTE(v uuid.UUID) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInviteeID), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Registration) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Registration) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
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
func Not(p predicate.Registration) predicate.Registration {
	return predicate.Registration(func(s *sql.Selector) {
		p(s.Not())
	})
}
