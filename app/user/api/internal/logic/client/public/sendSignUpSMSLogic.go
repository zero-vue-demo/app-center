package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSignUpSMSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSignUpSMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSignUpSMSLogic {
	return &SendSignUpSMSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSignUpSMSLogic) SendSignUpSMS(req *types.Client_Public_SendSignUpSMS_Request) (resp *types.Client_Public_SendSignUpSMS_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
