package service

import (
	"github.com/sirupsen/logrus"
	cacheService "mitra-kirim-be-mgmt/internal/gateways/cache/service"
	"mitra-kirim-be-mgmt/internal/location/repository"
)

type Location struct {
	Repo      *repository.Location
	Logger    *logrus.Logger
	Cache     *cacheService.Cache
	CacheTime int
}

func NewLocation(repo *repository.Location, Log *logrus.Logger, cache *cacheService.Cache, cacheTime int) *Location {
	return &Location{Repo: repo, Logger: Log, Cache: cache, CacheTime: cacheTime}
}
