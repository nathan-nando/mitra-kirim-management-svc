package model

type LocationRequest struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	IFrameLink  string `json:"iframeLink" validate:"required"`
}
