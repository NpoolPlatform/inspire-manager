// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/coupondiscount"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponfixamount"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponspecialoffer"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/schema"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	couponallocatedMixin := schema.CouponAllocated{}.Mixin()
	couponallocated.Policy = privacy.NewPolicies(couponallocatedMixin[0], schema.CouponAllocated{})
	couponallocated.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := couponallocated.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	couponallocatedMixinFields0 := couponallocatedMixin[0].Fields()
	_ = couponallocatedMixinFields0
	couponallocatedFields := schema.CouponAllocated{}.Fields()
	_ = couponallocatedFields
	// couponallocatedDescCreatedAt is the schema descriptor for created_at field.
	couponallocatedDescCreatedAt := couponallocatedMixinFields0[0].Descriptor()
	// couponallocated.DefaultCreatedAt holds the default value on creation for the created_at field.
	couponallocated.DefaultCreatedAt = couponallocatedDescCreatedAt.Default.(func() uint32)
	// couponallocatedDescUpdatedAt is the schema descriptor for updated_at field.
	couponallocatedDescUpdatedAt := couponallocatedMixinFields0[1].Descriptor()
	// couponallocated.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	couponallocated.DefaultUpdatedAt = couponallocatedDescUpdatedAt.Default.(func() uint32)
	// couponallocated.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	couponallocated.UpdateDefaultUpdatedAt = couponallocatedDescUpdatedAt.UpdateDefault.(func() uint32)
	// couponallocatedDescDeletedAt is the schema descriptor for deleted_at field.
	couponallocatedDescDeletedAt := couponallocatedMixinFields0[2].Descriptor()
	// couponallocated.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	couponallocated.DefaultDeletedAt = couponallocatedDescDeletedAt.Default.(func() uint32)
	// couponallocatedDescType is the schema descriptor for type field.
	couponallocatedDescType := couponallocatedFields[3].Descriptor()
	// couponallocated.DefaultType holds the default value on creation for the type field.
	couponallocated.DefaultType = couponallocatedDescType.Default.(string)
	// couponallocatedDescID is the schema descriptor for id field.
	couponallocatedDescID := couponallocatedFields[0].Descriptor()
	// couponallocated.DefaultID holds the default value on creation for the id field.
	couponallocated.DefaultID = couponallocatedDescID.Default.(func() uuid.UUID)
	coupondiscountMixin := schema.CouponDiscount{}.Mixin()
	coupondiscount.Policy = privacy.NewPolicies(coupondiscountMixin[0], schema.CouponDiscount{})
	coupondiscount.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := coupondiscount.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	coupondiscountMixinFields0 := coupondiscountMixin[0].Fields()
	_ = coupondiscountMixinFields0
	coupondiscountFields := schema.CouponDiscount{}.Fields()
	_ = coupondiscountFields
	// coupondiscountDescCreatedAt is the schema descriptor for created_at field.
	coupondiscountDescCreatedAt := coupondiscountMixinFields0[0].Descriptor()
	// coupondiscount.DefaultCreatedAt holds the default value on creation for the created_at field.
	coupondiscount.DefaultCreatedAt = coupondiscountDescCreatedAt.Default.(func() uint32)
	// coupondiscountDescUpdatedAt is the schema descriptor for updated_at field.
	coupondiscountDescUpdatedAt := coupondiscountMixinFields0[1].Descriptor()
	// coupondiscount.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	coupondiscount.DefaultUpdatedAt = coupondiscountDescUpdatedAt.Default.(func() uint32)
	// coupondiscount.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	coupondiscount.UpdateDefaultUpdatedAt = coupondiscountDescUpdatedAt.UpdateDefault.(func() uint32)
	// coupondiscountDescDeletedAt is the schema descriptor for deleted_at field.
	coupondiscountDescDeletedAt := coupondiscountMixinFields0[2].Descriptor()
	// coupondiscount.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	coupondiscount.DefaultDeletedAt = coupondiscountDescDeletedAt.Default.(func() uint32)
	// coupondiscountDescDiscount is the schema descriptor for discount field.
	coupondiscountDescDiscount := coupondiscountFields[2].Descriptor()
	// coupondiscount.DefaultDiscount holds the default value on creation for the discount field.
	coupondiscount.DefaultDiscount = coupondiscountDescDiscount.Default.(decimal.Decimal)
	// coupondiscountDescStartAt is the schema descriptor for start_at field.
	coupondiscountDescStartAt := coupondiscountFields[4].Descriptor()
	// coupondiscount.DefaultStartAt holds the default value on creation for the start_at field.
	coupondiscount.DefaultStartAt = coupondiscountDescStartAt.Default.(uint32)
	// coupondiscountDescDurationDays is the schema descriptor for duration_days field.
	coupondiscountDescDurationDays := coupondiscountFields[5].Descriptor()
	// coupondiscount.DefaultDurationDays holds the default value on creation for the duration_days field.
	coupondiscount.DefaultDurationDays = coupondiscountDescDurationDays.Default.(uint32)
	// coupondiscountDescMessage is the schema descriptor for message field.
	coupondiscountDescMessage := coupondiscountFields[6].Descriptor()
	// coupondiscount.DefaultMessage holds the default value on creation for the message field.
	coupondiscount.DefaultMessage = coupondiscountDescMessage.Default.(string)
	// coupondiscountDescName is the schema descriptor for name field.
	coupondiscountDescName := coupondiscountFields[7].Descriptor()
	// coupondiscount.DefaultName holds the default value on creation for the name field.
	coupondiscount.DefaultName = coupondiscountDescName.Default.(string)
	// coupondiscountDescID is the schema descriptor for id field.
	coupondiscountDescID := coupondiscountFields[0].Descriptor()
	// coupondiscount.DefaultID holds the default value on creation for the id field.
	coupondiscount.DefaultID = coupondiscountDescID.Default.(func() uuid.UUID)
	couponfixamountMixin := schema.CouponFixAmount{}.Mixin()
	couponfixamount.Policy = privacy.NewPolicies(couponfixamountMixin[0], schema.CouponFixAmount{})
	couponfixamount.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := couponfixamount.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	couponfixamountMixinFields0 := couponfixamountMixin[0].Fields()
	_ = couponfixamountMixinFields0
	couponfixamountFields := schema.CouponFixAmount{}.Fields()
	_ = couponfixamountFields
	// couponfixamountDescCreatedAt is the schema descriptor for created_at field.
	couponfixamountDescCreatedAt := couponfixamountMixinFields0[0].Descriptor()
	// couponfixamount.DefaultCreatedAt holds the default value on creation for the created_at field.
	couponfixamount.DefaultCreatedAt = couponfixamountDescCreatedAt.Default.(func() uint32)
	// couponfixamountDescUpdatedAt is the schema descriptor for updated_at field.
	couponfixamountDescUpdatedAt := couponfixamountMixinFields0[1].Descriptor()
	// couponfixamount.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	couponfixamount.DefaultUpdatedAt = couponfixamountDescUpdatedAt.Default.(func() uint32)
	// couponfixamount.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	couponfixamount.UpdateDefaultUpdatedAt = couponfixamountDescUpdatedAt.UpdateDefault.(func() uint32)
	// couponfixamountDescDeletedAt is the schema descriptor for deleted_at field.
	couponfixamountDescDeletedAt := couponfixamountMixinFields0[2].Descriptor()
	// couponfixamount.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	couponfixamount.DefaultDeletedAt = couponfixamountDescDeletedAt.Default.(func() uint32)
	// couponfixamountDescDenomination is the schema descriptor for denomination field.
	couponfixamountDescDenomination := couponfixamountFields[2].Descriptor()
	// couponfixamount.DefaultDenomination holds the default value on creation for the denomination field.
	couponfixamount.DefaultDenomination = couponfixamountDescDenomination.Default.(decimal.Decimal)
	// couponfixamountDescCirculation is the schema descriptor for circulation field.
	couponfixamountDescCirculation := couponfixamountFields[3].Descriptor()
	// couponfixamount.DefaultCirculation holds the default value on creation for the circulation field.
	couponfixamount.DefaultCirculation = couponfixamountDescCirculation.Default.(decimal.Decimal)
	// couponfixamountDescStartAt is the schema descriptor for start_at field.
	couponfixamountDescStartAt := couponfixamountFields[5].Descriptor()
	// couponfixamount.DefaultStartAt holds the default value on creation for the start_at field.
	couponfixamount.DefaultStartAt = couponfixamountDescStartAt.Default.(uint32)
	// couponfixamountDescDurationDays is the schema descriptor for duration_days field.
	couponfixamountDescDurationDays := couponfixamountFields[6].Descriptor()
	// couponfixamount.DefaultDurationDays holds the default value on creation for the duration_days field.
	couponfixamount.DefaultDurationDays = couponfixamountDescDurationDays.Default.(uint32)
	// couponfixamountDescMessage is the schema descriptor for message field.
	couponfixamountDescMessage := couponfixamountFields[7].Descriptor()
	// couponfixamount.DefaultMessage holds the default value on creation for the message field.
	couponfixamount.DefaultMessage = couponfixamountDescMessage.Default.(string)
	// couponfixamountDescName is the schema descriptor for name field.
	couponfixamountDescName := couponfixamountFields[8].Descriptor()
	// couponfixamount.DefaultName holds the default value on creation for the name field.
	couponfixamount.DefaultName = couponfixamountDescName.Default.(string)
	// couponfixamountDescID is the schema descriptor for id field.
	couponfixamountDescID := couponfixamountFields[0].Descriptor()
	// couponfixamount.DefaultID holds the default value on creation for the id field.
	couponfixamount.DefaultID = couponfixamountDescID.Default.(func() uuid.UUID)
	couponspecialofferMixin := schema.CouponSpecialOffer{}.Mixin()
	couponspecialoffer.Policy = privacy.NewPolicies(couponspecialofferMixin[0], schema.CouponSpecialOffer{})
	couponspecialoffer.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := couponspecialoffer.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	couponspecialofferMixinFields0 := couponspecialofferMixin[0].Fields()
	_ = couponspecialofferMixinFields0
	couponspecialofferFields := schema.CouponSpecialOffer{}.Fields()
	_ = couponspecialofferFields
	// couponspecialofferDescCreatedAt is the schema descriptor for created_at field.
	couponspecialofferDescCreatedAt := couponspecialofferMixinFields0[0].Descriptor()
	// couponspecialoffer.DefaultCreatedAt holds the default value on creation for the created_at field.
	couponspecialoffer.DefaultCreatedAt = couponspecialofferDescCreatedAt.Default.(func() uint32)
	// couponspecialofferDescUpdatedAt is the schema descriptor for updated_at field.
	couponspecialofferDescUpdatedAt := couponspecialofferMixinFields0[1].Descriptor()
	// couponspecialoffer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	couponspecialoffer.DefaultUpdatedAt = couponspecialofferDescUpdatedAt.Default.(func() uint32)
	// couponspecialoffer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	couponspecialoffer.UpdateDefaultUpdatedAt = couponspecialofferDescUpdatedAt.UpdateDefault.(func() uint32)
	// couponspecialofferDescDeletedAt is the schema descriptor for deleted_at field.
	couponspecialofferDescDeletedAt := couponspecialofferMixinFields0[2].Descriptor()
	// couponspecialoffer.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	couponspecialoffer.DefaultDeletedAt = couponspecialofferDescDeletedAt.Default.(func() uint32)
	// couponspecialofferDescAmount is the schema descriptor for amount field.
	couponspecialofferDescAmount := couponspecialofferFields[3].Descriptor()
	// couponspecialoffer.DefaultAmount holds the default value on creation for the amount field.
	couponspecialoffer.DefaultAmount = couponspecialofferDescAmount.Default.(decimal.Decimal)
	// couponspecialofferDescStartAt is the schema descriptor for start_at field.
	couponspecialofferDescStartAt := couponspecialofferFields[5].Descriptor()
	// couponspecialoffer.DefaultStartAt holds the default value on creation for the start_at field.
	couponspecialoffer.DefaultStartAt = couponspecialofferDescStartAt.Default.(uint32)
	// couponspecialofferDescDurationDays is the schema descriptor for duration_days field.
	couponspecialofferDescDurationDays := couponspecialofferFields[6].Descriptor()
	// couponspecialoffer.DefaultDurationDays holds the default value on creation for the duration_days field.
	couponspecialoffer.DefaultDurationDays = couponspecialofferDescDurationDays.Default.(uint32)
	// couponspecialofferDescMessage is the schema descriptor for message field.
	couponspecialofferDescMessage := couponspecialofferFields[7].Descriptor()
	// couponspecialoffer.DefaultMessage holds the default value on creation for the message field.
	couponspecialoffer.DefaultMessage = couponspecialofferDescMessage.Default.(string)
	// couponspecialofferDescID is the schema descriptor for id field.
	couponspecialofferDescID := couponspecialofferFields[0].Descriptor()
	// couponspecialoffer.DefaultID holds the default value on creation for the id field.
	couponspecialoffer.DefaultID = couponspecialofferDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.3" // Version of ent codegen.
)