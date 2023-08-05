package admin

import (
	"context"
	"net/http"

	"app/chatroom/api/internal/svc"
	"app/chatroom/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.Manager_Admin_Demo_Request) (resp *types.Manager_Admin_Demo_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
