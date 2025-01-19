package errorw

var (
	ErrInvalidSigningMethod = Error{Code: InvalidSigningMethod, Message: "invalid signing method"}
	ErrInvalidToken         = Error{Code: InvalidToken, Message: "invalid token"}
)

var (
	ErrConfigNotFound = Error{Code: ConfigModeNotFound, Message: "config method not found"}
)

var (
	ErrInternalServer   = Error{Code: 500000, Message: "internal server errorw"}
	ErrUnauthorized     = Error{Code: 401000, Message: "unauthorized"}
	ErrBadRequest       = Error{Code: 400000, Message: "bad request"}
	ErrForbidden        = Error{Code: 403000, Message: "you don't have access to this resource"}
	ErrSessionExpired   = Error{Code: 440000, Message: "the client's session has expired"}
	ErrResourceNotFound = Error{Code: 404000, Message: "resource not found"}
)
