package config

import (
	"fmt"
	deliveryHttp "jonathangunawan30/task-manager/internal/delivery/http"
	"jonathangunawan30/task-manager/internal/delivery/http/router"
	"jonathangunawan30/task-manager/internal/repository"
	repositoryImpl "jonathangunawan30/task-manager/internal/repository/impl"
	TxManager "jonathangunawan30/task-manager/internal/tx"
	usecaseImpl "jonathangunawan30/task-manager/internal/usecase/impl"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Bootstrap(config *viper.Viper, log *logrus.Logger, db *gorm.DB, validate *validator.Validate) {
	tx := TxManager.NewTx(db)

	base := repository.NewBaseRepository(db)

	userRepository := repositoryImpl.NewUserRepository(log, base)
	projectRepository := repositoryImpl.NewProjectRepository(log, base)
	taskRepository := repositoryImpl.NewTaskRepository(log, base)

	userUsecase := usecaseImpl.NewUserUsecase(log, tx, userRepository)
	projectUsecase := usecaseImpl.NewProjectUsecase(log, tx, projectRepository)
	taskUsecase := usecaseImpl.NewTaskUsecase(log, tx, taskRepository, projectRepository)

	userController := deliveryHttp.NewUserController(log, validate, userUsecase)
	projectController := deliveryHttp.NewProjectController(log, validate, projectUsecase)
	taskController := deliveryHttp.NewTaskController(log, validate, taskUsecase)

	r := router.NewRouter(log, config, *userController, *projectController, *taskController)

	port := config.GetInt("APP_PORT")
	if port == 0 {
		port = 3000
	}
	address := fmt.Sprintf(":%d", port)
	log.Infof("Server started at %s", address)
	err := http.ListenAndServe(address, r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
