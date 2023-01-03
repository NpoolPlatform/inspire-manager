// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/invitationcode"
	"github.com/google/uuid"
)

// InvitationCode is the model entity for the InvitationCode schema.
type InvitationCode struct {
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
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// InvitationCode holds the value of the "invitation_code" field.
	InvitationCode string `json:"invitation_code,omitempty"`
	// Confirmed holds the value of the "confirmed" field.
	Confirmed bool `json:"confirmed,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InvitationCode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case invitationcode.FieldConfirmed, invitationcode.FieldDisabled:
			values[i] = new(sql.NullBool)
		case invitationcode.FieldCreatedAt, invitationcode.FieldUpdatedAt, invitationcode.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case invitationcode.FieldInvitationCode:
			values[i] = new(sql.NullString)
		case invitationcode.FieldID, invitationcode.FieldAppID, invitationcode.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type InvitationCode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InvitationCode fields.
func (ic *InvitationCode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case invitationcode.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ic.ID = *value
			}
		case invitationcode.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ic.CreatedAt = uint32(value.Int64)
			}
		case invitationcode.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ic.UpdatedAt = uint32(value.Int64)
			}
		case invitationcode.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ic.DeletedAt = uint32(value.Int64)
			}
		case invitationcode.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ic.AppID = *value
			}
		case invitationcode.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ic.UserID = *value
			}
		case invitationcode.FieldInvitationCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field invitation_code", values[i])
			} else if value.Valid {
				ic.InvitationCode = value.String
			}
		case invitationcode.FieldConfirmed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field confirmed", values[i])
			} else if value.Valid {
				ic.Confirmed = value.Bool
			}
		case invitationcode.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				ic.Disabled = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this InvitationCode.
// Note that you need to call InvitationCode.Unwrap() before calling this method if this InvitationCode
// was returned from a transaction, and the transaction was committed or rolled back.
func (ic *InvitationCode) Update() *InvitationCodeUpdateOne {
	return (&InvitationCodeClient{config: ic.config}).UpdateOne(ic)
}

// Unwrap unwraps the InvitationCode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ic *InvitationCode) Unwrap() *InvitationCode {
	_tx, ok := ic.config.driver.(*txDriver)
	if !ok {
		panic("ent: InvitationCode is not a transactional entity")
	}
	ic.config.driver = _tx.drv
	return ic
}

// String implements the fmt.Stringer.
func (ic *InvitationCode) String() string {
	var builder strings.Builder
	builder.WriteString("InvitationCode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ic.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ic.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ic.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ic.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ic.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ic.UserID))
	builder.WriteString(", ")
	builder.WriteString("invitation_code=")
	builder.WriteString(ic.InvitationCode)
	builder.WriteString(", ")
	builder.WriteString("confirmed=")
	builder.WriteString(fmt.Sprintf("%v", ic.Confirmed))
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", ic.Disabled))
	builder.WriteByte(')')
	return builder.String()
}

// InvitationCodes is a parsable slice of InvitationCode.
type InvitationCodes []*InvitationCode

func (ic InvitationCodes) config(cfg config) {
	for _i := range ic {
		ic[_i].config = cfg
	}
}
