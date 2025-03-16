package sessions

import (
	"encoding/json"
	"fmt"
	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
	"net/url"
)

// Client представляет клиент для работы с сессиями пользователя
type Client struct {
	HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с сессиями
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
	return &Client{HTTPClient: httpClient}
}

// CreateSession создает новую сессию для пользователя [https://docs.stack-auth.com/next/rest-api/server/sessions/create-session]
//
// Входные параметры:
//   - request: данные для создания сессии
//
// Возвращаемое значение: объект CreateSessionResponse и ошибка, если она возникла
func (c *Client) CreateSession(request *CreateSessionRequest) (*CreateSessionResponse, error) {
	response := &CreateSessionResponse{}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/sessions", url.Values{}, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// SignOut завершает текущую сессию пользователя [https://docs.stack-auth.com/next/rest-api/server/sessions/sign-out-of-the-current-session]
//
// Входные параметры:
//   - refreshToken: refresh token из заголовка запроса
//
// Возвращаемое значение: объект SignOutResponse и ошибка, если она возникла
func (c *Client) SignOut(refreshToken string) (*SignOutResponse, error) {
	response := &SignOutResponse{}

	rawResponse, err := c.HTTPClient.SendRequest("DELETE", "/auth/sessions/current", url.Values{}, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// RefreshAccessToken обновляет access token с использованием refresh token [https://docs.stack-auth.com/next/rest-api/server/sessions/refresh-access-token]
//
// Входные параметры:
//   - refreshToken: refresh token из заголовка запроса
//
// Возвращаемое значение: объект RefreshAccessTokenResponse и ошибка, если она возникла
func (c *Client) RefreshAccessToken(refreshToken string) (*RefreshAccessTokenResponse, error) {
	response := &RefreshAccessTokenResponse{}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/sessions/current/refresh", url.Values{}, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}
