package public

import (
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSignInCaptchaLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSignInCaptchaLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *GetSignInCaptchaLogic {
	return &GetSignInCaptchaLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *GetSignInCaptchaLogic) GetSignInCaptcha(req *types.Manager_Public_GetSignInCaptcha_Request) (resp *types.Manager_Public_GetSignInCaptcha_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
