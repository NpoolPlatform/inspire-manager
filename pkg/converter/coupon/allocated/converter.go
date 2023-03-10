package allocated

import (
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/allocated"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.CouponAllocated) *npool.Allocated {
	if row == nil {
		return nil
	}
	ret := &npool.Allocated{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		CouponType: npool.CouponType(npool.CouponType_value[row.CouponType]),
		CouponID:   row.CouponID.String(),
		Value:      row.Value.String(),
		Used:       row.Used,
		UsedAt:     row.UsedAt,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}

	if row.Used {
		orderID := row.UsedByOrderID.String()
		ret.UsedByOrderID = &orderID
	}

	return ret
}

func Ent2GrpcMany(rows []*ent.CouponAllocated) []*npool.Allocated {
	infos := []*npool.Allocated{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
