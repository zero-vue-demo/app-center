package public

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInMobileLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInMobileLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SignInMobileLogic {
	return &SignInMobileLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SignInMobileLogic) SignInMobile(req *types.Client_Public_SignInMobile_Request) (resp *types.Client_Public_SignInMobile_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
