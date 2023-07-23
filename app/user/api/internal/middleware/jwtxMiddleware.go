package middleware

import (
	"app/user/api/internal/config"
	"common/response"

	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewJWTXMiddleware(c config.Config) *jwtx.Middleware {
	var (
		// 错误包装函数，需返回包装好的项目错误对象
		errorWraper = func(err error) error {

			// 解析 rpc error
			e := t.RPCErrorParse(err)

			// 内部错误信息日志
			// println(e.PrivateMessage)

			// 返回加工好的错误对象
			return response.Unauthorized().Message(e.PublicMessage)
		}

		jwtxConfig = c.Auth                                                   // jwtx 配置
		jwtxRpc    = jwtx.NewJwtxClient(zrpc.MustNewClient(c.JWTXRpc).Conn()) // jwtx rpc 客户端
	)
	return jwtx.NewMiddleware(errorWraper, jwtxConfig, jwtxRpc)
}
