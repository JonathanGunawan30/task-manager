package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func AuthMiddleware(log *logrus.Logger, config *viper.Viper, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		secretKey := config.GetString("X_API_KEY")

		if apiKey != secretKey {
			log.Warnf("Unauthorized access attempt with API Key: %s", apiKey)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"code":401,"status":"Unauthorized","error":"Unauthorized"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}
