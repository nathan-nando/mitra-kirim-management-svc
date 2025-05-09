package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mitra-kirim-be-mgmt/http/rest/handler"
)

const v1 = "/api/v1"

type RouteConfig struct {
	App                  *echo.Echo
	SuggestionHandler    *handler.SuggestionHandler
	LocationHandler      *handler.LocationHandler
	ConfigurationHandler *handler.ConfigurationHandler
	TestimonialHandler   *handler.TestimonialHandler
	DashboardHandler     *handler.DashboardHandler
	UserHandler          *handler.UserHandler
	Middleware           *CustomMiddleware
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
	api.POST("/auth/register", r.UserHandler.Register)
	api.POST("/public/configuration", r.ConfigurationHandler.PublicConfig)
	api.POST("/public/suggestion", r.SuggestionHandler.Create)
}

func (r *RouteConfig) SetupAuthRoute() {
	api := r.App.Group(v1)
	api.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}))

	api.Use(r.Middleware.DevMode())
	api.Use(r.Middleware.AuthMiddleware())

	api.GET("/dashboard", r.DashboardHandler.Get)

	api.GET("/suggestion", r.SuggestionHandler.List)
	api.POST("/suggestion/reply", r.SuggestionHandler.ReplyEmail)

	api.GET("/location", r.LocationHandler.List)
	api.POST("/location", r.LocationHandler.Create)
	api.PATCH("/location", r.LocationHandler.Update)
	api.DELETE("/location/:id", r.LocationHandler.Delete)

	api.POST("/configuration/type", r.ConfigurationHandler.ListByTypes)
	api.PATCH("/configuration", r.ConfigurationHandler.UpdateConfiguration)
	api.PATCH("/configuration/appLogo", r.ConfigurationHandler.UpdateAppLogo)
	api.PATCH("/configuration/hero", r.ConfigurationHandler.UpdateHero)
	api.PATCH("/configuration/services", r.ConfigurationHandler.UpdateServices)

	api.GET("/testimonial", r.TestimonialHandler.List)
	api.POST("/testimonial", r.TestimonialHandler.Create)
	api.PATCH("/testimonial/slide", r.TestimonialHandler.UpdateSlide)
	api.DELETE("/testimonial", r.TestimonialHandler.Delete)

	api.PATCH("/user/picture", r.UserHandler.ChangeProfile)
	api.PATCH("/user/information", r.UserHandler.Update)
	api.GET("/user/information", r.UserHandler.Information)
}
