package admin

import (
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminListLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminListLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *GetAdminListLogic {
	return &GetAdminListLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *GetAdminListLogic) GetAdminList(req *types.Manager_Admin_GetAdminList_Request) (resp *types.Manager_Admin_GetAdminList_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
