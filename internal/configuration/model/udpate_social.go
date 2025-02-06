package model

type UpdateSocialRequest struct {
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
	Whatsapp  string `json:"whatsapp"`
}
