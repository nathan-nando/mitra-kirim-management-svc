package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mitra-kirim-be-mgmt/http/rest/handler"
)

const v1 = "/api/v1"

type RouteConfig struct {
	App               *echo.Echo
	SuggestionHandler *handler.SuggestionHandler
	LocationHandler   *handler.LocationHandler
	UserHandler       *handler.UserHandler
	Middleware        *CustomMiddleware
}

func (r *RouteConfig) Setup() {
	r.App.Validator = &CustomValidator{Validator: validator.New()}
	r.SetupGuestRoute()
	r.SetupAuthRoute()
}

func (r *RouteConfig) SetupGuestRoute() {
	r.App.GET("/health", handler.Health)

	api := r.App.Group(v1)
	api.Use(r.Middleware.DevMode())
	api.POST("/auth/login", r.UserHandler.Login)
	api.POST("/auth/refresh", r.UserHandler.Refresh)
	api.POST("/auth/logout", handler.Health)
}

func (c *RouteConfig) SetupAuthRoute() {
	api := c.App.Group(v1)
	api.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}))

	api.Use(c.Middleware.DevMode())
	//api.Use(c.Middleware.AuthMiddleware())

	api.GET("/suggestion", c.SuggestionHandler.List)
	api.POST("/suggestion", c.SuggestionHandler.Create)
	api.POST("/suggestion/reply", c.SuggestionHandler.ReplyEmail)

	api.GET("/location", c.LocationHandler.List)
	api.POST("/location", c.LocationHandler.Create)
	api.PATCH("/location", c.LocationHandler.Update)
	api.DELETE("/location/:id", c.LocationHandler.Delete)
}
