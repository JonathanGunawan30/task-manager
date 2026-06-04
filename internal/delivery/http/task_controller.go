package http

import (
	"jonathangunawan30/task-manager/internal/delivery/http/helper"
	"jonathangunawan30/task-manager/internal/entity"
	"jonathangunawan30/task-manager/internal/usecase"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type TaskController struct {
	log         *logrus.Logger
	validate    *validator.Validate
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(log *logrus.Logger, validate *validator.Validate, taskUsecase usecase.TaskUsecase) *TaskController {
	return &TaskController{log: log, validate: validate, taskUsecase: taskUsecase}
}

func (t *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	taskRequest := &entity.TaskCreateRequest{}
	err := helper.ReadFromRequestBody(r, taskRequest)
	if err != nil {
		t.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	taskRequest.UserID = userID

	if err = t.validate.Struct(taskRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	taskResponse, err := t.taskUsecase.Create(r.Context(), taskRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusCreated
	helper.WriteJSON(t.log, w, status, entity.WebResponse[*entity.TaskResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   taskResponse,
	})
}

func (t *TaskController) GetAll(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	taskResponse, err := t.taskUsecase.GetAll(r.Context(), userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(t.log, w, status, entity.WebResponse[[]*entity.TaskResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   taskResponse,
	})
}

func (t *TaskController) GetDetail(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	taskIDstr := r.PathValue("taskID")
	taskID, err := strconv.Atoi(taskIDstr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid task ID",
		})
	}

	taskResponse, err := t.taskUsecase.GetDetail(r.Context(), taskID, userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(t.log, w, status, entity.WebResponse[*entity.TaskResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   taskResponse,
	})
}

func (t *TaskController) Update(w http.ResponseWriter, r *http.Request) {
	taskRequest := &entity.TaskUpdateRequest{}
	err := helper.ReadFromRequestBody(r, taskRequest)
	if err != nil {
		t.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	taskID, err := strconv.Atoi(r.PathValue("taskID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	taskRequest.UserID = userID
	taskRequest.ID = taskID
	if err = t.validate.Struct(&taskRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
	}

	taskResponse, err := t.taskUsecase.Update(r.Context(), taskRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
	}

	status := http.StatusOK
	helper.WriteJSON(t.log, w, status, entity.WebResponse[*entity.TaskResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   taskResponse,
	})
}

func (t *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	taskIDstr := r.PathValue("taskID")
	taskID, err := strconv.Atoi(taskIDstr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid task ID",
		})
	}

	err = t.taskUsecase.Delete(r.Context(), taskID, userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(t.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(t.log, w, status, entity.WebResponse[*entity.TaskResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   nil,
	})
}
