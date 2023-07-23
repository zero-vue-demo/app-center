package middleware

import (
	"fmt"
	"net/http"

	"github.com/5-say/zero-services/public/jwtx"
)

type AuthUserMiddleware struct {
}

func NewAuthUserMiddleware() *AuthUserMiddleware {
	return &AuthUserMiddleware{}
}

func (m *AuthUserMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		result := jwtx.GetMiddlewareResult(r.Context())
		fmt.Println(result.TokenGroup)

		// Passthrough to next handler if need
		next(w, r)
	}
}

// assignment copies lock value to jwtxResponse: github.com/5-say/zero-services/public/jwtx.CheckToken_Response contains google.golang.org/protobuf/internal/impl.MessageState contains sync.Mutex
