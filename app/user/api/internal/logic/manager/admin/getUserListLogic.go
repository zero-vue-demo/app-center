package admin

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.Manager_Admin_GetUserList_Request) (resp *types.Manager_Admin_GetUserList_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
