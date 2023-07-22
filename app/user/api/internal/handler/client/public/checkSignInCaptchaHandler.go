package public

import (
	"net/http"

	"app/user/api/internal/logic/client/public"
	"app/user/api/internal/svc"
	"app/user/api/internal/types"
	"common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CheckSignInCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.Client_Public_CheckSignInCaptcha_Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewCheckSignInCaptchaLogic(svcCtx, w, r)
		resp, err := l.CheckSignInCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Data(resp).OK(r.Context(), w)
		}
	}
}
