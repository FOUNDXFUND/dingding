package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	HTTPStatus int
	Code       int
	Message    string
}

func (msg *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"code":    msg.Code,
		"message": msg.Message,
	})
}

func (msg Error) Error() string {
	return fmt.Sprintf("%v:%v/%v", msg.HTTPStatus, msg.Code, msg.Message)
}

func newError(status, code int, message string) *Error {
	return &Error{
		HTTPStatus: status,
		Code:       code,
		Message:    message,
	}
}

var (
	ErrWrongParam = newError(http.StatusBadRequest, -1000, "Parameter Error")
	ErrRateLimit = newError(http.StatusBadRequest, -1001, "Rate Limit!")
	ErrDataNotFound= newError(http.StatusNotFound, -4004, "Data Not Found!")
)
