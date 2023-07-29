package svc

import (
	"app/admin/rpc/internal/config"

	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	JWTXRpc jwtx.JwtxClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		JWTXRpc: jwtx.NewJwtxClient(zrpc.MustNewClient(c.JWTXRpc).Conn()),
	}
}
