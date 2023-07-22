package public

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpMobileLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpMobileLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SignUpMobileLogic {
	return &SignUpMobileLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SignUpMobileLogic) SignUpMobile(req *types.Client_Public_SignUpMobile_Request) (resp *types.Client_Public_SignUpMobile_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
