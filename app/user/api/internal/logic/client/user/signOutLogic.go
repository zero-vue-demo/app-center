package user

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignOutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignOutLogic {
	return &SignOutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignOutLogic) SignOut(req *types.Client_User_SignOut_Request) (resp *types.Client_User_SignOut_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
