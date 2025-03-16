package otp

import (
    "encoding/json"
    "fmt"
    base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
)

// Client представляет клиент для работы с OTP аутентификацией
type Client struct {
    HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с OTP
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
    return &Client{HTTPClient: httpClient}
}

// SignInWithCode выполняет вход с использованием одноразового кода [https://docs.stack-auth.com/next/rest-api/server/otp/sign-in-with-a-code]
//
// Входные параметры:
//   - request: данные запроса содержащие код
//
// Возвращаемое значение: объект AuthResponse и ошибка, если она возникла
func (c *Client) SignInWithCode(request *SignInWithCodeRequest) (*AuthResponse, error) {
    response := &AuthResponse{}
    bodyBytes, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/otp/sign-in", nil, bodyBytes)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// SendSignInCode отправляет код для входа на email пользователя [https://docs.stack-auth.com/next/rest-api/server/otp/send-sign-in-code]
//
// Входные параметры:
//   - request: данные запроса содержащие email и callback URL
//
// Возвращаемое значение: объект SendSignInCodeResponse и ошибка, если она возникла
func (c *Client) SendSignInCode(request *SendSignInCodeRequest) (*SendSignInCodeResponse, error) {
    response := &SendSignInCodeResponse{}
    bodyBytes, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/otp/send-sign-in-code", nil, bodyBytes)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// MFASignIn выполняет MFA аутентификацию с использованием TOTP [https://docs.stack-auth.com/next/rest-api/server/otp/mfa-sign-in]
//
// Входные параметры:
//   - request: данные запроса содержащие тип аутентификации, TOTP и код
//
// Возвращаемое значение: объект AuthResponse и ошибка, если она возникла
func (c *Client) MFASignIn(request *MFASignInRequest) (*AuthResponse, error) {
    response := &AuthResponse{}
    bodyBytes, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/mfa/sign-in", nil, bodyBytes)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// CheckSignInCode проверяет валидность кода для входа [https://docs.stack-auth.com/next/rest-api/server/otp/check-sign-in-code]
//
// Входные параметры:
//   - request: данные запроса содержащие проверяемый код
//
// Возвращаемое значение: объект CheckSignInCodeResponse и ошибка, если она возникла
func (c *Client) CheckSignInCode(request *CheckSignInCodeRequest) (*CheckSignInCodeResponse, error) {
    response := &CheckSignInCodeResponse{}
    bodyBytes, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/otp/sign-in/check-code", nil, bodyBytes)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}
