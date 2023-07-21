package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSignInCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSignInCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSignInCaptchaLogic {
	return &GetSignInCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSignInCaptchaLogic) GetSignInCaptcha(req *types.Client_Public_GetSignInCaptcha_Request) (resp *types.Client_Public_GetSignInCaptcha_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
