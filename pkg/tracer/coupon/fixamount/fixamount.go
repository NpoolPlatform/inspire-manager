package allocated

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/coupon/fixamount"
)

func trace(span trace1.Span, in *npool.FixAmountReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("Denomination.%v", index), in.GetDenomination()),
		attribute.String(fmt.Sprintf("Circulation.%v", index), in.GetCirculation()),
		attribute.String(fmt.Sprintf("ReleasedByUserID.%v", index), in.GetReleasedByUserID()),
		attribute.Int(fmt.Sprintf("StartAt.%v", index), int(in.GetStartAt())),
		attribute.Int(fmt.Sprintf("DurationDays.%v", index), int(in.GetDurationDays())),
		attribute.String(fmt.Sprintf("Message.%v", index), in.GetMessage()),
		attribute.String(fmt.Sprintf("Name.%v", index), in.GetName()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.FixAmountReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("ReleasedByUserID.Op", in.GetReleasedByUserID().GetOp()),
		attribute.String("ReleasedByUserID.Value", in.GetReleasedByUserID().GetValue()),
		attribute.String("StartAt.Op", in.GetStartAt().GetOp()),
		attribute.Int("StartAt.Value", int(in.GetStartAt().GetValue())),
		attribute.String("DurationDays.Op", in.GetDurationDays().GetOp()),
		attribute.Int("DurationDays.Value", int(in.GetDurationDays().GetValue())),
		attribute.String("Name.Op", in.GetName().GetOp()),
		attribute.String("Name.Value", in.GetName().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.FixAmountReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
