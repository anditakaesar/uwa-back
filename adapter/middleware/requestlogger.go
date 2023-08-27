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
		if env.Env() == "development" && requesterIP == "" {
			requesterIP = "127.0.0.1"
		}

		if requesterIP != "" {
			_, err := m.IplogModel.UpdateCounter(requesterIP)
			if err != nil {
				m.Log.Error("[Middleware][IpLogging] UpdateCounter", err)
			}
		}

		go m.Log.Info(fmt.Sprintf("%s %s", r.Method, r.URL),
			zap.Any("ip", requesterIP),
			zap.Any("headers", r.Header))
		h.ServeHTTP(w, r)
	}
}
