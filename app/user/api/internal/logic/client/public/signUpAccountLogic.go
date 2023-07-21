package public

import (
	"context"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpAccountLogic {
	return &SignUpAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpAccountLogic) SignUpAccount(req *types.Client_Public_SignUpAccount_Request) (resp *types.Client_Public_SignUpAccount_Response, err error) {
	// todo: add your logic here and delete this line

	return
}
