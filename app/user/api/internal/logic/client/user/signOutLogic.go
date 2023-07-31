package user

import (
	"common/response"
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-auth/public/jwtx"
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

func (l *SignOutLogic) SignOut(req *types.Client_User_SignOut_Request) (resp *types.Client_User_SignOut_Response, err error) {

	// 获取上下文信息
	middlewareResult := jwtx.GetValue(l.ctx)

	// 调用 RPC 移除 Token
	_, rpcError := l.svcCtx.JWTXRpc.DeleteToken(l.ctx, &jwtx.DeleteToken_Request{
		Group:     "user",
		TokenID:   middlewareResult.TokenID,
		AccountID: middlewareResult.AccountID,
		Terminal:  middlewareResult.Terminal,
	})
	if rpcError != nil {
		err := t.RPCErrorParse(rpcError)
		return nil, response.Error(err.Error())
	}

	return &types.Client_User_SignOut_Response{}, nil
}
