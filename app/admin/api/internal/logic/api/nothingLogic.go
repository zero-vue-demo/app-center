package api

import (
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type NothingLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNothingLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *NothingLogic {
	return &NothingLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *NothingLogic) Nothing() error {
	// todo: add your logic here and delete this line

	return nil
}
