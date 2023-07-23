package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"app/user/api/internal/config"
	"app/user/api/internal/handler"
	"app/user/api/internal/svc"
	"app/user/db/dao"
	"common/response"

	"github.com/5-say/go-tools/tools/db"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(
		c.RestConf,                      // 配置
		rest.WithCors(c.CorsOrigins...), // 跨域
	)
	defer server.Stop()

	// 初始化公共数据库连接
	dao.SetCommon(c.DB, db.LogWriter())

	// 注册全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "c", c)
			next(w, r.WithContext(ctx))
		}
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *response.CustomError:
			return e.Body.Meta.HTTPStatusCode, e.Body
		default:
			fmt.Println("- - - - - - - - - - - - - - - 未使用自定义错误 -")
			fmt.Println(err)
			fmt.Println("- - - - - - - - - - - - - - -")
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
