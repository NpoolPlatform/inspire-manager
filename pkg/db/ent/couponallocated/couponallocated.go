// Code generated by ent, DO NOT EDIT.

package couponallocated

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the couponallocated type in the database.
	Label = "coupon_allocated"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCouponType holds the string denoting the coupon_type field in the database.
	FieldCouponType = "coupon_type"
	// FieldCouponID holds the string denoting the coupon_id field in the database.
	FieldCouponID = "coupon_id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldUsed holds the string denoting the used field in the database.
	FieldUsed = "used"
	// FieldUsedAt holds the string denoting the used_at field in the database.
	FieldUsedAt = "used_at"
	// FieldUsedByOrderID holds the string denoting the used_by_order_id field in the database.
	FieldUsedByOrderID = "used_by_order_id"
	// Table holds the table name of the couponallocated in the database.
	Table = "coupon_allocateds"
)

// Columns holds all SQL columns for couponallocated fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldUserID,
	FieldCouponType,
	FieldCouponID,
	FieldValue,
	FieldUsed,
	FieldUsedAt,
	FieldUsedByOrderID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/inspire-manager/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultCouponType holds the default value on creation for the "coupon_type" field.
	DefaultCouponType string
	// DefaultValue holds the default value on creation for the "value" field.
	DefaultValue decimal.Decimal
	// DefaultUsed holds the default value on creation for the "used" field.
	DefaultUsed bool
	// DefaultUsedAt holds the default value on creation for the "used_at" field.
	DefaultUsedAt uint32
	// DefaultUsedByOrderID holds the default value on creation for the "used_by_order_id" field.
	DefaultUsedByOrderID func() uuid.UUID
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
