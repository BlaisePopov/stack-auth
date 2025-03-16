package contactchannels

import (
	"encoding/json"
	"fmt"
	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client/interface"
	"github.com/BlaisePopov/stack-auth/base-http-client/utils"
	"net/url"
)

// Client представляет клиент для работы с контактными каналами
type Client struct {
	HTTPClient base_http_client.BaseHTTPClient
}

// NewClient создает новый экземпляр клиента для работы с контактными каналами
func NewClient(httpClient base_http_client.BaseHTTPClient) *Client {
	return &Client{HTTPClient: httpClient}
}

// ListContactChannels возвращает список контактных каналов пользователя. [https://docs.stack-auth.com/next/rest-api/server/contact-channels/list-contact-channels]
//
// Входные параметры:
//   - userID: идентификатор пользователя
//   - contactChannelID: идентификатор контактного канала (опционально)
//
// Возвращаемое значение: объект ListContactChannelsResponse и ошибка, если она возникла
func (c *Client) ListContactChannels(userID, contactChannelID string) (*ListContactChannelsResponse, error) {
	response := &ListContactChannelsResponse{}
	queryParams := url.Values{}
	utils.AddOptionalStringParam(queryParams, "user_id", userID)
	utils.AddOptionalStringParam(queryParams, "contact_channel_id", contactChannelID)

	rawResponse, err := c.HTTPClient.SendRequest("GET", "/contact-channels", queryParams, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// CreateContactChannel создает новый контактный канал для пользователя. [https://docs.stack-auth.com/next/rest-api/server/contact-channels/create-a-contact-channel]
//
// Входные параметры:
//   - request: данные для создания контактного канала
//
// Возвращаемое значение: объект ContactChannelResponse и ошибка, если она возникла
func (c *Client) CreateContactChannel(request *CreateContactChannelRequest) (*ContactChannelResponse, error) {
	response := &ContactChannelResponse{}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/contact-channels", nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// VerifyEmail подтверждает email пользователя с помощью кода верификации. [https://docs.stack-auth.com/next/rest-api/server/contact-channels/verify-an-email]
//
// Входные параметры:
//   - request: данные с кодом подтверждения
//
// Возвращаемое значение: объект VerifyResponse и ошибка, если она возникла
func (c *Client) VerifyEmail(request *VerifyRequest) (*VerifyResponse, error) {
	response := &VerifyResponse{}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/contact-channels/verify", nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// CheckEmailVerificationCode проверяет валидность кода подтверждения email. [https://docs.stack-auth.com/next/rest-api/server/contact-channels/check-email-verification-code]
//
// Входные параметры:
//   - request: данные с кодом подтверждения
//
// Возвращаемое значение: объект CheckCodeResponse и ошибка, если она возникла
func (c *Client) CheckEmailVerificationCode(request *CheckCodeRequest) (*CheckCodeResponse, error) {
	response := &CheckCodeResponse{}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", "/contact-channels/verify/check-code", nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// GetContactChannel возвращает контактный канал по идентификаторам пользователя и канала. [https://docs.stack-auth.com/next/rest-api/server/contact-channels/get-a-contact-channel]
//
// Входные параметры:
//   - userID: идентификатор пользователя
//   - contactChannelID: идентификатор контактного канала
//
// Возвращаемое значение: объект ContactChannelResponse и ошибка, если она возникла
func (c *Client) GetContactChannel(userID, contactChannelID string) (*ContactChannelResponse, error) {
	response := &ContactChannelResponse{}
	path := fmt.Sprintf("/contact-channels/%s/%s", userID, contactChannelID)

	rawResponse, err := c.HTTPClient.SendRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// DeleteContactChannel удаляет контактный канал пользователя. [https://docs.stack-auth.com/next/rest-api/server/contact-channels/delete-a-contact-channel]
//
// Входные параметры:
//   - userID: идентификатор пользователя
//   - contactChannelID: идентификатор контактного канала
//
// Возвращаемое значение: объект DeleteResponse и ошибка, если она возникла
func (c *Client) DeleteContactChannel(userID, contactChannelID string) (*DeleteResponse, error) {
	response := &DeleteResponse{}
	path := fmt.Sprintf("/contact-channels/%s/%s", userID, contactChannelID)

	rawResponse, err := c.HTTPClient.SendRequest("DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}

// UpdateContactChannel обновляет существующий контактный канал. [https://docs.stack-auth.com/next/rest-api/server/contact-channels/update-a-contact-channel]
//
// Входные параметры:
//   - userID: идентификатор пользователя
//   - contactChannelID: идентификатор контактного канала
//   - request: данные для обновления
//
// Возвращаемое значение: объект ContactChannelResponse и ошибка, если она возникла
func (c *Client) UpdateContactChannel(userID, contactChannelID string, request *UpdateContactChannelRequest) (*ContactChannelResponse, error) {
	response := &ContactChannelResponse{}
	path := fmt.Sprintf("/contact-channels/%s/%s", userID, contactChannelID)

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

// SendVerificationCode отправляет код подтверждения на контактный канал. [https://docs.stack-auth.com/next/rest-api/server/contact-channels/send-contact-channel-verification-code]
//
// Входные параметры:
//   - userID: идентификатор пользователя
//   - contactChannelID: идентификатор контактного канала
//   - request: данные с callback URL
//
// Возвращаемое значение: объект SendCodeResponse и ошибка, если она возникла
func (c *Client) SendVerificationCode(userID, contactChannelID string, request *SendCodeRequest) (*SendCodeResponse, error) {
	response := &SendCodeResponse{}
	path := fmt.Sprintf("/contact-channels/%s/%s/send-verification-code", userID, contactChannelID)

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования запроса: %w", err)
	}

	rawResponse, err := c.HTTPClient.SendRequest("POST", path, nil, body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rawResponse, response); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}
	return response, nil
}
