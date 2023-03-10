package registration

import (
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/invitation/registration"
)

func Ent2Grpc(row *ent.Registration) *npool.Registration {
	if row == nil {
		return nil
	}

	return &npool.Registration{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		InviterID: row.InviterID.String(),
		InviteeID: row.InviteeID.String(),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Registration) []*npool.Registration {
	infos := []*npool.Registration{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
