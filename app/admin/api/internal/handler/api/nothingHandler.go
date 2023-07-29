package api

import (
	"net/http"

	"app/admin/api/internal/logic/api"
	"app/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func NothingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := api.NewNothingLogic(svcCtx, w, r)
		err := l.Nothing()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
