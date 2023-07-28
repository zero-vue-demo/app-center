package logic

import (
	"context"

	"app/com.docker.devenvironments.code/public/user"
	"app/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 移除 token（安全退出）
func (l *DeleteTokenLogic) DeleteToken(in *user.DeleteToken_Request) (*user.DeleteToken_Response, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteToken_Response{}, nil
}
