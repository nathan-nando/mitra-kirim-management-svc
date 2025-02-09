package model

type UserUpdate struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}
