package logic

import (
	"context"

	"app/admin/rpc/internal/svc"

	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zero-vue-demo/app-center-public/rpc/admin"

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
func (l *MakeTokenLogic) MakeToken(in *admin.MakeToken_Request) (*admin.MakeToken_Response, error) {
	jwtxConfig := l.svcCtx.Config.JWTXConfig
	resp, err := l.svcCtx.JWTXRpc.MakeToken(l.ctx, &jwtx.MakeToken_Request{
		AccessGroup:        jwtxConfig.AccessGroup,
		AccessExpireByHour: jwtxConfig.AccessExpireByHour,
		AccessTerminal:     in.AccessTerminal,
		AccountID:          in.AccountID,
		RequestIP:          in.RequestIP,
		SingleEnd:          jwtxConfig.SingleEnd,
	})

	if err != nil {
		return nil, err
	}

	return &admin.MakeToken_Response{
		Token: resp.Token,
	}, nil
}
