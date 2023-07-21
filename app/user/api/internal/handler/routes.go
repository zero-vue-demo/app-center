// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	clientpublic "app/user/api/internal/handler/client/public"
	clientuser "app/user/api/internal/handler/client/user"
	manageradmin "app/user/api/internal/handler/manager/admin"
	"app/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/sign-up/captcha",
				Handler: clientpublic.GetSignUpCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/sign-up/check-captcha",
				Handler: clientpublic.CheckSignUpCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sign-up/account",
				Handler: clientpublic.SignUpAccountHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sign-up/mobile",
				Handler: clientpublic.SignUpMobileHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/sign-up/send-sms",
				Handler: clientpublic.SendSignUpSMSHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/sign-in/captcha",
				Handler: clientpublic.GetSignInCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sign-in/check-captcha",
				Handler: clientpublic.CheckSignInCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sign-in/account",
				Handler: clientpublic.SignInAccountHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sign-in/mobile",
				Handler: clientpublic.SignInMobileHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/sign-in/send-sms",
				Handler: clientpublic.SendSignInSMSHandler(serverCtx),
			},
		},
		rest.WithPrefix("/client/public"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.RefreshTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/sign-out",
					Handler: clientuser.SignOutHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/password",
					Handler: clientuser.EditPasswordHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/client/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.RefreshTokenMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/users",
					Handler: manageradmin.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/users/:id",
					Handler: manageradmin.GetUserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/manager/admin"),
	)
}
