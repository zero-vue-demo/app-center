package svc

import (
	"app/user/api/internal/config"
	"app/user/api/internal/middleware"

	"github.com/5-say/zero-auth/public/jwtx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	JWTXRpc             jwtx.JwtxClient
	AuthUserMiddleware  rest.Middleware
	AuthAdminMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	jwtxRpc := jwtx.NewJwtxClient(zrpc.MustNewClient(c.JWTXRpc).Conn())
	return &ServiceContext{
		Config:              c,
		JWTXRpc:             jwtxRpc,
		AuthUserMiddleware:  middleware.NewAuthUserMiddleware(jwtxRpc, "user").Handle,
		AuthAdminMiddleware: middleware.NewAuthAdminMiddleware(jwtxRpc, "admin").Handle,
	}
}
