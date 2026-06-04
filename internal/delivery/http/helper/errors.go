package helper

import (
	"errors"
	errDto "jonathangunawan30/task-manager/internal/errors"
	"net/http"
)

func DomainErrorToHTTPStatus(err error) int {
	switch {
	case errors.Is(err, errDto.TaskNotFound):
		return http.StatusNotFound
	case errors.Is(err, errDto.ProjectNotFound):
		return http.StatusNotFound
	case errors.Is(err, errDto.EmailAlreadyExists):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
