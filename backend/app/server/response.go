package server

import (
	"backend/app/common/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResponseStatusOK 200
func ResponseStatusOK(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusOK, models.CommonResponse{
		Error:         false,
		Status:        http.StatusOK,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusOK Image 200
func ResponseStatusImageOK(c echo.Context, msg string, imageData []byte, enc interface{}, meta interface{}) error {
	c.Response().Header().Set(echo.HeaderContentType, "image/png")
	return c.Blob(http.StatusOK, "image/png", imageData)
}

// ResponseStatusCreated 201
func ResponseStatusCreated(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusCreated, models.CommonResponse{
		Error:         false,
		Status:        http.StatusCreated,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusAccepted 202
func ResponseStatusAccepted(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusAccepted, models.CommonResponse{
		Error:         false,
		Status:        http.StatusAccepted,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusNoContent 204
func ResponseStatusNoContent(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusNoContent, models.CommonResponse{
		Error:         false,
		Status:        http.StatusNoContent,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusBadRequest 400
func ResponseStatusBadRequest(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusBadRequest, models.CommonResponse{
		Error:         false,
		Status:        http.StatusBadRequest,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusUnauthorized 401
func ResponseStatusUnauthorized(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusUnauthorized, models.CommonResponse{
		Error:         false,
		Status:        http.StatusUnauthorized,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusConflict 409
func ResponseStatusConflict(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusConflict, models.CommonResponse{
		Error:         false,
		Status:        http.StatusConflict,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusForbidden 403
func ResponseStatusForbidden(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusForbidden, models.CommonResponse{
		Error:         false,
		Status:        http.StatusForbidden,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusForbidden 404
func ResponseStatusNotFound(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusNotFound, models.CommonResponse{
		Error:         false,
		Status:        http.StatusNotFound,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusForbidden 422
func ResponseStatusUnprocessableEntity(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusUnprocessableEntity, models.CommonResponse{
		Error:         false,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusServiceUnavailable 503
func ResponseStatusServiceUnavailable(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusServiceUnavailable, models.CommonResponse{
		Error:         false,
		Status:        http.StatusServiceUnavailable,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}

// ResponseStatusInternalServerError 500
func ResponseStatusInternalServerError(c echo.Context, msg string, data interface{}, enc interface{}, meta interface{}) error {
	return c.JSON(http.StatusInternalServerError, models.CommonResponse{
		Error:         false,
		Status:        http.StatusInternalServerError,
		Data:          data,
		Message:       msg,
		EncryptedData: enc,
		Meta:          meta,
	})
}
