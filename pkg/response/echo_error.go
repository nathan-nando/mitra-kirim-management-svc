package response

type ErrorResponse struct {
	RequestID string `json:"requestId"`
	Success   bool   `json:"success" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode,omitempty"`
	Internal  error  `json:"-"`
}
