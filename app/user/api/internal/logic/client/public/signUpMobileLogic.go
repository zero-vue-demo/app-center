package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpMobileLogic {
	return &SignUpMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpMobileLogic) SignUpMobile(req *types.Client_Public_SignUpMobile_Request) (resp *types.Client_Public_SignUpMobile_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
