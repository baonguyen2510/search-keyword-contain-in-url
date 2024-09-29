package failure

import (
	"fmt"
	"net/http"
)

type AppErr struct {
	Code         ErrCode
	OriginalErr  error
	CustomErrMsg *string
}

func (e *AppErr) Error() string {
	if e.CustomErrMsg != nil && *e.CustomErrMsg != "" {
		return *e.CustomErrMsg
	}
	if val, ok := ErrMsgMap[e.Code]; ok {
		return val
	}
	return fmt.Sprintf("Lỗi hệ thống: %d\nVui lòng thử lại sau", e.Code)
}

func (e *AppErr) HTTPCode() int {
	if val, ok := ErrCodeMap[e.Code]; ok {
		return val
	}
	return http.StatusBadRequest
}

type ErrCode int

var ErrMsgMap = map[ErrCode]string{}

var ErrCodeMap = map[ErrCode]int{}
