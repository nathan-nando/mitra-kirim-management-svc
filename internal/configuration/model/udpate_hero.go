package model

import "mime/multipart"

type UpdateHeroLogoRequest struct {
	HeroDesc string                `json:"heroDesc"`
	HeroImg  *multipart.FileHeader `json:"heroImg"`
}

type UpdateHeroKeyVal struct {
	HeroImg  string `json:"heroImg"`
	HeroDesc string `json:"heroDesc"`
}
