// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/registration"
	"github.com/google/uuid"
)

// Registration is the model entity for the Registration schema.
type Registration struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// InviterID holds the value of the "inviter_id" field.
	InviterID uuid.UUID `json:"inviter_id,omitempty"`
	// InviteeID holds the value of the "invitee_id" field.
	InviteeID uuid.UUID `json:"invitee_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Registration) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case registration.FieldCreatedAt, registration.FieldUpdatedAt, registration.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case registration.FieldID, registration.FieldAppID, registration.FieldInviterID, registration.FieldInviteeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Registration", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Registration fields.
func (r *Registration) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case registration.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case registration.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = uint32(value.Int64)
			}
		case registration.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = uint32(value.Int64)
			}
		case registration.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = uint32(value.Int64)
			}
		case registration.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				r.AppID = *value
			}
		case registration.FieldInviterID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field inviter_id", values[i])
			} else if value != nil {
				r.InviterID = *value
			}
		case registration.FieldInviteeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field invitee_id", values[i])
			} else if value != nil {
				r.InviteeID = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Registration.
// Note that you need to call Registration.Unwrap() before calling this method if this Registration
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Registration) Update() *RegistrationUpdateOne {
	return (&RegistrationClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Registration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Registration) Unwrap() *Registration {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Registration is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Registration) String() string {
	var builder strings.Builder
	builder.WriteString("Registration(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", r.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", r.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", r.AppID))
	builder.WriteString(", ")
	builder.WriteString("inviter_id=")
	builder.WriteString(fmt.Sprintf("%v", r.InviterID))
	builder.WriteString(", ")
	builder.WriteString("invitee_id=")
	builder.WriteString(fmt.Sprintf("%v", r.InviteeID))
	builder.WriteByte(')')
	return builder.String()
}

// Registrations is a parsable slice of Registration.
type Registrations []*Registration

func (r Registrations) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}