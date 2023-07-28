package config

import (
	"github.com/5-say/go-tools/tools/db"
	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB         db.Config
	JWTXRpc    zrpc.RpcClientConf
	JWTXConfig jwtx.Config
}
