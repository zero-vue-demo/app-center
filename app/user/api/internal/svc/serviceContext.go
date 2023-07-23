package svc

import (
	"app/user/api/internal/config"
	"app/user/api/internal/middleware"

	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	JWTXRpc             jwtx.JwtxClient
	JWTXMiddleware      rest.Middleware
	AuthUserMiddleware  rest.Middleware
	AuthAdminMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		JWTXRpc:             jwtx.NewJwtxClient(zrpc.MustNewClient(c.JWTXRpc).Conn()),
		JWTXMiddleware:      middleware.NewJWTXMiddleware(c).Handle,
		AuthUserMiddleware:  middleware.NewAuthUserMiddleware().Handle,
		AuthAdminMiddleware: middleware.NewAuthAdminMiddleware().Handle,
	}
}
