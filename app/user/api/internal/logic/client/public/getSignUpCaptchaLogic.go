package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSignUpCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSignUpCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSignUpCaptchaLogic {
	return &GetSignUpCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSignUpCaptchaLogic) GetSignUpCaptcha(req *types.Client_Public_GetSignUpCaptcha_Request) (resp *types.Client_Public_GetSignUpCaptcha_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
