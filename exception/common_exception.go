package exception

import "fmt"

type CommonException struct {
	HttpStatus int
	Code       string
	Message    string
}

func (c CommonException) Error() string {
	return fmt.Sprintf("CommonException: %s", c.Message)
}

func NewCommonException(httpStatus int, code string, message string) *CommonException {
	return &CommonException{
		HttpStatus: httpStatus,
		Code:       code,
		Message:    message,
	}
}
