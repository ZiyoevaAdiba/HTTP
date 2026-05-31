package middleware

import (
	"go.uber.org/zap"
	"httpProject/pkg/Logger"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.L.Info("Request data:",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("id", r.PathValue("id")),
		)

		next.ServeHTTP(w, r)
	})
}
