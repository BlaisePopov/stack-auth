package projects

import (
    "encoding/json"
    "fmt"
    base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
    "net/url"
)

// Client представляет клиент для работы с проектами
type Client struct {
    HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с проектами
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
    return &Client{HTTPClient: httpClient}
}

// GetCurrentProject возвращает информацию о текущем проекте. [https://docs.stack-auth.com/next/rest-api/server/projects/get-the-current-project]
//
// Возвращаемое значение:
//   - объект GetCurrentProjectResponse с данными проекта
//   - ошибка, если возникла при выполнении запроса или декодировании ответа
func (c *Client) GetCurrentProject() (*GetCurrentProjectResponse, error) {
    response := &GetCurrentProjectResponse{}

    rawResponse, err := c.HTTPClient.SendRequest("GET", "/projects/current", url.Values{}, nil)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    
    return response, nil
}
