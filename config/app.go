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
	testimonialRepository "mitra-kirim-be-mgmt/internal/testimonial/repository"
	testimonialService "mitra-kirim-be-mgmt/internal/testimonial/service"
	userRepository "mitra-kirim-be-mgmt/internal/user/repository"
	userService "mitra-kirim-be-mgmt/internal/user/service"
	"os"

	serviceCache "mitra-kirim-be-mgmt/internal/gateways/cache/service"
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
	cacheTime := appCfg.Config.CacheExpiration

	if err != nil {
		logrus.Error(err)
	}
	cache := serviceCache.NewCache(appCfg.Cache.Client)
	publisher := pubService.NewPublisher(appCfg.Publisher.Client, appCfg.Log, appCfg.Config.RedisMaxRetry)

	suggestionRepo := repository.NewSuggestion(appCfg.Db)
	testimonialRepo := testimonialRepository.NewRepository(appCfg.Db)
	fileUploaderRepo := fileUploaderRepository.NewFileUploader(homeDir, appCfg.Log)
	locationRepo := locationRepository.NewLocation(appCfg.Db)
	configRepo := configurationRepository.NewConfiguration(appCfg.Db)
	userRepo := userRepository.NewUser(appCfg.Db)

	fileUploaderSvc := fileUploaderService.NewFileUploader(fileUploaderRepo, appCfg.Log)
	suggestionSvc := service.NewSuggestion(suggestionRepo, appCfg.Log, publisher)
	testimonialSvc := testimonialService.NewTestimonial(testimonialRepo, fileUploaderSvc, appCfg.Log, cache, cacheTime)
	locationSvc := locationService.NewLocation(locationRepo, appCfg.Log, cache, cacheTime)
	configSvc := configurationService.NewConfiguration(configRepo, fileUploaderSvc, appCfg.Log, cache, cacheTime)
	userSvc := userService.NewUser(appCfg.Db, fileUploaderSvc, userRepo, appCfg.Config.JwtSigningKey, appCfg.Config.JwtTokenExp, appCfg.Config.JwtRefreshTokenExp)

	userHandler := handler.NewUserHandler(userSvc, appCfg.Log)
	//dashboardHandler := handler.NewDashboardHandler()
	locationHandler := handler.NewLocationHandler(locationSvc, appCfg.Log)
	testimonialHandler := handler.NewTestimonialHandler(testimonialSvc, appCfg.Log)
	suggestionHandler := handler.NewSuggestionHandler(suggestionSvc, appCfg.Log)
	configurationHandler := handler.NewConfigurationHandler(configSvc, testimonialSvc, locationSvc, appCfg.Log)

	appMiddleware := middleware.NewCustomMiddleware(appCfg.Log, appCfg.Config.JwtSigningKey)

	routeConfig := middleware.RouteConfig{
		App:                  appCfg.App,
		SuggestionHandler:    suggestionHandler,
		LocationHandler:      locationHandler,
		UserHandler:          userHandler,
		TestimonialHandler:   testimonialHandler,
		ConfigurationHandler: configurationHandler,
		Middleware:           appMiddleware,
	}

	routeConfig.Setup()
}
