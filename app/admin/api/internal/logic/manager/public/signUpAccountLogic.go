package public

import (
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpAccountLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpAccountLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SignUpAccountLogic {
	return &SignUpAccountLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SignUpAccountLogic) SignUpAccount(req *types.Manager_Public_SignUpAccount_Request) (resp *types.Manager_Public_SignUpAccount_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
