package review_handler

import (
	review_request "github.com/keanutaufan/anitrackr-server/internal/domain/review/dto/request"
	review_usecase "github.com/keanutaufan/anitrackr-server/internal/domain/review/usecase"
	"github.com/keanutaufan/anitrackr-server/pkg/http_response"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"net/http"
)

type handler struct {
	reviewUseCase review_usecase.UseCase
}

func NewHandler(reviewUseCase review_usecase.UseCase) Handler {
	return &handler{
		reviewUseCase: reviewUseCase,
	}
}

func (h *handler) Store(c echo.Context) error {
	var req review_request.StoreReview
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.reviewUseCase.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Review created successfully!",
		Data:    response,
	})
}

func (h *handler) Index(c echo.Context) error {
	var req review_request.IndexReview
	if err := c.Bind(&req); err != nil {
		return err
	}

	response, meta, err := h.reviewUseCase.FindWithPagination(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Reviews retrieved successfully!",
		Data:    response,
		Meta:    meta,
	})
}

func (h *handler) Show(c echo.Context) error {
	id := cast.ToInt64(c.Param("reviewId"))
	response, err := h.reviewUseCase.FindOne(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Review retrieved successfully!",
		Data:    response,
	})
}

func (h *handler) Update(c echo.Context) error {
	var req review_request.UpdateReview
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	response, err := h.reviewUseCase.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Review updated successfully!",
		Data:    response,
	})
}

func (h *handler) Delete(c echo.Context) error {
	var req review_request.DeleteReview
	if err := c.Bind(&req); err != nil {
		return err
	}

	userId, ok := c.Get("userId").(int64)
	if !ok {
		return nil
	}
	req.UserId = userId

	err := h.reviewUseCase.Delete(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, http_response.Response{
		Success: true,
		Message: "Review deleted successfully!",
	})
}
