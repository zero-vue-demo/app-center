package api

import (
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoErrorLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoErrorLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *DemoErrorLogic {
	return &DemoErrorLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *DemoErrorLogic) DemoError(req *types.DemoError_Request) (resp *types.DemoError_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
