package backend

import "fmt"

// HTTPError is API Responseとして返すError
type HTTPError struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

// StatusCode is Http Response Status Codeを返す
func (he *HTTPError) StatusCode() int {
	return he.Code
}

// ErrorMessage is Clientに返すErrorMessageを返す
func (he *HTTPError) ErrorMessage() interface{} {
	return he
}

// Error is error interfaceを実装
func (he *HTTPError) Error() string {
	return fmt.Sprintf("status code %d: %s", he.StatusCode(), he.ErrorMessage())
}
