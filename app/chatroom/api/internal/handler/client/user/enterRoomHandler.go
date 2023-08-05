package user

import (
	"net/http"

	"app/chatroom/api/internal/logic/client/user"
	"app/chatroom/api/internal/svc"
	"app/chatroom/api/internal/types"
	"common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EnterRoomHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.Client_User_EnterRoom_Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewEnterRoomLogic(svcCtx, w, r)
		resp, err := l.EnterRoom(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Data(resp).OK(r.Context(), w)
		}
	}
}
