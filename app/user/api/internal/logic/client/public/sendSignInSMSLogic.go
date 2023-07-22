package public

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSignInSMSLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSignInSMSLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SendSignInSMSLogic {
	return &SendSignInSMSLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SendSignInSMSLogic) SendSignInSMS(req *types.Client_Public_SendSignInSMS_Request) (resp *types.Client_Public_SendSignInSMS_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
