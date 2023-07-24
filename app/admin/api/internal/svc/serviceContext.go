package svc

import (
	"app/admin/api/internal/config"
	"app/admin/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	JWTXMiddleware      rest.Middleware
	AuthAdminMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		JWTXMiddleware:      middleware.NewJWTXMiddleware().Handle,
		AuthAdminMiddleware: middleware.NewAuthAdminMiddleware().Handle,
	}
}
