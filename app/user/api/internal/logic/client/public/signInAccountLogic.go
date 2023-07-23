package public

import (
	"common/response"
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"
	"app/user/db/dao"

	"github.com/5-say/go-tools/tools/ip"
	"github.com/5-say/go-tools/tools/password"
	"github.com/5-say/go-tools/tools/random"
	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-services/public/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
)

type SignInAccountLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInAccountLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SignInAccountLogic {
	return &SignInAccountLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SignInAccountLogic) SignInAccount(req *types.Client_Public_SignInAccount_Request) (resp *types.Client_Public_SignInAccount_Response, err error) {
	// 校验图形验证码

	// 查找用户
	userQuery := dao.Common().User
	user, err := userQuery.Where(userQuery.Account.Eq(req.Account)).First()
	if err != nil {
		l.Logger.Error(err)
		return nil, response.Error("请输入正确的账号")
	}

	// 校验密码
	err = password.Compare(req.Password, user.Password)
	if err != nil {
		return nil, response.Error("密码错误")
	}

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

	return &types.Client_Public_SignInAccount_Response{
		Token: rpcResponse.Token,
	}, nil
}
