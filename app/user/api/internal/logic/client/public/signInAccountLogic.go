package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInAccountLogic {
	return &SignInAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInAccountLogic) SignInAccount(req *types.Client_Public_SignInAccount_Request) (resp *types.Client_Public_SignInAccount_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
