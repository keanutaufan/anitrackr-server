package auth_handler

import (
	auth_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/auth/usecase"
	"github.com/keanutaufan/anitrackr-server/pkg/http_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	authUseCase auth_usecase.UseCase
}

func NewHandler(authUseCase auth_usecase.UseCase) Handler {
	return &handler{
		authUseCase: authUseCase,
	}
}
func (h *handler) Me(ctx echo.Context) error {
	userId, ok := ctx.Get("userId").(int64)
	if !ok {
		return nil
	}

	response, err := h.authUseCase.GetCurrentUser(ctx.Request().Context(), userId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Current user retrieved successfully!",
		Data:    response,
	})
}
