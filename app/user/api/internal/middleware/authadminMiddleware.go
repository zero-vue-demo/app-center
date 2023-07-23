package middleware

import (
	"app/user/api/internal/config"
	"common/response"
	"net/http"

	"github.com/5-say/go-tools/tools/random"
	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthAdminMiddleware struct {
}

func NewAuthAdminMiddleware() *AuthAdminMiddleware {
	return &AuthAdminMiddleware{}
}

func (m *AuthAdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			result = jwtx.GetMiddlewareResult(r.Context())
			c      = r.Context().Value("c").(config.Config)
		)

		// 分组校验
		if result.TokenGroup != c.Name+".admin" {
			appErr := response.Forbidden().Message("token group fail")
			httpx.ErrorCtx(r.Context(), w, appErr)
			return
		}

		// AccountID 防篡改
		if int64(result.AccountID) != random.Simple(c.SimpleRandom).Decode(result.RandomAccountID) {
			appErr := response.Unauthorized().Message("token was tampered with")
			httpx.ErrorCtx(r.Context(), w, appErr)
			return
		}

		next(w, r)
	}
}
