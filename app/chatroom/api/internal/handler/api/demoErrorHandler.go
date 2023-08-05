package api

import (
	"net/http"

	"app/chatroom/api/internal/logic/api"
	"app/chatroom/api/internal/svc"
	"app/chatroom/api/internal/types"
	"common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DemoErrorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.DemoError_Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := api.NewDemoErrorLogic(svcCtx, w, r)
		resp, err := l.DemoError(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Data(resp).OK(r.Context(), w)
		}
	}
}
