package middleware

import (
	"fmt"
	"net/http"

	"github.com/anditakaesar/uwa-back/internal/env"
	"go.uber.org/zap"
)

func (m Middleware) IpLogging(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requesterIP := r.Header.Get(env.IPHeaderKey())
		m.Log.Info(fmt.Sprintf("%s %s", r.Method, r.URL), zap.Any("ip", requesterIP))
		h.ServeHTTP(w, r)
	}
}
