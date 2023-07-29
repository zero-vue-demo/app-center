package public

import (
	"net/http"

	"app/admin/api/internal/logic/manager/public"
	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"
	"common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSignUpCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.Manager_Public_GetSignUpCaptcha_Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewGetSignUpCaptchaLogic(svcCtx, w, r)
		resp, err := l.GetSignUpCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Data(resp).OK(r.Context(), w)
		}
	}
}
