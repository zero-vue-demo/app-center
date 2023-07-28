package logic

import (
	"context"

	"app/com.docker.devenvironments.code/public/user"
	"app/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTokenLogic {
	return &CheckTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验 token（拓展校验、刷新 token）
func (l *CheckTokenLogic) CheckToken(in *user.CheckToken_Request) (*user.CheckToken_Response, error) {
	// todo: add your logic here and delete this line

	return &user.CheckToken_Response{}, nil
}
