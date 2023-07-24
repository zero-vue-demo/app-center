package api

import (
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SwaggerDocLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwaggerDocLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SwaggerDocLogic {
	return &SwaggerDocLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SwaggerDocLogic) SwaggerDoc() error {
	// todo: add your logic here and delete this line

	return nil
}
