package way

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anditakaesar/uwa-back/adapter/models/iplog"
	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/anditakaesar/uwa-back/internal/log"
	"go.uber.org/zap"
)

func NotFoundHandlerWithIpLogging(internalLogger log.LoggerInterface, ipLogRepo iplog.IplogModelInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requesterIp := r.Header.Get(env.IPHeaderKey())
		if env.Env() == "development" && requesterIp == "" {
			requesterIp = "127.0.0.1"
		}

		httpRequest := map[string]interface{}{
			"requestMethod": strings.ToUpper(r.Method),
			"requestUrl":    r.URL.RequestURI(),
			"userAgent":     r.UserAgent(),
			"rawQuery":      r.URL.RawQuery,
			"query":         r.URL.Query(),
			"host":          r.Host,
		}

		if requesterIp != "" {
			_, err := ipLogRepo.UpdateCounter(requesterIp)
			if err != nil {
				internalLogger.Error("[NotFoundHandlerWithIpLogging] UpdateCounter", err)
			} else {
				httpRequest["ipRequester"] = requesterIp
			}
		}

		go internalLogger.Info(fmt.Sprintf("%s %s", r.Method, r.URL),
			zap.Any("httpRequest", httpRequest))
		http.Error(w, "404 page not found", http.StatusNotFound)
	})
}
