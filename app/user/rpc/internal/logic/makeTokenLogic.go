package logic

import (
	"context"

	"app/com.docker.devenvironments.code/public/user"
	"app/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeTokenLogic {
	return &MakeTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成 token（登录）
func (l *MakeTokenLogic) MakeToken(in *user.MakeToken_Request) (*user.MakeToken_Response, error) {
	// todo: add your logic here and delete this line

	return &user.MakeToken_Response{}, nil
}
