package public

import (
	"net/http"

	"app/user/api/internal/logic/client/public"
	"app/user/api/internal/svc"
	"app/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CheckSignUpCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Client_Public_CheckSignUpCaptcha_Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewCheckSignUpCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.CheckSignUpCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
