package svc

import (
	"app/user/api/internal/config"
	"app/user/api/internal/middleware"

	"github.com/zero-vue-demo/app-center-public/rpc/user"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	UserRpc             user.UserClient
	AuthUserMiddleware  rest.Middleware
	AuthAdminMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRpc := user.NewUserClient(zrpc.MustNewClient(c.UserRpc).Conn())
	return &ServiceContext{
		Config:              c,
		UserRpc:             userRpc,
		AuthUserMiddleware:  middleware.NewAuthUserMiddleware(userRpc).Handle,
		AuthAdminMiddleware: middleware.NewAuthAdminMiddleware().Handle,
	}
}
