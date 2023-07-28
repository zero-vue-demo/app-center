package logic

import (
	"context"

	"app/user/rpc/internal/svc"

	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zero-vue-demo/app-center-public/rpc/user"

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

	jwtxConfig := l.svcCtx.Config.JWTXConfig
	resp, err := l.svcCtx.JWTXRpc.CheckToken(l.ctx, &jwtx.CheckToken_Request{
		AccessGroup:        jwtxConfig.AccessGroup,
		RefreshInterval:    jwtxConfig.RefreshInterval,
		FaultTolerance:     jwtxConfig.FaultTolerance,
		AutomaticRenewal:   jwtxConfig.AutomaticRenewal,
		AccessExpireByHour: jwtxConfig.AccessExpireByHour,
		CheckIP:            jwtxConfig.CheckIP,
		RequestIP:          in.RequestIP,
		RequestToken:       in.RequestToken,
	})

	if err != nil {
		return nil, err
	}

	return &user.CheckToken_Response{
		TokenID:             resp.TokenID,
		AccountID:           resp.AccountID,
		AccessTerminal:      resp.AccessTerminal,
		MakeTokenIP:         resp.MakeTokenIP,
		ExpirationTimestamp: resp.ExpirationTimestamp,
		NewToken:            resp.NewToken,
	}, nil
}
