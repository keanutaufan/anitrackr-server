package list_handler

import (
	list_request "github.com/keanutaufan/anitrackr-server/internal/domain/list/dto/request"
	list_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/list/usecase"
	"github.com/keanutaufan/anitrackr-server/pkg/http_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	listUseCase list_usecase.UseCase
}

func NewHandler(listUseCase list_usecase.UseCase) Handler {
	return &handler{
		listUseCase: listUseCase,
	}
}

func (h *handler) Store(c echo.Context) error {
	var req list_request.StoreList
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.listUseCase.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "List created successfully!",
		Data:    response,
	})
}

func (h *handler) Show(c echo.Context) error {
	var req list_request.ShowList
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.listUseCase.FindOne(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "List retrieved successfully!",
		Data:    response,
	})
}

func (h *handler) Update(c echo.Context) error {
	var req list_request.UpdateList
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.listUseCase.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "List updated successfully!",
		Data:    response,
	})
}

func (h *handler) Delete(c echo.Context) error {
	var req list_request.DeleteList
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	err := h.listUseCase.Delete(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "List deleted successfully!",
	})
}
