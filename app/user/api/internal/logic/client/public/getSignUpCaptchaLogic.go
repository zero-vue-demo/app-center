package public

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSignUpCaptchaLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSignUpCaptchaLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *GetSignUpCaptchaLogic {
	return &GetSignUpCaptchaLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *GetSignUpCaptchaLogic) GetSignUpCaptcha(req *types.Client_Public_GetSignUpCaptcha_Request) (resp *types.Client_Public_GetSignUpCaptcha_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
