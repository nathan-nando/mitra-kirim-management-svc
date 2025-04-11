package model

import "mime/multipart"

type UpdateServicesRequest struct {
	ServicesDescription string `json:"appDescription"`
}

type UpdateServicesLogoRequest struct {
	ServicesLogo *multipart.FileHeader `json:"heroLogo"`
}

type UpdateServicesLogoFileName struct {
	ServicesLogo string `json:"heroLogo"`
}
