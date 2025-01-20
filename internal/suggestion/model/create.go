package model

type SuggestionCreate struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
