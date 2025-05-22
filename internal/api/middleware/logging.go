package custommiddleware

import (
	"net/http"

	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId, _ := utils.GenerateUUIDv7()
		logger := log.BaseLogger.WithFields(
			"request_id",
			requestId,
			"remote_addr",
			r.RemoteAddr,
			"method",
			r.Method,
			"url",
			r.URL.String(),
		)
		ctx := logger.WithCtx(r.Context())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
