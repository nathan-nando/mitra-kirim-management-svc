package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mitra-kirim-be-mgmt/http/rest/handler"
)

const v1 = "/api/v1"

type RouteConfig struct {
	App               *echo.Echo
	SuggestionHandler *handler.SuggestionHandler
	Middleware        *CustomMiddleware
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.GET("/health", handler.Health)
}

func (c *RouteConfig) SetupAuthRoute() {
	api := c.App.Group(v1)
	api.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}))
	api.Use(c.Middleware.DevMode())
	api.GET("/suggestion", c.SuggestionHandler.List)
}
