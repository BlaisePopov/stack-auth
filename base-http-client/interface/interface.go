package _interface

import "net/url"

type BaseHTTPClient interface {
	SendRequest(method, path string, queryParams url.Values, body []byte) ([]byte, error)
}
