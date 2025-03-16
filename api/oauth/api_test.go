package oauth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client"
	"github.com/stretchr/testify/assert"
)

func setupTestClient(baseURL string) *Client {
	baseClient := base_http_client.NewClient(base_http_client.Config{
		BaseURL: baseURL,
	})
	return NewClient(baseClient)
}

func TestClient_Token(t *testing.T) {
	expectedResponse := &TokenResponse{
		Key:   "access_token",
		Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/oauth/token", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request TokenRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		assert.NoError(t, err)
		assert.Equal(t, "authorization_code", request.GrantType)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.Token(&TokenRequest{GrantType: "authorization_code"})

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Key, response.Key)
	assert.Equal(t, expectedResponse.Value, response.Value)
}

func TestClient_Authorize(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/auth/oauth/authorize/test-provider"
		assert.Equal(t, expectedPath, r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		query := r.URL.Query()
		assert.Equal(t, "test-client", query.Get("client_id"))
		assert.Equal(t, "secret", query.Get("client_secret"))
		assert.Equal(t, "http://callback", query.Get("redirect_uri"))
		assert.Equal(t, "link", query.Get("type"))
		assert.Equal(t, "email", query.Get("provider_scope"))

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	query := &AuthorizeQuery{
		ClientID:            "test-client",
		ClientSecret:        "secret",
		RedirectURI:         "http://callback",
		Scope:               "profile",
		State:               "state123",
		GrantType:           "authorization_code",
		CodeChallenge:       "challenge",
		CodeChallengeMethod: "S256",
		ResponseType:        "code",
		Type:                "link",
		ProviderScope:       "email",
	}

	err := client.Authorize("test-provider", query)
	assert.NoError(t, err)
}
