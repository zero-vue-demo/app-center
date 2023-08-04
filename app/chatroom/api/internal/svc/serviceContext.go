package svc

import (
	"app/chatroom/api/internal/config"
	"app/chatroom/api/internal/middleware"

	"github.com/5-say/zero-auth/public/jwtx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	JWTXRpc             jwtx.JwtxClient
	AuthAdminMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	jwtxRpc := jwtx.NewJwtxClient(zrpc.MustNewClient(c.JWTXRpc).Conn())
	return &ServiceContext{
		Config:              c,
		JWTXRpc:             jwtxRpc,
		AuthAdminMiddleware: middleware.NewAuthAdminMiddleware(jwtxRpc, "admin").Handle,
	}
}
