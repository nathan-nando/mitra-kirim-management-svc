package model

type TestimonialCreate struct {
	Name        string `json:"name"`
	Img         string `json:"img"`
	Description string `json:"description"`
}

type TestimonialResponse struct {
	Id       int    `json:"id"`
	Filename string `json:"filename"`
}
