package logic

import (
	"context"

	"app/admin/rpc/internal/svc"

	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zero-vue-demo/app-center-public/rpc/admin"

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
func (l *DeleteTokenLogic) DeleteToken(in *admin.DeleteToken_Request) (*admin.DeleteToken_Response, error) {

	_, err := l.svcCtx.JWTXRpc.DeleteToken(l.ctx, &jwtx.DeleteToken_Request{
		TokenID:        in.TokenID,
		AccountID:      in.AccountID,
		AccessTerminal: in.AccessTerminal,
		SingleEnd:      l.svcCtx.Config.JWTXConfig.SingleEnd,
	})
	if err != nil {
		return nil, err
	}

	return &admin.DeleteToken_Response{}, nil
}
