package service

import (
	"github.com/sirupsen/logrus"
	"mitra-kirim-be-mgmt/internal/location/repository"
)

type Location struct {
	Repo *repository.Location
	Log  *logrus.Logger
}

func NewLocation(repo *repository.Location, Log *logrus.Logger) *Location {
	return &Location{Repo: repo, Log: Log}
}
