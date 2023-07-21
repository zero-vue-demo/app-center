package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSignUpCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckSignUpCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSignUpCaptchaLogic {
	return &CheckSignUpCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckSignUpCaptchaLogic) CheckSignUpCaptcha(req *types.Client_Public_CheckSignUpCaptcha_Request) (resp *types.Client_Public_CheckSignUpCaptcha_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
