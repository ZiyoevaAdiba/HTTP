package middleware

import (
	"go.uber.org/zap"
	"httpProject/pkg/Logger"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.L.Info("user authorization",
			zap.String("token:", r.Header.Get("token")),
		)

		token := r.Header.Get("Authorization")

		if token != "secret" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
