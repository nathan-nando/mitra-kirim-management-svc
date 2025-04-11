package model

type Dashboard struct {
	ViewerCount      int64 `json:"viewerCount"`
	SuggestionCount  int64 `json:"suggestionCount"`
	TestimonialCount int64 `json:"testimonialCount"`
	LocationCount    int64 `json:"locationCount"`
}
