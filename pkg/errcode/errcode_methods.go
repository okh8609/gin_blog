package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	// 錯誤碼
	ErrCode int `json:"err_code"`
	// 錯誤消息
	ErrMsg string `json:"err_msg"`
	// 詳細信息
	ErrDetails []string `json:"err_details"`
}

var codes = map[int]string{}

func NewError(c int, m string) *Error {
	if _, ok := codes[c]; ok {
		panic(fmt.Sprintf("錯誤碼 %d 已經存在，請更換一個", c))
	}
	codes[c] = m
	return &Error{ErrCode: c, ErrMsg: m}
}

func (e *Error) GetError() string {
	return fmt.Sprintf("錯誤碼：%d, 錯誤信息:：%s", e.GetCode(), e.GetMsg())
}

func (e *Error) GetCode() int {
	return e.ErrCode
}

func (e *Error) GetMsg() string {
	return e.ErrMsg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.ErrMsg, args...)
}

func (e *Error) GetDetails() []string {
	return e.ErrDetails
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.ErrDetails = []string{}
	newError.ErrDetails = append(newError.ErrDetails, details...)
	return &newError
}

func (e *Error) GetHttpStatusCode() int {
	switch e.GetCode() {
	case Success.GetCode():
		return http.StatusOK
	case ServerError.GetCode():
		return http.StatusInternalServerError
	case InvalidParams.GetCode():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.GetCode():
		fallthrough
	case UnauthorizedTokenError.GetCode():
		fallthrough
	case UnauthorizedTokenGenerate.GetCode():
		fallthrough
	case UnauthorizedTokenTimeout.GetCode():
		return http.StatusUnauthorized
	case TooManyRequests.GetCode():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
