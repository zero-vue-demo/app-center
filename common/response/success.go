package response

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// CustomSuccess ..
type CustomSuccess struct {
	Meta struct {
		HTTPStatusCode   int    `json:"http_status_code"`   // HTTP 状态码
		SystemStatusCode int    `json:"system_status_code"` // 系统 状态码
		Message          string `json:"message"`            // 响应附加消息
	} `json:"meta"`

	Data interface{} `json:"data"`
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 实例化

// Data ..
func Data(resp interface{}) *CustomSuccess {
	v := &CustomSuccess{
		Data: resp,
	}
	v.Meta.HTTPStatusCode = 200
	v.Meta.SystemStatusCode = 200
	v.Meta.Message = "OK"
	return v
}

// OK ..
func (s *CustomSuccess) OK(ctx context.Context, w http.ResponseWriter) {
	httpx.OkJsonCtx(ctx, w, s)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 修改实例

// SYS ..
func (s *CustomSuccess) SYS(systemStatusCode int) *CustomSuccess {
	s.Meta.SystemStatusCode = systemStatusCode
	return s
}

// Message ..
func (s *CustomSuccess) Message(message string) *CustomSuccess {
	s.Meta.Message = message
	return s
}
