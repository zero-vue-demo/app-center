package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSignInSMSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSignInSMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSignInSMSLogic {
	return &SendSignInSMSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSignInSMSLogic) SendSignInSMS(req *types.Client_Public_SendSignInSMS_Request) (resp *types.Client_Public_SendSignInSMS_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
