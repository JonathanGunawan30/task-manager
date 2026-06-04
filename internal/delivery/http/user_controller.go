package http

import (
	"jonathangunawan30/task-manager/internal/delivery/http/helper"
	"jonathangunawan30/task-manager/internal/entity"
	"jonathangunawan30/task-manager/internal/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	log         *logrus.Logger
	validate    *validator.Validate
	userUsecase usecase.UserUsecase
}

func NewUserController(log *logrus.Logger, validate *validator.Validate, userUsecase usecase.UserUsecase) *UserController {
	return &UserController{log: log, validate: validate, userUsecase: userUsecase}
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	userRequest := &entity.UserRegisterRequest{}
	err := helper.ReadFromRequestBody(r, userRequest)
	if err != nil {
		u.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(u.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}
	if err = u.validate.Struct(userRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(u.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userResponse, err := u.userUsecase.Register(r.Context(), userRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(u.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}
	status := http.StatusCreated
	helper.WriteJSON(u.log, w, status, entity.WebResponse[*entity.UserResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   userResponse,
	})
}

func (u *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	userResponses, err := u.userUsecase.GetAll(r.Context())
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(u.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(u.log, w, status, entity.WebResponse[[]*entity.UserResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   userResponses,
	})
}
