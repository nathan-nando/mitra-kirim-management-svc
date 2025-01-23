package config

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/http/rest/handler"
	"mitra-kirim-be-mgmt/http/rest/middleware"
	"mitra-kirim-be-mgmt/internal/suggestion/repository"
	"mitra-kirim-be-mgmt/internal/suggestion/service"
	userService "mitra-kirim-be-mgmt/internal/user/service"
	//serviceCache "mitra-kirim-be-mgmt/internal/gateways/cache/service"
	pubService "mitra-kirim-be-mgmt/internal/gateways/publisher/service"
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
	publisher := pubService.NewPublisher(appCfg.Publisher.Client, appCfg.Log, appCfg.Config.RedisMaxRetry)

	suggestionRepo := repository.NewSuggestion(appCfg.Db)

	userSvc := userService.NewUser(appCfg.Db)
	suggestionSvc := service.NewSuggestion(suggestionRepo, appCfg.Log, publisher)

	userHandler := handler.NewUserHandler(userSvc, appCfg.Log)
	suggestionHandler := handler.NewSuggestionHandler(suggestionSvc, appCfg.Log)

	appMiddleware := middleware.NewCustomMiddleware(appCfg.Log)

	routeConfig := middleware.RouteConfig{
		App:               appCfg.App,
		SuggestionHandler: suggestionHandler,
		UserHandler:       userHandler,
		Middleware:        appMiddleware,
	}

	routeConfig.Setup()
}
