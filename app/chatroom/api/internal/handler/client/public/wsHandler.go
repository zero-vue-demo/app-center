package public

import (
	"net/http"

	"app/chatroom/api/internal/logic/client/public"
	"app/chatroom/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := public.NewWsLogic(svcCtx, w, r)
		err := l.Ws()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
