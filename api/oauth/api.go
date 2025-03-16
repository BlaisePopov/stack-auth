package oauth

import (
	"encoding/json"
	"fmt"
	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
	"github.com/BlaisePopov/stack-auth/base-http-client/utils"
	"net/url"
)

// Client представляет клиент для работы с OAuth аутентификацией
type Client struct {
	HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с OAuth
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
	return &Client{HTTPClient: httpClient}
}

// Token обменивает код авторизации или refresh token на access token [https://docs.stack-auth.com/next/rest-api/server/oauth/o-auth-token-endpoints]
//
// Входные параметры:
//   - request: данные запроса для получения токена
//
// Возвращаемое значение: объект TokenResponse и ошибка, если она возникла
func (c *Client) Token(request *TokenRequest) (*TokenResponse, error) {
	response := &TokenResponse{}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка сериализации запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/auth/oauth/token", nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// Authorize инициирует OAuth авторизацию или связывание аккаунта [https://docs.stack-auth.com/next/rest-api/server/oauth/o-auth-authorize-endpoint]
//
// Входные параметры:
//   - providerID: идентификатор OAuth-провайдера
//   - request: параметры запроса для авторизации
//
// Возвращаемое значение: ошибка, если она возникла
func (c *Client) Authorize(providerID string, query *AuthorizeQuery) error {
	path := fmt.Sprintf("/auth/oauth/authorize/%s", url.PathEscape(providerID))

	queryParams := url.Values{}
	utils.AddOptionalStringParam(queryParams, "type", query.Type)
	utils.AddOptionalStringParam(queryParams, "token", query.Token)
	utils.AddOptionalStringParam(queryParams, "provider_scope", query.ProviderScope)
	utils.AddOptionalStringParam(queryParams, "error_redirect_uri", query.ErrorRedirectURI)
	utils.AddOptionalStringParam(queryParams, "after_callback_redirect_url", query.AfterCallbackRedirectURL)

	queryParams.Add("client_id", query.ClientID)
	queryParams.Add("client_secret", query.ClientSecret)
	queryParams.Add("redirect_uri", query.RedirectURI)
	queryParams.Add("scope", query.Scope)
	queryParams.Add("state", query.State)
	queryParams.Add("grant_type", query.GrantType)
	queryParams.Add("code_challenge", query.CodeChallenge)
	queryParams.Add("code_challenge_method", query.CodeChallengeMethod)
	queryParams.Add("response_type", query.ResponseType)

	_, err := c.HTTPClient.SendRequest("GET", path, queryParams, nil)
	return err
}
