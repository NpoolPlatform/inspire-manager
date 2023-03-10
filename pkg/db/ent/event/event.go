// Code generated by ent, DO NOT EDIT.

package entevent

import (
	"entgo.io/ent"
	"github.com/NpoolPlatform/message/npool/inspire/mgr/v1/event"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
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
	// FieldEventType holds the string denoting the event_type field in the database.
	FieldEventType = "event_type"
	// FieldCoupons holds the string denoting the coupons field in the database.
	FieldCoupons = "coupons"
	// FieldCredits holds the string denoting the credits field in the database.
	FieldCredits = "credits"
	// FieldCreditsPerUsd holds the string denoting the credits_per_usd field in the database.
	FieldCreditsPerUsd = "credits_per_usd"
	// FieldMaxConsecutive holds the string denoting the max_consecutive field in the database.
	FieldMaxConsecutive = "max_consecutive"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldInviterLayers holds the string denoting the inviter_layers field in the database.
	FieldInviterLayers = "inviter_layers"
	// Table holds the table name of the event in the database.
	Table = "events"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldEventType,
	FieldCoupons,
	FieldCredits,
	FieldCreditsPerUsd,
	FieldMaxConsecutive,
	FieldGoodID,
	FieldInviterLayers,
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
	// DefaultEventType holds the default value on creation for the "event_type" field.
	DefaultEventType string
	// DefaultCoupons holds the default value on creation for the "coupons" field.
	DefaultCoupons []event.Coupon
	// DefaultCredits holds the default value on creation for the "credits" field.
	DefaultCredits decimal.Decimal
	// DefaultCreditsPerUsd holds the default value on creation for the "credits_per_usd" field.
	DefaultCreditsPerUsd decimal.Decimal
	// DefaultMaxConsecutive holds the default value on creation for the "max_consecutive" field.
	DefaultMaxConsecutive uint32
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultInviterLayers holds the default value on creation for the "inviter_layers" field.
	DefaultInviterLayers uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
