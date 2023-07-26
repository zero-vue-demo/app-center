package middleware

import (
	"common/response"
	"net/http"

	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthUserMiddleware struct {
	jwtxConfig jwtx.Config
	jwtxClient jwtx.JwtxClient
}

func NewAuthUserMiddleware(jwtxConfig jwtx.Config, jwtxClient jwtx.JwtxClient) *AuthUserMiddleware {
	return &AuthUserMiddleware{
		jwtxConfig: jwtxConfig,
		jwtxClient: jwtxClient,
	}
}

func (m *AuthUserMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 调用 jwtx-rpc 中间件
		newCtx, middlewareResult, rpcError := jwtx.Middleware(r.Header.Get("Authorization"), r, m.jwtxConfig, m.jwtxClient)

		// token 校验失败
		if rpcError != nil {

			// 内部错误信息日志
			// println(rpcError.PrivateMessage)

			// 加工错误对象
			appErr := response.Unauthorized().Message(rpcError.Error())

			// 为响应对象写入错误信息
			httpx.ErrorCtx(r.Context(), w, appErr)

		} else {

			// 响应头填充刷新后的 token
			w.Header().Add("NEW-TOKEN", middlewareResult.NewToken)

			// Passthrough to next handler if need
			next(w, r.WithContext(newCtx))
		}

	}
}
