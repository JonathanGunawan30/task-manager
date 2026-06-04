package router

import (
	deliveryHttp "jonathangunawan30/task-manager/internal/delivery/http"
	"jonathangunawan30/task-manager/internal/delivery/http/helper"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewRouter(log *logrus.Logger, config *viper.Viper, userController deliveryHttp.UserController, projectController deliveryHttp.ProjectController, taskController deliveryHttp.TaskController) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/users/register", userController.Register)
	mux.HandleFunc("GET /api/users", userController.GetAll)

	mux.Handle("POST /api/projects", helper.Auth(log, config, projectController.Create))
	mux.Handle("GET /api/projects", helper.Auth(log, config, projectController.GetAll))
	mux.Handle("GET /api/projects/{projectID}", helper.Auth(log, config, projectController.GetDetail))
	mux.Handle("PUT /api/projects/{projectID}", helper.Auth(log, config, projectController.Update))
	mux.Handle("DELETE /api/projects/{projectID}", helper.Auth(log, config, projectController.Delete))

	mux.Handle("POST /api/tasks", helper.Auth(log, config, taskController.Create))
	mux.Handle("GET /api/tasks", helper.Auth(log, config, taskController.GetAll))
	mux.Handle("GET /api/tasks/{taskID}", helper.Auth(log, config, taskController.GetDetail))
	mux.Handle("PUT /api/tasks/{taskID}", helper.Auth(log, config, taskController.Update))
	mux.Handle("DELETE /api/tasks/{taskID}", helper.Auth(log, config, taskController.Delete))

	return mux
}
