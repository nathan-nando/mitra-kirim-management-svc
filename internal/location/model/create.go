package model

type LocationRequest struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"nama" validate:"required"`
	Address     string `json:"alamat" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Whatsapp    string `json:"whatsapp" validate:"required"`
	Description string `json:"deskripsi" validate:"required"`
	IFrameLink  string `json:"iframeLink" validate:"required"`
}
