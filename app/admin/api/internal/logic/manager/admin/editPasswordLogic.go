package admin

import (
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditPasswordLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditPasswordLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *EditPasswordLogic {
	return &EditPasswordLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *EditPasswordLogic) EditPassword(req *types.Manager_Admin_EditPassword_Request) (resp *types.Manager_Admin_EditPassword_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
