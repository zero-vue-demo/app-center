package public

import (
	"common/response"
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"
	"app/user/db/dao"
	"app/user/db/dao/model"

	"github.com/5-say/go-tools/tools/ip"
	"github.com/5-say/go-tools/tools/password"
	"github.com/5-say/go-tools/tools/random"
	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpAccountLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpAccountLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SignUpAccountLogic {
	return &SignUpAccountLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SignUpAccountLogic) SignUpAccount(req *types.Client_Public_SignUpAccount_Request) (resp *types.Client_Public_SignUpAccount_Response, err error) {
	var (
		userQuery = dao.Common().User
	)
	// 校验 图形验证码

	// 校验 账号格式

	// 校验 账号唯一
	if count, err := userQuery.Where(userQuery.Account.Eq(req.Account)).Count(); err != nil {
		l.Logger.Error(err)
		return nil, response.Error("注册失败")
	} else if count != 0 {
		return nil, response.Error("账号已注册")
	}

	// 密码生成
	hashPassword, err := password.Generate(req.Password)
	if err != nil {
		l.Logger.Error(err)
		return nil, response.Error("注册失败")
	}

	// 新增账号
	user := model.User{
		Account:  req.Account,
		Password: hashPassword,
	}
	err = userQuery.Create(&user)
	if err != nil {
		l.Logger.Error(err)
		return nil, response.Error("注册失败")
	}

	// 注册成功
	// 请求 rpc 生成 token
	rpcResponse, err := l.svcCtx.JWTXRpc.MakeToken(l.ctx, &jwtx.MakeToken_Request{
		AccessSecret:    l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire:    l.svcCtx.Config.Auth.AccessExpire,
		Group:           l.svcCtx.Config.Name + ".user",
		AccountID:       user.ID,
		RandomAccountID: random.Simple(l.svcCtx.Config.SimpleRandom).Encode(int64(user.ID), 8),
		RequestIP:       ip.GetRequestIP(l.r),
	})
	if err != nil {
		return nil, response.Error(t.RPCErrorParse(err).PublicMessage)
	}

	return &types.Client_Public_SignUpAccount_Response{
		Token: rpcResponse.Token,
	}, nil
}
