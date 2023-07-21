package user

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPasswordLogic {
	return &EditPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditPasswordLogic) EditPassword(req *types.Client_User_EditPassword_Request) (resp *types.Client_User_EditPassword_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
