package public

import (
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSignUpCaptchaLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckSignUpCaptchaLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *CheckSignUpCaptchaLogic {
	return &CheckSignUpCaptchaLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *CheckSignUpCaptchaLogic) CheckSignUpCaptcha(req *types.Manager_Public_CheckSignUpCaptcha_Request) (resp *types.Manager_Public_CheckSignUpCaptcha_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
