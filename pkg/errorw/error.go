package errorw

import "fmt"

type ErrorCode int

type Error struct {
	Message string
	Code    ErrorCode
}

func (e Error) Error() string {
	return fmt.Sprintf("[%d] - %s \n", e.Code, e.Message)
}
