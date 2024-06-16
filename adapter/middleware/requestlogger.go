package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anditakaesar/uwa-back/internal/env"
	"go.uber.org/zap"
)

func (m Middleware) IpLogging(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requesterIP := r.Header.Get(env.IPHeaderKey())
		if env.Env() == "development" && requesterIP == "" {
			requesterIP = "127.0.0.1"
		}

		httpRequest := map[string]interface{}{
			"requestMethod": strings.ToUpper(r.Method),
			"requestUrl":    r.URL.RequestURI(),
			"userAgent":     r.UserAgent(),
			"rawQuery":      r.URL.RawQuery,
			"query":         r.URL.Query(),
			"host":          r.Host,
		}

		if requesterIP != "" {
			_, err := m.IplogModel.UpdateCounter(requesterIP)
			if err != nil {
				m.Log.Error("[Middleware][IpLogging] UpdateCounter", err)
			} else {
				httpRequest["ipRequester"] = requesterIP
			}
		}

		go m.Log.Info(fmt.Sprintf("%s %s", r.Method, r.URL),
			zap.Any("httpRequest", httpRequest))
		h.ServeHTTP(w, r)
	}
}
