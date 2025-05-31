package app_errors

import (
	"github.com/keanutaufan/anitrackr-server/pkg/app_error"
	"net/http"
)

var (
	ErrNotFound       = app_error.New(http.StatusNotFound, "Resource not found")
	ErrAlreadyExists  = app_error.New(http.StatusConflict, "Resource already exists")
	ErrForbidden      = app_error.New(http.StatusForbidden, "You are not allowed to perform this operation")
	ErrInternalServer = app_error.New(http.StatusInternalServerError, "Internal server error")
)
