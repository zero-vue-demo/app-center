package user

import (
	"common/response"
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-services/public/jwtx"
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
	// 调用 RPC 移除 Token
	_, err = l.svcCtx.JWTXRpc.DeleteToken(l.ctx, &jwtx.DeleteToken_Request{
		TokenID: l.ctx.Value("tokenID").(uint64),
	})
	if err != nil {
		return nil, response.Error(t.RPCErrorParse(err).PublicMessage)
	}

	return &types.Client_User_SignOut_Response{}, nil
}
