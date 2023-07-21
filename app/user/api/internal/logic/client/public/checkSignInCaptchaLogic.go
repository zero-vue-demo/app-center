package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSignInCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckSignInCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSignInCaptchaLogic {
	return &CheckSignInCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckSignInCaptchaLogic) CheckSignInCaptcha(req *types.Client_Public_CheckSignInCaptcha_Request) (resp *types.Client_Public_CheckSignInCaptcha_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
