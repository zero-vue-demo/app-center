package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInMobileLogic {
	return &SignInMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInMobileLogic) SignInMobile(req *types.Client_Public_SignInMobile_Request) (resp *types.Client_Public_SignInMobile_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
