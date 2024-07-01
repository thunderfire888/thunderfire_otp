package logic

import (
	"context"
	"github.com/thunderfire888/thunderfire_otp/helper/otpx"
	"github.com/thunderfire888/thunderfire_otp/rpc/otpclient"

	"github.com/thunderfire888/thunderfire_otp/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateLogic {
	return &ValidateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateLogic) Validate(in *otpclient.OtpVaildRequest) (*otpclient.OtpVaildResponse, error) {

	return &otpclient.OtpVaildResponse{
		Vaild: otpx.Validate(in.PassCode, in.Secret),
	}, nil
}
