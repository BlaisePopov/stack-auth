package others

import (
    "encoding/json"
    "fmt"
    base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
    "net/url"
)

// Client представляет клиент для работы с дополнительными методами API
type Client struct {
    HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с дополнительными методами
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
    return &Client{HTTPClient: httpClient}
}

// ListTeamInvitations возвращает список приглашений в команду. [https://docs.stack-auth.com/next/rest-api/server/others/get-team-invitations]
//
// Возвращаемое значение: объект ListTeamInvitationsResponse и ошибка, если она возникла
func (c *Client) ListTeamInvitations() (*ListTeamInvitationsResponse, error) {
    response := &ListTeamInvitationsResponse{}
    rawResponse, err := c.HTTPClient.SendRequest("GET", "/team-invitations", url.Values{}, nil)
    if err != nil {
        return nil, err
    }
    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// DeleteTeamInvitation удаляет приглашение в команду по ID. [https://docs.stack-auth.com/next/rest-api/server/others/delete-team-invitations-id]
//
// Входные параметры:
//   - id: идентификатор приглашения
//
// Возвращаемое значение: объект DeleteTeamInvitationResponse и ошибка, если она возникла
func (c *Client) DeleteTeamInvitation(id string) (*DeleteTeamInvitationResponse, error) {
    response := &DeleteTeamInvitationResponse{}
    path := fmt.Sprintf("/team-invitations/%s", id)
    rawResponse, err := c.HTTPClient.SendRequest("DELETE", path, url.Values{}, nil)
    if err != nil {
        return nil, err
    }
    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}

// ConfirmNeonTransferCheck подтверждает проверку передачи проекта Neon. [https://docs.stack-auth.com/next/rest-api/server/others/post-integrations-neon-projects-transfer-confirm-check]
//
// Входные параметры:
//   - request: данные для подтверждения проверки
//
// Возвращаемое значение: объект ConfirmNeonTransferCheckResponse и ошибка, если она возникла
func (c *Client) ConfirmNeonTransferCheck(request *ConfirmNeonTransferCheckRequest) (*ConfirmNeonTransferCheckResponse, error) {
    response := &ConfirmNeonTransferCheckResponse{}
    body, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
    }
    rawResponse, err := c.HTTPClient.SendRequest("POST", "/integrations/neon/projects/transfer/confirm/check", url.Values{}, body)
    if err != nil {
        return nil, err
    }
    if err := json.Unmarshal(rawResponse, response); err != nil {
        return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
    }
    return response, nil
}
