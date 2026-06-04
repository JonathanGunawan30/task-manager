package helper

import (
	"jonathangunawan30/task-manager/internal/delivery/http/middleware"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Auth(log *logrus.Logger, config *viper.Viper, h http.HandlerFunc) http.Handler {
	return middleware.AuthMiddleware(log, config, h)
}
