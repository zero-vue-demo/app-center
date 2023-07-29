package public

import (
	"common/response"
	"context"
	"net/http"

	"app/admin/api/internal/svc"
	"app/admin/api/internal/types"
	"app/admin/db/dao"

	"github.com/5-say/go-tools/tools/ip"
	"github.com/5-say/go-tools/tools/password"
	"github.com/5-say/go-tools/tools/t"
	adminRpc "github.com/zero-vue-demo/app-center-public/rpc/admin"
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

func (l *SignInAccountLogic) SignInAccount(req *types.Manager_Public_SignInAccount_Request) (resp *types.Manager_Public_SignInAccount_Response, err error) {
	// 校验图形验证码

	// 查找用户
	adminQuery := dao.Common().Admin
	admin, err := adminQuery.Where(adminQuery.Account.Eq(req.Account)).First()
	if err != nil {
		l.Logger.Error(err)
		return nil, response.Error("请输入正确的账号")
	}

	// 校验密码
	err = password.Compare(req.Password, admin.Password)
	if err != nil {
		return nil, response.Error("密码错误")
	}

	// 校验成功 - - - - - - - -

	// 获取登录终端
	terminal := ""

	// 请求 rpc 生成 token
	rpcResp, err := l.svcCtx.AdminRpc.MakeToken(l.ctx, &adminRpc.MakeToken_Request{
		AccessTerminal: terminal,
		AccountID:      admin.ID,
		RequestIP:      ip.GetRequestIP(l.r),
	})
	if err != nil {
		rpcError := t.RPCErrorParse(err)
		return nil, response.Error(rpcError.Error())
	}

	return &types.Manager_Public_SignInAccount_Response{
		Token: rpcResp.Token,
	}, nil
}
