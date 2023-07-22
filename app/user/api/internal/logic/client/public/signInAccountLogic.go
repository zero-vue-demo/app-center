package public

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInAccountLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInAccountLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SignInAccountLogic {
	return &SignInAccountLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SignInAccountLogic) SignInAccount(req *types.Client_Public_SignInAccount_Request) (resp *types.Client_Public_SignInAccount_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
