package public

import (
	"context"
	"net/http"

	"app/chatroom/api/internal"
	"app/chatroom/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type WsLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *WsLogic {
	return &WsLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws() error {
	internal.ServeWs(l.svcCtx.Hub, l.w, l.r)

	return nil
}
