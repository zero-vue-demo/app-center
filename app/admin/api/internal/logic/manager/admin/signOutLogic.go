package admin

import (
	"common/response"
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"

	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-services/public/jwtx"
	adminRpc "github.com/zero-vue-demo/app-center-public/rpc/admin"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignOutLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignOutLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SignOutLogic {
	return &SignOutLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SignOutLogic) SignOut(req *types.Manager_Admin_SignOut_Request) (resp *types.Manager_Admin_SignOut_Response, err error) {

	// 获取上下文信息
	middlewareResult := l.ctx.Value(jwtx.MIDDLEWARE_RESULT).(jwtx.MiddlewareResult)

	// 调用 RPC 移除 Token
	_, err = l.svcCtx.AdminRpc.DeleteToken(l.ctx, &adminRpc.DeleteToken_Request{
		TokenID:        middlewareResult.TokenID,
		AccountID:      middlewareResult.AccountID,
		AccessTerminal: middlewareResult.AccessTerminal,
	})
	if err != nil {
		rpcError := t.RPCErrorParse(err)
		return nil, response.Error(rpcError.Error())
	}

	return &types.Manager_Admin_SignOut_Response{}, nil
}
