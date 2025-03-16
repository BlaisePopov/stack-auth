package permissions

import (
	"encoding/json"
	"fmt"
	"github.com/BlaisePopov/stack-auth/base-http-client/utils"
	"net/url"

	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
)

// Client представляет клиент для работы с разрешениями
type Client struct {
	HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с разрешениями
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
	return &Client{HTTPClient: httpClient}
}

// ListTeamPermissions возвращает список командных разрешений пользователя [https://docs.stack-auth.com/next/rest-api/server/permissions/list-team-permissions-of-a-user]
//
// Параметры:
//   - teamID: идентификатор команды (опционально)
//   - userID: идентификатор пользователя (опционально)
//   - permissionID: идентификатор разрешения (опционально)
//   - recursive: флаг рекурсивного поиска (опционально)
func (c *Client) ListTeamPermissions(teamID, userID, permissionID, recursive string) (*ListTeamPermissionsResponse, error) {
	response := &ListTeamPermissionsResponse{}
	queryParams := url.Values{}

	utils.AddOptionalStringParam(queryParams, "team_id", teamID)
	utils.AddOptionalStringParam(queryParams, "user_id", userID)
	utils.AddOptionalStringParam(queryParams, "permission_id", permissionID)
	utils.AddOptionalStringParam(queryParams, "recursive", recursive)

	rawResponse, err := c.HTTPClient.SendRequest("GET", "/team-permissions", queryParams, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// GrantTeamPermissionToUser выдает пользователю командное разрешение [https://docs.stack-auth.com/next/rest-api/server/permissions/grant-a-team-permission-to-a-user]
//
// Параметры:
//   - teamID: идентификатор команды
//   - userID: идентификатор пользователя
//   - permissionID: идентификатор разрешения
//   - recursive: флаг рекурсивного применения (опционально)
func (c *Client) GrantTeamPermissionToUser(teamID, userID, permissionID, recursive string) (*GrantTeamPermissionResponse, error) {
	response := &GrantTeamPermissionResponse{}
	path := fmt.Sprintf("/team-permissions/%s/%s/%s", teamID, userID, permissionID)
	queryParams := url.Values{}

	utils.AddOptionalStringParam(queryParams, "team_id", teamID)
	utils.AddOptionalStringParam(queryParams, "user_id", userID)
	utils.AddOptionalStringParam(queryParams, "permission_id", permissionID)
	utils.AddOptionalStringParam(queryParams, "recursive", recursive)

	rawResponse, err := c.HTTPClient.SendRequest("POST", path, queryParams, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// RevokeTeamPermissionFromUser отзывает командное разрешение у пользователя [https://docs.stack-auth.com/next/rest-api/server/permissions/revoke-a-team-permission-from-a-user]
//
// Параметры:
//   - teamID: идентификатор команды
//   - userID: идентификатор пользователя
//   - permissionID: идентификатор разрешения
//   - recursive: флаг рекурсивного отзыва (опционально)
func (c *Client) RevokeTeamPermissionFromUser(teamID, userID, permissionID, recursive string) (*RevokeTeamPermissionResponse, error) {
	response := &RevokeTeamPermissionResponse{}
	path := fmt.Sprintf("/team-permissions/%s/%s/%s", teamID, userID, permissionID)
	queryParams := url.Values{}

	utils.AddOptionalStringParam(queryParams, "team_id", teamID)
	utils.AddOptionalStringParam(queryParams, "user_id", userID)
	utils.AddOptionalStringParam(queryParams, "permission_id", permissionID)
	utils.AddOptionalStringParam(queryParams, "recursive", recursive)

	rawResponse, err := c.HTTPClient.SendRequest("DELETE", path, queryParams, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}
