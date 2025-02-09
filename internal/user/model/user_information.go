package model

type UserInformation struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Img      string `json:"img"`
	Status   int    `json:"status"`
}
