package admin

import (
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminInfoLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminInfoLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *GetAdminInfoLogic {
	return &GetAdminInfoLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *GetAdminInfoLogic) GetAdminInfo(req *types.Manager_Admin_GetAdminInfo_Request) (resp *types.Manager_Admin_GetAdminInfo_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
