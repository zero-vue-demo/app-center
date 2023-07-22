package user

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

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

func (l *EditPasswordLogic) EditPassword(req *types.Client_User_EditPassword_Request) (resp *types.Client_User_EditPassword_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
