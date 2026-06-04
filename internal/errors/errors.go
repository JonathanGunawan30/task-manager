package errors

import "errors"

var (
	TaskNotFound    = errors.New("task not found")
	ProjectNotFound = errors.New("project not found")
)
