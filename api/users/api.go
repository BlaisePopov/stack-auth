package users

import (
	"encoding/json"
	"fmt"
	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
	"github.com/BlaisePopov/stack-auth/base-http-client/utils"
	"net/url"
	"strconv"
)

// Client представляет клиент для работы с пользователями
type Client struct {
	HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с пользователями
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
	return &Client{HTTPClient: httpClient}
}

// ListUsers возвращает список пользователей проекта. [https://docs.stack-auth.com/next/rest-api/server/users/list-users]
//
// Входные параметры:
//   - teamID: идентификатор команды (опционально)
//   - cursor: курсор для пагинации (опционально)
//   - orderBy: поле для сортировки (опционально)
//   - query: поисковый запрос (опционально)
//   - desc: обратный порядок сортировки (опционально)
//   - limit: ограничение количества результатов (опционально)
//
// Возвращаемое значение: объект ListUsersResponse и ошибка, если она возникла
func (c *Client) ListUsers(teamID, cursor, orderBy, query string, desc bool, limit int) (*ListUsersResponse, error) {
	response := &ListUsersResponse{}
	queryParams := url.Values{}

	utils.AddOptionalStringParam(queryParams, "team_id", teamID)
	utils.AddOptionalIntParam(queryParams, "limit", limit)
	utils.AddOptionalStringParam(queryParams, "cursor", cursor)
	utils.AddOptionalStringParam(queryParams, "order_by", orderBy)
	utils.AddOptionalStringParam(queryParams, "query", query)
	queryParams.Add("desc", strconv.FormatBool(desc))

	rawResponse, err := c.HTTPClient.SendRequest("GET", "/users", queryParams, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// CreateUser создает нового пользователя. [https://docs.stack-auth.com/next/rest-api/server/users/create-user]
//
// Входные параметры:
//   - request: данные для создания пользователя
//
// Возвращаемое значение: объект UserResponse и ошибка, если она возникла
func (c *Client) CreateUser(request *CreateUserRequest) (*UserResponse, error) {
	response := &UserResponse{}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/users", nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// GetCurrentUser возвращает данные текущего аутентифицированного пользователя. [https://docs.stack-auth.com/next/rest-api/server/users/get-current-user]
//
// Возвращаемое значение: объект UserResponse и ошибка, если она возникла
func (c *Client) GetCurrentUser() (*UserResponse, error) {
	response := &UserResponse{}
	rawResponse, err := c.HTTPClient.SendRequest("GET", "/users/me", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// DeleteCurrentUser удаляет текущего аутентифицированного пользователя. [https://docs.stack-auth.com/next/rest-api/server/users/delete-current-user]
//
// Возвращаемое значение: объект SuccessResponse и ошибка, если она возникла
func (c *Client) DeleteCurrentUser() (*SuccessResponse, error) {
	response := &SuccessResponse{}
	rawResponse, err := c.HTTPClient.SendRequest("DELETE", "/users/me", nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// UpdateCurrentUser обновляет данные текущего пользователя. [https://docs.stack-auth.com/next/rest-api/server/users/update-current-user]
//
// Входные параметры:
//   - request: данные для обновления
//
// Возвращаемое значение: объект UserResponse и ошибка, если она возникла
func (c *Client) UpdateCurrentUser(request *UpdateUserRequest) (*UserResponse, error) {
	return c.updateUser("me", request)
}

// GetUser возвращает пользователя по ID. [https://docs.stack-auth.com/next/rest-api/server/users/get-user]
//
// Входные параметры:
//   - userID: идентификатор пользователя
//
// Возвращаемое значение: объект UserResponse и ошибка, если она возникла
func (c *Client) GetUser(userID string) (*UserResponse, error) {
	response := &UserResponse{}
	path := fmt.Sprintf("/users/%s", userID)

	rawResponse, err := c.HTTPClient.SendRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// DeleteUser удаляет пользователя по ID. [https://docs.stack-auth.com/next/rest-api/server/users/delete-user]
//
// Входные параметры:
//   - userID: идентификатор пользователя
//
// Возвращаемое значение: объект SuccessResponse и ошибка, если она возникла
func (c *Client) DeleteUser(userID string) (*SuccessResponse, error) {
	response := &SuccessResponse{}
	path := fmt.Sprintf("/users/%s", userID)

	rawResponse, err := c.HTTPClient.SendRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// UpdateUser обновляет данные пользователя по ID. [https://docs.stack-auth.com/next/rest-api/server/users/update-user]
//
// Входные параметры:
//   - userID: идентификатор пользователя
//   - request: данные для обновления
//
// Возвращаемое значение: объект UserResponse и ошибка, если она возникла
func (c *Client) UpdateUser(userID string, request *UpdateUserRequest) (*UserResponse, error) {
	return c.updateUser(userID, request)
}

func (c *Client) updateUser(userID string, request *UpdateUserRequest) (*UserResponse, error) {
	path := fmt.Sprintf("/users/%s", userID)
	response := &UserResponse{}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("PATCH", path, nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}
