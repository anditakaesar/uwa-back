package way

import (
	"fmt"
	"net/http"

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

		if requesterIp != "" {
			_, err := ipLogRepo.UpdateCounter(requesterIp)
			if err != nil {
				internalLogger.Error("[NotFoundHandlerWithIpLogging] UpdateCounter", err)
			}
		}

		go internalLogger.Info(fmt.Sprintf("%s %s", r.Method, r.URL),
			zap.Any("ip", requesterIp),
			zap.Any("headers", r.Header))
		http.Error(w, "404 page not found", http.StatusNotFound)
	})
}
