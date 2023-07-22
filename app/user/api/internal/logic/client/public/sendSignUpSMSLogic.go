package public

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSignUpSMSLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSignUpSMSLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SendSignUpSMSLogic {
	return &SendSignUpSMSLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SendSignUpSMSLogic) SendSignUpSMS(req *types.Client_Public_SendSignUpSMS_Request) (resp *types.Client_Public_SendSignUpSMS_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
