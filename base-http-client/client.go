package base_http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	ProjectID            string
	SecretServerKey      string
	AccessType           string
	AccessToken          string
	RefreshToken         string
	SuperSecretAdminKey  string
	PublishableClientKey string
	BaseURL              string
	HTTPClient           *http.Client
}

const (
	DefaultBaseURL        = "https://api.stack-auth.com/api/v1"
	DefaultRequestTimeout = 30 * time.Second
	DefaultAccessType     = "Server"
)

type Client struct {
	httpClient *http.Client
	config     Config
}

func NewClient(cfg Config) *Client {
	client := &Client{
		config: cfg,
	}

	if client.config.BaseURL == "" {
		client.config.BaseURL = DefaultBaseURL
	}

	if client.config.AccessType == "" {
		client.config.AccessType = DefaultAccessType
	}

	if client.config.HTTPClient != nil {
		client.httpClient = client.config.HTTPClient
	} else {
		client.httpClient = &http.Client{
			Timeout: DefaultRequestTimeout,
		}
	}

	return client
}

// SendRequest отправляет HTTP-запрос к API.
func (c *Client) SendRequest(method, path string, queryParams url.Values, body []byte) ([]byte, error) {
	fullURL := c.config.BaseURL + path

	u, err := url.Parse(fullURL)
	if err != nil {
		return nil, err
	}

	if queryParams != nil {
		params := u.Query()
		for key, values := range queryParams {
			params[key] = values
		}
		u.RawQuery = params.Encode()
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Stack-Access-Type", c.config.AccessType)
	req.Header.Set("X-Stack-Project-Id", c.config.ProjectID)
	req.Header.Set("X-Stack-Secret-Server-Key", c.config.SecretServerKey)
	req.Header.Set("X-Stack-Access-Token", c.config.AccessToken)
	req.Header.Set("X-Stack-Refresh-Token", c.config.RefreshToken)
	req.Header.Set("X-Stack-Publishable-Client-Key", c.config.PublishableClientKey)
	req.Header.Set("X-Stack-Super-Secret-Admin-Key", c.config.SuperSecretAdminKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		apiError := &APIError{}
		err := json.Unmarshal(responseBody, &apiError)
		if err == nil && apiError.Message != "" {
			return nil, apiError
		}

		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}
