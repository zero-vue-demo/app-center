package user

import (
	"net/http"

	"app/user/api/internal/logic/client/user"
	"app/user/api/internal/svc"
	"app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignOutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Client_User_SignOut_Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSignOutLogic(r.Context(), svcCtx)
		resp, err := l.SignOut(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
