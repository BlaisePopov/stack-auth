package base_http_client

import "fmt"

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"error"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
