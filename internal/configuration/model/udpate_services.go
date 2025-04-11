package model

type ServicesLayout struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

type UpdateServicesRequest struct {
	Services string `json:"services"`
}
