package routes

import (
	"backend/app/middleware/jwt"
	"backend/app/server"
	"backend/pkg/api/v1/user"
	"backend/pkg/api/v1/user/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	svc user.Service
}

func NewHTTP(svc user.Service, err *echo.Group) {
	h := HTTP{svc: svc}

	g := err.Group("/user", jwt.JWTMiddleware)
	g.GET("/search", h.Search)
	g.POST("", h.CreateUser)
	g.PUT("/:user_id", h.UpdateUser)
	g.DELETE("/:user_id", h.DeleteUser)
}

// User godoc
// @Summary Search
// @Description Search in the system
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param q query string false "query"
// @Success 200 {object} models.ResponseCreateUser
// @Failure 400 {object} models.CommonResponse
// @Failure 401 {object} models.CommonResponse
// @Failure 500 {object} models.CommonResponse
// @Router /v1/user/search [get]
func (h *HTTP) Search(c echo.Context) error {
	q := new(models.QuerySearch)
	if err := c.Bind(q); err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	if err := q.Validate(); err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	ctx := c.Request().Context()
	result, err := h.svc.Search(ctx, *q)
	if err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	return server.ResponseStatusOK(c, "success", result, nil, nil)
}

// User godoc
// @Summary Create a new user
// @Description Register a new user in the system
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.RequestCreateUser true "User creation payload"
// @Success 201 {object} models.ResponseCreateUser
// @Failure 400 {object} models.CommonResponse
// @Failure 401 {object} models.CommonResponse
// @Failure 500 {object} models.CommonResponse
// @Router /v1/user [post]
func (h *HTTP) CreateUser(c echo.Context) error {
	requestBody := new(models.RequestCreateUser)
	if err := c.Bind(requestBody); err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	if err := requestBody.Validate(); err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	ctx := c.Request().Context()
	result, err := h.svc.CreateUser(ctx, *requestBody)
	if err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	return server.ResponseStatusOK(c, "success", result, nil, nil)
}

// User godoc
// @Summary Update user
// @Description Update data user in the system
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param request body models.RequestUpdateUser true "User update payload"
// @Success 200 {object} models.CommonResponse
// @Failure 400 {object} models.CommonResponse
// @Failure 401 {object} models.CommonResponse
// @Failure 500 {object} models.CommonResponse
// @Router /v1/user/{user_id} [put]
func (h *HTTP) UpdateUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	// Bind request body
	requestBody := new(models.RequestUpdateUser)
	if err := c.Bind(requestBody); err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	// Validate input (only validate fields that are provided)
	if err := requestBody.Validate(); err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	// Prepare update parameters
	requestParam := models.UpdateUser{IDUser: userId}

	if requestBody.Name != nil {
		requestParam.Name = requestBody.Name
	}
	if requestBody.Password != nil {
		requestParam.Password = requestBody.Password
	}

	// If no fields are provided, return an error
	if requestBody.Name == nil && requestBody.Password == nil {
		return server.ResponseStatusBadRequest(c, "No fields to update", nil, nil, nil)
	}

	// Call the service layer to update user
	ctx := c.Request().Context()
	err = h.svc.UpdateUser(ctx, requestParam)
	if err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	return server.ResponseStatusOK(c, "success", nil, nil, nil)
}

// User godoc
// @Summary Delete user
// @Description Delete a user from the system
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} models.CommonResponse
// @Failure 400 {object} models.CommonResponse
// @Failure 401 {object} models.CommonResponse
// @Failure 500 {object} models.CommonResponse
// @Router /v1/user/{user_id} [delete]
func (h *HTTP) DeleteUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	ctx := c.Request().Context()
	err = h.svc.DeleteUser(ctx, userId)
	if err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	return server.ResponseStatusOK(c, "success", nil, nil, nil)
}
