package middleware

import (
	"common/response"
	"context"
	"net/http"

	"github.com/5-say/go-tools/tools/ip"
	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zero-vue-demo/app-center-public/rpc/admin"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthAdminMiddleware struct {
	adminClient admin.AdminClient
}

func NewAuthAdminMiddleware(adminClient admin.AdminClient) *AuthAdminMiddleware {
	return &AuthAdminMiddleware{
		adminClient: adminClient,
	}
}

func (m *AuthAdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 调用 admin-rpc 中间件
		result, err := m.adminClient.CheckToken(r.Context(), &admin.CheckToken_Request{
			RequestIP:    ip.GetRequestIP(r),
			RequestToken: r.Header.Get("Authorization"),
		})

		// token 校验失败
		if err != nil {

			rpcError := t.RPCErrorParse(err)

			// 内部错误信息日志
			// println(rpcError.PrivateMessage)

			// 加工错误对象
			appErr := response.Unauthorized().Message(rpcError.Error())

			// 为响应对象写入错误信息
			httpx.ErrorCtx(r.Context(), w, appErr)

		} else {

			// 响应头填充刷新后的 token
			w.Header().Add("NEW-TOKEN", result.NewToken)

			// 附加上下文
			newCtx := context.WithValue(r.Context(), jwtx.MIDDLEWARE_RESULT, jwtx.MiddlewareResult{
				TokenID:             result.TokenID,
				AccountID:           result.AccountID,
				AccessTerminal:      result.AccessTerminal,
				MakeTokenIP:         result.MakeTokenIP,
				ExpirationTimestamp: result.ExpirationTimestamp,
				NewToken:            result.NewToken,
			})

			// Passthrough to next handler if need
			next(w, r.WithContext(newCtx))
		}

	}
}
