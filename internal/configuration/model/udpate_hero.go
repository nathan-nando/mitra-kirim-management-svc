package model

import "mime/multipart"

type UpdateHeroLogoRequest struct {
	HeroLogo *multipart.FileHeader `json:"heroLogo"`
}

type UpdateHeroLogoFileName struct {
	HeroLogo string `json:"heroLogo"`
}
