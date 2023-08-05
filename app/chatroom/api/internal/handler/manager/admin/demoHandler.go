package admin

import (
	"net/http"

	"app/chatroom/api/internal/logic/manager/admin"
	"app/chatroom/api/internal/svc"
	"app/chatroom/api/internal/types"
	"common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.Manager_Admin_Demo_Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := admin.NewDemoLogic(svcCtx, w, r)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Data(resp).OK(r.Context(), w)
		}
	}
}
