package public

import (
	"common/response"
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"
	"app/admin/db/dao"
	"app/admin/db/dao/model"

	"github.com/5-say/go-tools/tools/ip"
	"github.com/5-say/go-tools/tools/password"
	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-auth/public/jwtx"
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

func (l *SignUpAccountLogic) SignUpAccount(req *types.Manager_Public_SignUpAccount_Request) (resp *types.Manager_Public_SignUpAccount_Response, err error) {
	var (
		adminQuery = dao.Common().Admin
	)
	// 校验 图形验证码

	// 校验 账号格式

	// 校验 账号唯一
	if count, err := adminQuery.Where(adminQuery.Account.Eq(req.Account)).Count(); err != nil {
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
	admin := model.Admin{
		Account:  req.Account,
		Password: hashPassword,
	}
	err = adminQuery.Create(&admin)
	if err != nil {
		l.Logger.Error(err)
		return nil, response.Error("注册失败")
	}

	// 注册成功 - - - - - - - -

	// 获取登录终端
	terminal := ""

	// 请求 rpc 生成 token
	rpcResp, err := l.svcCtx.JWTXRpc.MakeToken(l.ctx, &jwtx.MakeToken_Request{
		Group:     "admin",
		Terminal:  terminal,
		AccountID: admin.ID,
		RequestIP: ip.GetRequestIP(l.r),
	})
	if err != nil {
		rpcError := t.RPCErrorParse(err)
		return nil, response.Error(rpcError.Error())
	}

	return &types.Manager_Public_SignUpAccount_Response{
		Token: rpcResp.Token,
	}, nil
}
