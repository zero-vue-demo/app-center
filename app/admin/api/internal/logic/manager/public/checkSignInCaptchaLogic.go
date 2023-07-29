package public

import (
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSignInCaptchaLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckSignInCaptchaLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *CheckSignInCaptchaLogic {
	return &CheckSignInCaptchaLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *CheckSignInCaptchaLogic) CheckSignInCaptcha(req *types.Manager_Public_CheckSignInCaptcha_Request) (resp *types.Manager_Public_CheckSignInCaptcha_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
