package api

import (
	"backend/app/config"
	"backend/app/server"

	usersRoutes "backend/pkg/api/v1/user/routes"
	userUC "backend/pkg/api/v1/user/usecase"

	authRoutes "backend/pkg/api/v1/auth/routes"
	authUC "backend/pkg/api/v1/auth/usecase"

	_ "backend/docs"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	e := server.InitEcho()

	e.GET("/docs/*", echoSwagger.WrapHandler)

	routeV1 := e.Group("/v1")
	routeV1.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(config.GetConfig().RateLimit))))
	authRoutes.NewHTTP(authUC.Initialize(db), routeV1)
	usersRoutes.NewHTTP(userUC.Initialize(db), routeV1)
	server.Start(e)
}
