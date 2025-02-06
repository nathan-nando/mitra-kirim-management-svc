package config

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mitra-kirim-be-mgmt/http/rest/handler"
	"mitra-kirim-be-mgmt/http/rest/middleware"
	configurationRepository "mitra-kirim-be-mgmt/internal/configuration/repository"
	configurationService "mitra-kirim-be-mgmt/internal/configuration/service"
	fileUploaderRepository "mitra-kirim-be-mgmt/internal/file-uploader/repository"
	fileUploaderService "mitra-kirim-be-mgmt/internal/file-uploader/service"
	locationRepository "mitra-kirim-be-mgmt/internal/location/repository"
	locationService "mitra-kirim-be-mgmt/internal/location/service"
	"mitra-kirim-be-mgmt/internal/suggestion/repository"
	"mitra-kirim-be-mgmt/internal/suggestion/service"
	userService "mitra-kirim-be-mgmt/internal/user/service"
	"os"

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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.Error(err)
	}
	//newCache := serviceCache.NewCache(appCfg.Redis)
	publisher := pubService.NewPublisher(appCfg.Publisher.Client, appCfg.Log, appCfg.Config.RedisMaxRetry)

	suggestionRepo := repository.NewSuggestion(appCfg.Db)
	fileUploaderRepo := fileUploaderRepository.NewFileUploader(homeDir, appCfg.Log)
	locationRepo := locationRepository.NewLocation(appCfg.Db)
	configRepo := configurationRepository.NewConfiguration(appCfg.Db)

	userSvc := userService.NewUser(appCfg.Db)
	suggestionSvc := service.NewSuggestion(suggestionRepo, appCfg.Log, publisher)
	locationSvc := locationService.NewLocation(locationRepo, appCfg.Log)
	fileUploaderSvc := fileUploaderService.NewFileUploader(fileUploaderRepo, appCfg.Log)
	configSvc := configurationService.NewConfiguration(configRepo, fileUploaderSvc, appCfg.Log)

	userHandler := handler.NewUserHandler(userSvc, appCfg.Log)
	//dashboardHandler := handler.NewDashboardHandler()
	locationHandler := handler.NewLocationHandler(locationSvc, appCfg.Log)
	//testimonialHandler := handler.NewTestimonialHandler()
	suggestionHandler := handler.NewSuggestionHandler(suggestionSvc, appCfg.Log)
	configurationHandler := handler.NewConfigurationHandler(configSvc, fileUploaderSvc, appCfg.Log)
	//settingsHandler := handler.NewSettingsHandler()

	appMiddleware := middleware.NewCustomMiddleware(appCfg.Log)

	routeConfig := middleware.RouteConfig{
		App:                  appCfg.App,
		SuggestionHandler:    suggestionHandler,
		LocationHandler:      locationHandler,
		UserHandler:          userHandler,
		ConfigurationHandler: configurationHandler,
		Middleware:           appMiddleware,
	}

	routeConfig.Setup()
}
