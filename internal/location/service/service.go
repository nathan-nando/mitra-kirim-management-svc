package service

import "mitra-kirim-be-mgmt/internal/location/repository"

type Location struct {
	Repo *repository.Location
}

func NewService(repo *repository.Location) *Location {
	return &Location{Repo: repo}
}
