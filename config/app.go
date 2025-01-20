package config

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/http/rest/handler"
	"mitra-kirim-be-mgmt/http/rest/middleware"
	"mitra-kirim-be-mgmt/internal/suggestion/repository"
	"mitra-kirim-be-mgmt/internal/suggestion/service"
)

type AppConfig struct {
	Db        *gorm.DB
	App       *echo.Echo
	Log       *logrus.Logger
	Publisher *redis.Client
}

func BuildInternal(appCfg *AppConfig) {
	suggestionRepo := repository.NewSuggestion(appCfg.Db)

	suggestionSvc := service.NewService(suggestionRepo, appCfg.Publisher, appCfg.Log)

	suggestionHandler := handler.NewSuggestionHandler(suggestionSvc, appCfg.Log)

	appMiddleware := middleware.NewCustomMiddleware(appCfg.Log)

	routeConfig := middleware.RouteConfig{
		App:               appCfg.App,
		SuggestionHandler: suggestionHandler,
		Middleware:        appMiddleware,
	}

	routeConfig.Setup()
}
