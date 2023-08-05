package user

import (
	"context"
	"net/http"

	"app/chatroom/api/internal/svc"
	"app/chatroom/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnterRoomLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnterRoomLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *EnterRoomLogic {
	return &EnterRoomLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *EnterRoomLogic) EnterRoom(req *types.Client_User_EnterRoom_Request) (resp *types.Client_User_EnterRoom_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
