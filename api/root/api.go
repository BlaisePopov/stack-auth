package root

import (
    "encoding/json"
    "fmt"
    base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
)

// Client представляет клиент для работы с корневым эндпоинтом API
type Client struct {
    HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создаёт новый экземпляр клиента для корневого эндпоинта
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
    return &Client{HTTPClient: httpClient}
}

// GetAPIInfo возвращает информацию о API. [https://docs.stack-auth.com/next/rest-api/server//api-v-1]
//
// Возвращаемое значение: объект GetAPIInfoResponse и ошибка, если она возникла
func (c *Client) GetAPIInfo() (*GetAPIInfoResponse, error) {
    response := &GetAPIInfoResponse{}

    rawResponse, err := c.HTTPClient.SendRequest("GET", "", nil, nil)
    if err != nil {
        return nil, err
    }

    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}
