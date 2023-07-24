// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	api "app/admin/api/internal/handler/api"
	manageradmin "app/admin/api/internal/handler/manager/admin"
	managerpublic "app/admin/api/internal/handler/manager/public"
	"app/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: api.NothingHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/demo/error",
				Handler: api.DemoErrorHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/doc/swagger",
				Handler: api.SwaggerDocHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/sign-up/captcha",
				Handler: managerpublic.GetSignUpCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/sign-up/check-captcha",
				Handler: managerpublic.CheckSignUpCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sign-up/account",
				Handler: managerpublic.SignUpAccountHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/sign-in/captcha",
				Handler: managerpublic.GetSignInCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sign-in/check-captcha",
				Handler: managerpublic.CheckSignInCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sign-in/account",
				Handler: managerpublic.SignInAccountHandler(serverCtx),
			},
		},
		rest.WithPrefix("/manager/public"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JWTXMiddleware, serverCtx.AuthAdminMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/sign-out",
					Handler: manageradmin.SignOutHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/password",
					Handler: manageradmin.EditPasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/users",
					Handler: manageradmin.GetAdminListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/users/:id",
					Handler: manageradmin.GetAdminInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/manager/admin"),
	)
}
