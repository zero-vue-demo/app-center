package admin

import (
	"common/response"
	"context"
	"net/http"

	"app/user/api/internal/svc"
	"app/user/api/internal/types"
	"app/user/db/dao"
	"app/user/db/dao/query"

	"github.com/5-say/go-tools/tools/t"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gen/field"
)

type GetUserListLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.Manager_Admin_GetUserList_Request) (resp *types.Manager_Admin_GetUserList_Response, err error) {
	var (
		page        = t.Pointer(req.Page, 1)
		perPage     = t.Pointer(req.PerPage, 10)
		orderColumn = t.Pointer(req.OrderColumn, "id")
		orderType   = t.Pointer(req.OrderType, "desc")
	)

	var (
		total int64
		items []types.Manager_Admin_GetUserList_Response_Item
	)
	{
		q := dao.Common().User
		var do query.IUserDo

		// 排序条件
		orderCond, ok := l.UserOrderCond(orderColumn, orderType)
		if !ok {
			return nil, response.Error("无效的排序参数")
		}
		do = q.Order(orderCond)

		// where 条件
		if req.SearchAccountLike != nil {
			do = do.Where(q.Account.Like(*req.SearchAccountLike))
		}

		// 分页查询
		total, err = do.ScanByPage(&items, (page-1)*perPage, perPage)
		if err != nil {
			l.Logger.Error(err)
			return nil, response.Error("sql fail")
		}
	}

	return &types.Manager_Admin_GetUserList_Response{
		Total: total,
		Items: items,
	}, nil
}

func (l *GetUserListLogic) UserOrderCond(orderColumn string, orderType string) (orderCond field.Expr, ok bool) {
	q := dao.Common().User

	// 排序条件
	orderCond, ok = map[string]field.Expr{
		"id asc":  q.ID,
		"id desc": q.ID.Desc(),
	}[orderColumn+" "+orderType]

	return
}
