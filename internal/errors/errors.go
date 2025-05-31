package app_errors

import (
	"github.com/keanutaufan/anitrackr-server/pkg/app_error"
	"net/http"
)

var (
	ErrNotFound       = app_error.New(http.StatusNotFound, "Resource not found")
	ErrAlreadyExists  = app_error.New(http.StatusConflict, "Resource already exists")
	ErrInternalServer = app_error.New(http.StatusInternalServerError, "Internal server error")
)
