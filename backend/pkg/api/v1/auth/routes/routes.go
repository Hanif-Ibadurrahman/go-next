package routes

import (
	"backend/app/server"
	"backend/pkg/api/v1/auth"
	"backend/pkg/api/v1/auth/models"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	svc auth.Service
}

func NewHTTP(svc auth.Service, err *echo.Group) {
	h := HTTP{svc: svc}

	g := err.Group("/auth")
	g.POST("/login", h.Login)
}

// CreateUser godoc
// @Summary Create login
// @Description Register a login in the system
// @Tags Login
// @Accept json
// @Produce json
// @Param request body models.RequestAuthLogin true "Login creation payload"
// @Success 201 {object} models.ResponseAuthLogin
// @Failure 400 {object} models.CommonResponse
// @Failure 401 {object} models.CommonResponse
// @Failure 500 {object} models.CommonResponse
// @Router /v1/auth/login [post]
func (h *HTTP) Login(c echo.Context) error {
	req := new(models.RequestAuthLogin)
	if err := c.Bind(req); err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	if err := req.Validate(); err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}

	ctx := c.Request().Context()
	result, err := h.svc.LoginAuth(ctx, *req)
	if err != nil {
		return server.ResponseStatusBadRequest(c, err.Error(), nil, nil, nil)
	}
	return server.ResponseStatusOK(c, "success", result, nil, nil)
}
