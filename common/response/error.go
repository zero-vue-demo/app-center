package response

// CustomError ..
type CustomError struct {
	Body struct {
		Meta struct {
			HTTPStatusCode   int    `json:"http_status_code"`   // HTTP 状态码
			SystemStatusCode int    `json:"system_status_code"` // 系统 状态码
			Message          string `json:"message"`            // 响应附加消息
		} `json:"meta"`

		Data struct {
			Errors []ErrorsItem `json:"errors"`
		} `json:"data"`
	}
}

// ErrorsItem ..
type ErrorsItem struct {
	Field   string `json:"field"`   // 错误字段
	Message string `json:"message"` // 错误信息
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 实现 原生 error 接口

// CustomError ..
func (e *CustomError) Error() string {
	return e.Body.Meta.Message
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 实例化

// Error ..
func Error(message string) *CustomError {
	return (&CustomError{}).BadRequest(message)
}

// Set ..
func (e *CustomError) Set(httpStatusCode, systemStatusCode int, message string) *CustomError {
	e.Body.Meta.HTTPStatusCode = httpStatusCode
	e.Body.Meta.SystemStatusCode = systemStatusCode
	e.Body.Meta.Message = message
	return e
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 预定义状态码

// 400 Bad Request 客户端请求的语法错误，服务器无法理解
func (e *CustomError) BadRequest(message string) *CustomError {
	return e.Set(400, 400, message)
}

// 401 Unauthorized 请求要求用户的身份认证
func Unauthorized() *CustomError {
	return (&CustomError{}).Set(401, 401, "Unauthorized")
}

// 403 Forbidden 服务器理解请求客户端的请求，但是拒绝执行此请求
func Forbidden() *CustomError {
	return (&CustomError{}).Set(403, 403, "Forbidden")
}

// 404 Not Found 服务器无法根据客户端的请求找到资源
func NotFound() *CustomError {
	return (&CustomError{}).Set(404, 404, "Not Found")
}

// 503 Service Unavailable 由于超载或系统维护，服务器暂时的无法处理客户端的请求。延时的长度可包含在服务器的 Retry-After 头信息中
func ServiceUnavailable() *CustomError {
	return (&CustomError{}).Set(503, 503, "Service Unavailable")
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 修改实例

// SYS ..
func (e *CustomError) SYS(systemStatusCode int) *CustomError {
	e.Body.Meta.SystemStatusCode = systemStatusCode
	return e
}

// Message ..
func (e *CustomError) Message(message string) *CustomError {
	e.Body.Meta.Message = message
	return e
}

// Append .. 附加错误信息
func (e *CustomError) Append(field, message string) *CustomError {
	e.Body.Data.Errors = append(e.Body.Data.Errors, ErrorsItem{
		Field:   field,
		Message: message,
	})
	return e
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 其它方法

// Count .. 统计附加的错误信息数量
func (e *CustomError) Count() int {
	return len(e.Body.Data.Errors)
}
