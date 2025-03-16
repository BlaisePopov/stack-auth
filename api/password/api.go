package password

import (
    "encoding/json"
    "fmt"
    base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
    "net/url"
)

// Client представляет клиент для работы с аутентификацией по паролю
type Client struct {
    HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с паролями
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
    return &Client{HTTPClient: httpClient}
}

// UpdatePassword обновляет пароль текущего пользователя [https://docs.stack-auth.com/next/rest-api/server/password/update-password]
//
// Входные параметры:
//   - request: данные для обновления пароля
//
// Возвращаемое значение: объект UpdatePasswordResponse и ошибка, если она возникла
func (c *Client) UpdatePassword(request *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
    response := &UpdatePasswordResponse{}
    body, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/password/update", url.Values{}, body)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// SignUpWithEmail создает новую учетную запись с email и паролем [https://docs.stack-auth.com/next/rest-api/server/password/sign-up-with-email-and-password]
//
// Входные параметры:
//   - request: данные для регистрации
//
// Возвращаемое значение: объект SignUpResponse и ошибка, если она возникла
func (c *Client) SignUpWithEmail(request *SignUpWithEmailRequest) (*SignUpResponse, error) {
    response := &SignUpResponse{}
    body, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/password/sign-up", url.Values{}, body)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// SignInWithEmail выполняет вход в учетную запись с email и паролем [https://docs.stack-auth.com/next/rest-api/server/password/sign-in-with-email-and-password]
//
// Входные параметры:
//   - request: данные для входа
//
// Возвращаемое значение: объект SignInResponse и ошибка, если она возникла
func (c *Client) SignInWithEmail(request *SignInRequest) (*SignInResponse, error) {
    response := &SignInResponse{}
    body, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/password/sign-in", url.Values{}, body)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// SetPassword устанавливает новый пароль для текущего пользователя [https://docs.stack-auth.com/next/rest-api/server/password/set-password]
//
// Входные параметры:
//   - request: данные для установки пароля
//
// Возвращаемое значение: объект SetPasswordResponse и ошибка, если она возникла
func (c *Client) SetPassword(request *SetPasswordRequest) (*SetPasswordResponse, error) {
    response := &SetPasswordResponse{}
    body, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/password/set", url.Values{}, body)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// SendResetPasswordCode отправляет код сброса пароля на email [https://docs.stack-auth.com/next/rest-api/server/password/send-reset-password-code]
//
// Входные параметры:
//   - request: данные для отправки кода
//
// Возвращаемое значение: объект SendResetCodeResponse и ошибка, если она возникла
func (c *Client) SendResetPasswordCode(request *SendResetCodeRequest) (*SendResetCodeResponse, error) {
    response := &SendResetCodeResponse{}
    body, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/password/send-reset-code", url.Values{}, body)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// ResetPasswordWithCode сбрасывает пароль с использованием кода [https://docs.stack-auth.com/next/rest-api/server/password/reset-password-with-a-code]
//
// Входные параметры:
//   - request: данные для сброса пароля
//
// Возвращаемое значение: объект ResetPasswordResponse и ошибка, если она возникла
func (c *Client) ResetPasswordWithCode(request *ResetPasswordRequest) (*ResetPasswordResponse, error) {
    response := &ResetPasswordResponse{}
    body, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/password/reset", url.Values{}, body)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// CheckResetPasswordCode проверяет валидность кода сброса пароля [https://docs.stack-auth.com/next/rest-api/server/password/check-reset-password-code]
//
// Входные параметры:
//   - request: данные для проверки кода
//
// Возвращаемое значение: объект CheckCodeResponse и ошибка, если она возникла
func (c *Client) CheckResetPasswordCode(request *CheckCodeRequest) (*CheckCodeResponse, error) {
    response := &CheckCodeResponse{}
    body, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }

    rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/password/reset/check-code", url.Values{}, body)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}
