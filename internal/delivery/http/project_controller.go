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

type ProjectController struct {
	log            *logrus.Logger
	validate       *validator.Validate
	projectUsecase usecase.ProjectUsecase
}

func NewProjectController(log *logrus.Logger, validate *validator.Validate, projectUsecase usecase.ProjectUsecase) *ProjectController {
	return &ProjectController{log: log, validate: validate, projectUsecase: projectUsecase}
}

func (p *ProjectController) Create(w http.ResponseWriter, r *http.Request) {
	projectRequest := &entity.ProjectCreateRequest{}
	err := helper.ReadFromRequestBody(r, projectRequest)
	if err != nil {
		p.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	projectRequest.UserID = userID

	if err = p.validate.Struct(projectRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	projectResponse, err := p.projectUsecase.Create(r.Context(), projectRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusCreated
	helper.WriteJSON(p.log, w, status, entity.WebResponse[*entity.ProjectResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   projectResponse,
	})
}

func (p *ProjectController) GetAll(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	projectResponse, err := p.projectUsecase.GetAll(r.Context(), userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(p.log, w, status, entity.WebResponse[[]*entity.ProjectResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   projectResponse,
	})
}

func (p *ProjectController) GetDetail(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	projectIDStr := r.PathValue("projectID")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid project ID",
		})
		return
	}

	projectResponse, err := p.projectUsecase.GetDetail(r.Context(), projectID, userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(p.log, w, status, entity.WebResponse[*entity.ProjectResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   projectResponse,
	})
}

func (p *ProjectController) Update(w http.ResponseWriter, r *http.Request) {
	projectRequest := &entity.ProjectUpdateRequest{}
	err := helper.ReadFromRequestBody(r, projectRequest)
	if err != nil {
		p.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	projectIDStr := r.PathValue("projectID")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid project ID",
		})
		return
	}

	projectRequest.UserID = userID
	projectRequest.ID = projectID

	if err = p.validate.Struct(projectRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	projectResponse, err := p.projectUsecase.Update(r.Context(), projectRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(p.log, w, status, entity.WebResponse[*entity.ProjectResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   projectResponse,
	})
}

func (p *ProjectController) Delete(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	projectIDStr := r.PathValue("projectID")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid project ID",
		})
		return
	}

	err = p.projectUsecase.Delete(r.Context(), projectID, userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(p.log, w, status, entity.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(p.log, w, status, entity.WebResponse[any]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   nil,
	})
}
