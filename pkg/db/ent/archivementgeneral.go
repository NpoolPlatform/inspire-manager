// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/archivementgeneral"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ArchivementGeneral is the model entity for the ArchivementGeneral schema.
type ArchivementGeneral struct {
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
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// TotalUnits holds the value of the "total_units" field.
	TotalUnits uint32 `json:"total_units,omitempty"`
	// TotalUnitsV1 holds the value of the "total_units_v1" field.
	TotalUnitsV1 decimal.Decimal `json:"total_units_v1,omitempty"`
	// SelfUnits holds the value of the "self_units" field.
	SelfUnits uint32 `json:"self_units,omitempty"`
	// SelfUnitsV1 holds the value of the "self_units_v1" field.
	SelfUnitsV1 decimal.Decimal `json:"self_units_v1,omitempty"`
	// TotalAmount holds the value of the "total_amount" field.
	TotalAmount decimal.Decimal `json:"total_amount,omitempty"`
	// SelfAmount holds the value of the "self_amount" field.
	SelfAmount decimal.Decimal `json:"self_amount,omitempty"`
	// TotalCommission holds the value of the "total_commission" field.
	TotalCommission decimal.Decimal `json:"total_commission,omitempty"`
	// SelfCommission holds the value of the "self_commission" field.
	SelfCommission decimal.Decimal `json:"self_commission,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ArchivementGeneral) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case archivementgeneral.FieldTotalUnitsV1, archivementgeneral.FieldSelfUnitsV1, archivementgeneral.FieldTotalAmount, archivementgeneral.FieldSelfAmount, archivementgeneral.FieldTotalCommission, archivementgeneral.FieldSelfCommission:
			values[i] = new(decimal.Decimal)
		case archivementgeneral.FieldCreatedAt, archivementgeneral.FieldUpdatedAt, archivementgeneral.FieldDeletedAt, archivementgeneral.FieldTotalUnits, archivementgeneral.FieldSelfUnits:
			values[i] = new(sql.NullInt64)
		case archivementgeneral.FieldID, archivementgeneral.FieldAppID, archivementgeneral.FieldUserID, archivementgeneral.FieldGoodID, archivementgeneral.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ArchivementGeneral", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ArchivementGeneral fields.
func (ag *ArchivementGeneral) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case archivementgeneral.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ag.ID = *value
			}
		case archivementgeneral.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ag.CreatedAt = uint32(value.Int64)
			}
		case archivementgeneral.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ag.UpdatedAt = uint32(value.Int64)
			}
		case archivementgeneral.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ag.DeletedAt = uint32(value.Int64)
			}
		case archivementgeneral.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ag.AppID = *value
			}
		case archivementgeneral.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ag.UserID = *value
			}
		case archivementgeneral.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				ag.GoodID = *value
			}
		case archivementgeneral.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				ag.CoinTypeID = *value
			}
		case archivementgeneral.FieldTotalUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_units", values[i])
			} else if value.Valid {
				ag.TotalUnits = uint32(value.Int64)
			}
		case archivementgeneral.FieldTotalUnitsV1:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total_units_v1", values[i])
			} else if value != nil {
				ag.TotalUnitsV1 = *value
			}
		case archivementgeneral.FieldSelfUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field self_units", values[i])
			} else if value.Valid {
				ag.SelfUnits = uint32(value.Int64)
			}
		case archivementgeneral.FieldSelfUnitsV1:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field self_units_v1", values[i])
			} else if value != nil {
				ag.SelfUnitsV1 = *value
			}
		case archivementgeneral.FieldTotalAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total_amount", values[i])
			} else if value != nil {
				ag.TotalAmount = *value
			}
		case archivementgeneral.FieldSelfAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field self_amount", values[i])
			} else if value != nil {
				ag.SelfAmount = *value
			}
		case archivementgeneral.FieldTotalCommission:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total_commission", values[i])
			} else if value != nil {
				ag.TotalCommission = *value
			}
		case archivementgeneral.FieldSelfCommission:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field self_commission", values[i])
			} else if value != nil {
				ag.SelfCommission = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ArchivementGeneral.
// Note that you need to call ArchivementGeneral.Unwrap() before calling this method if this ArchivementGeneral
// was returned from a transaction, and the transaction was committed or rolled back.
func (ag *ArchivementGeneral) Update() *ArchivementGeneralUpdateOne {
	return (&ArchivementGeneralClient{config: ag.config}).UpdateOne(ag)
}

// Unwrap unwraps the ArchivementGeneral entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ag *ArchivementGeneral) Unwrap() *ArchivementGeneral {
	_tx, ok := ag.config.driver.(*txDriver)
	if !ok {
		panic("ent: ArchivementGeneral is not a transactional entity")
	}
	ag.config.driver = _tx.drv
	return ag
}

// String implements the fmt.Stringer.
func (ag *ArchivementGeneral) String() string {
	var builder strings.Builder
	builder.WriteString("ArchivementGeneral(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ag.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ag.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ag.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ag.UserID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", ag.GoodID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", ag.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("total_units=")
	builder.WriteString(fmt.Sprintf("%v", ag.TotalUnits))
	builder.WriteString(", ")
	builder.WriteString("total_units_v1=")
	builder.WriteString(fmt.Sprintf("%v", ag.TotalUnitsV1))
	builder.WriteString(", ")
	builder.WriteString("self_units=")
	builder.WriteString(fmt.Sprintf("%v", ag.SelfUnits))
	builder.WriteString(", ")
	builder.WriteString("self_units_v1=")
	builder.WriteString(fmt.Sprintf("%v", ag.SelfUnitsV1))
	builder.WriteString(", ")
	builder.WriteString("total_amount=")
	builder.WriteString(fmt.Sprintf("%v", ag.TotalAmount))
	builder.WriteString(", ")
	builder.WriteString("self_amount=")
	builder.WriteString(fmt.Sprintf("%v", ag.SelfAmount))
	builder.WriteString(", ")
	builder.WriteString("total_commission=")
	builder.WriteString(fmt.Sprintf("%v", ag.TotalCommission))
	builder.WriteString(", ")
	builder.WriteString("self_commission=")
	builder.WriteString(fmt.Sprintf("%v", ag.SelfCommission))
	builder.WriteByte(')')
	return builder.String()
}

// ArchivementGenerals is a parsable slice of ArchivementGeneral.
type ArchivementGenerals []*ArchivementGeneral

func (ag ArchivementGenerals) config(cfg config) {
	for _i := range ag {
		ag[_i].config = cfg
	}
}
