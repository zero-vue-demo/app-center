package svc

import (
	"app/user/api/internal/config"
	"app/user/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config                 config.Config
	RefreshTokenMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		RefreshTokenMiddleware: middleware.NewRefreshTokenMiddleware().Handle,
	}
}
