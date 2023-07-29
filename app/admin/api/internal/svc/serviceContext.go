package svc

import (
	"app/admin/api/internal/config"
	"app/admin/api/internal/middleware"

	"github.com/zero-vue-demo/app-center-public/rpc/admin"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	AdminRpc            admin.AdminClient
	AuthAdminMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	adminRpc := admin.NewAdminClient(zrpc.MustNewClient(c.AdminRpc).Conn())
	return &ServiceContext{
		Config:              c,
		AdminRpc:            adminRpc,
		AuthAdminMiddleware: middleware.NewAuthAdminMiddleware(adminRpc).Handle,
	}
}
