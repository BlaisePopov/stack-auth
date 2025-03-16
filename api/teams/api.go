package teams

import (
	"encoding/json"
	"fmt"
	"github.com/BlaisePopov/stack-auth/base-http-client/utils"
	"net/url"

	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
)

// Client представляет клиент для работы с командами
type Client struct {
	HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с командами
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
	return &Client{HTTPClient: httpClient}
}

// ListTeams возвращает список команд проекта [https://docs.stack-auth.com/next/rest-api/server/teams/list-teams]
//
// Входные параметры:
//   - userID: идентификатор пользователя
//
// Возвращаемое значение: объект ListTeamsResponse и ошибка
func (c *Client) ListTeams(userID string) (*ListTeamsResponse, error) {
	response := &ListTeamsResponse{}
	queryParams := url.Values{}
	utils.AddOptionalStringParam(queryParams, "user_id", userID)

	rawResponse, err := c.HTTPClient.SendRequest("GET", "/teams", queryParams, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// CreateTeam создает новую команду [https://docs.stack-auth.com/next/rest-api/server/teams/create-a-team]
//
// Входные параметры:
//   - request: данные для создания команды
//
// Возвращаемое значение: объект TeamResponse и ошибка
func (c *Client) CreateTeam(request *CreateTeamRequest) (*TeamResponse, error) {
	response := &TeamResponse{}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/teams", nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// GetTeam возвращает информацию о команде по ID [https://docs.stack-auth.com/next/rest-api/server/teams/get-a-team]
//
// Входные параметры:
//   - teamID: идентификатор команды
//
// Возвращаемое значение: объект TeamResponse и ошибка
func (c *Client) GetTeam(teamID string) (*TeamResponse, error) {
	response := &TeamResponse{}
	path := fmt.Sprintf("/teams/%s", teamID)

	rawResponse, err := c.HTTPClient.SendRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// DeleteTeam удаляет команду по ID [https://docs.stack-auth.com/next/rest-api/server/teams/delete-a-team]
//
// Входные параметры:
//   - teamID: идентификатор команды
//
// Возвращаемое значение: объект SuccessResponse и ошибка
func (c *Client) DeleteTeam(teamID string) (*SuccessResponse, error) {
	response := &SuccessResponse{}
	path := fmt.Sprintf("/teams/%s", teamID)

	rawResponse, err := c.HTTPClient.SendRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// UpdateTeam обновляет информацию о команде [https://docs.stack-auth.com/next/rest-api/server/teams/update-a-team]
//
// Входные параметры:
//   - teamID: идентификатор команды
//   - request: данные для обновления
//
// Возвращаемое значение: объект TeamResponse и ошибка
func (c *Client) UpdateTeam(teamID string, request *UpdateTeamRequest) (*TeamResponse, error) {
	response := &TeamResponse{}
	path := fmt.Sprintf("/teams/%s", teamID)

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

// ListTeamMembersProfiles возвращает профили участников команды [https://docs.stack-auth.com/next/rest-api/server/teams/list-team-members-profiles]
//
// Входные параметры:
//   - teamID: идентификатор команды (опционально)
//   - userID: идентификатор пользователя (опционально)
//
// Возвращаемое значение: объект ListTeamMembersResponse и ошибка
func (c *Client) ListTeamMembersProfiles(teamID, userID string) (*ListTeamMembersResponse, error) {
	response := &ListTeamMembersResponse{}
	queryParams := url.Values{}

	utils.AddOptionalStringParam(queryParams, "team_id", teamID)
	utils.AddOptionalStringParam(queryParams, "user_id", userID)

	rawResponse, err := c.HTTPClient.SendRequest("GET", "/team-member-profiles", queryParams, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// SendInviteEmail отправляет приглашение в команду по email [https://docs.stack-auth.com/next/rest-api/server/teams/send-an-email-to-invite-a-user-to-a-team]
//
// Входные параметры:
//   - request: данные для отправки приглашения
//
// Возвращаемое значение: объект SuccessResponse и ошибка
func (c *Client) SendInviteEmail(request *SendInviteEmailRequest) (*SuccessResponse, error) {
	response := &SuccessResponse{}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/team-invitations/send-code", nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// AcceptInvite принимает приглашение в команду [https://docs.stack-auth.com/next/rest-api/server/teams/invite-a-user-to-a-team]
//
// Входные параметры:
//   - request: данные приглашения
//
// Возвращаемое значение: ошибка выполнения операции
func (c *Client) AcceptInvite(request *AcceptInviteRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	_, err = c.HTTPClient.SendRequest("POST", "/team-invitations/accept", nil, body)
	return err
}

// AddTeamMember добавляет пользователя в команду [https://docs.stack-auth.com/next/rest-api/server/teams/add-a-user-to-a-team]
//
// Входные параметры:
//   - teamID: идентификатор команды
//   - userID: идентификатор пользователя
//
// Возвращаемое значение: объект TeamMembershipResponse и ошибка
func (c *Client) AddTeamMember(teamID, userID string) (*TeamMembershipResponse, error) {
	response := &TeamMembershipResponse{}
	path := fmt.Sprintf("/team-memberships/%s/%s", teamID, userID)

	rawResponse, err := c.HTTPClient.SendRequest("POST", path, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// RemoveTeamMember удаляет пользователя из команды [https://docs.stack-auth.com/next/rest-api/server/teams/remove-a-user-from-a-team]
//
// Входные параметры:
//   - teamID: идентификатор команды
//   - userID: идентификатор пользователя
//
// Возвращаемое значение: объект SuccessResponse и ошибка
func (c *Client) RemoveTeamMember(teamID, userID string) (*SuccessResponse, error) {
	response := &SuccessResponse{}
	path := fmt.Sprintf("/team-memberships/%s/%s", teamID, userID)

	rawResponse, err := c.HTTPClient.SendRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// GetTeamMemberProfile возвращает профиль участника команды [https://docs.stack-auth.com/next/rest-api/server/teams/get-a-team-member-profile]
//
// Входные параметры:
//   - teamID: идентификатор команды
//   - userID: идентификатор пользователя
//
// Возвращаемое значение: объект TeamMemberProfileResponse и ошибка
func (c *Client) GetTeamMemberProfile(teamID, userID string) (*TeamMemberProfileResponse, error) {
	response := &TeamMemberProfileResponse{}
	path := fmt.Sprintf("/team-member-profiles/%s/%s", teamID, userID)

	rawResponse, err := c.HTTPClient.SendRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// UpdateTeamMemberProfile обновляет профиль участника команды [https://docs.stack-auth.com/next/rest-api/server/teams/update-a-team-member-profile]
//
// Входные параметры:
//   - teamID: идентификатор команды
//   - userID: идентификатор пользователя
//   - request: данные для обновления
//
// Возвращаемое значение: объект TeamMemberProfileResponse и ошибка
func (c *Client) UpdateTeamMemberProfile(teamID, userID string, request *UpdateTeamMemberProfileRequest) (*TeamMemberProfileResponse, error) {
	response := &TeamMemberProfileResponse{}
	path := fmt.Sprintf("/team-member-profiles/%s/%s", teamID, userID)

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

// GetInvitationDetails возвращает информацию о приглашении [https://docs.stack-auth.com/next/rest-api/server/teams/get-team-invitation-details]
//
// Входные параметры:
//   - code: код приглашения
//
// Возвращаемое значение: объект InvitationDetailsResponse и ошибка
func (c *Client) GetInvitationDetails(code string) (*InvitationDetailsResponse, error) {
	response := &InvitationDetailsResponse{}
	request := AcceptInviteRequest{Code: code}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/team-invitations/accept/details", nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// CheckInviteCode проверяет валидность кода приглашения [https://docs.stack-auth.com/next/rest-api/server/teams/check-if-a-team-invitation-code-is-valid]
//
// Входные параметры:
//   - code: код приглашения
//
// Возвращаемое значение: объект CheckCodeResponse и ошибка
func (c *Client) CheckInviteCode(code string) (*CheckCodeResponse, error) {
	response := &CheckCodeResponse{}
	request := AcceptInviteRequest{Code: code}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/team-invitations/accept/check-code", nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}
