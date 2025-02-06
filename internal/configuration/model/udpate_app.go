package model

import "mime/multipart"

type UpdateAppRequest struct {
	AppName        string `json:"appName"`
	AppDescription string `json:"appDescription"`
}

type UpdateAppLogoRequest struct {
	AppLogo *multipart.FileHeader `json:"appLogo"`
}

type UpdateAppLogoFileName struct {
	AppLogo string `json:"appLogo"`
}
