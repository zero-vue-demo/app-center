package admin

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.Manager_Admin_GetUserInfo_Request) (resp *types.Manager_Admin_GetUserInfo_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
