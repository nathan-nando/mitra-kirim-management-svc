package config

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/http/rest/handler"
	"mitra-kirim-be-mgmt/http/rest/middleware"
	"mitra-kirim-be-mgmt/internal/suggestion/repository"
	"mitra-kirim-be-mgmt/internal/suggestion/service"

	//serviceCache "mitra-kirim-be-mgmt/internal/gateways/cache/service"
	servicePub "mitra-kirim-be-mgmt/internal/gateways/publisher/service"
)

type AppConfig struct {
	Db        *gorm.DB
	App       *echo.Echo
	Log       *logrus.Logger
	Publisher *RedisConfig
	Cache     *RedisConfig
	Config    *Config
}

func BuildInternal(appCfg *AppConfig) {
	//newCache := serviceCache.NewCache(appCfg.Redis)
	publisher := servicePub.NewPublisher(appCfg.Publisher.Client, appCfg.Log, appCfg.Config.RedisMaxRetry)

	suggestionRepo := repository.NewSuggestion(appCfg.Db)
	suggestionSvc := service.NewSuggestion(suggestionRepo, appCfg.Log, publisher)

	suggestionHandler := handler.NewSuggestionHandler(suggestionSvc, appCfg.Log)

	appMiddleware := middleware.NewCustomMiddleware(appCfg.Log)

	routeConfig := middleware.RouteConfig{
		App:               appCfg.App,
		SuggestionHandler: suggestionHandler,
		Middleware:        appMiddleware,
	}

	routeConfig.Setup()
}
