package sessions

import (
	"encoding/json"
	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestClient(baseURL string) *Client {
	baseClient := base_http_client.NewClient(base_http_client.Config{
		BaseURL: baseURL,
	})
	return NewClient(baseClient)
}

func TestCreateSession(t *testing.T) {
	expectedResponse := &CreateSessionResponse{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/sessions", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request CreateSessionRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		assert.NoError(t, err)
		assert.Equal(t, "test_user_id", request.UserID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.CreateSession(&CreateSessionRequest{UserID: "test_user_id"})

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.AccessToken, response.AccessToken)
	assert.Equal(t, expectedResponse.RefreshToken, response.RefreshToken)
}

func TestSignOut(t *testing.T) {
	expectedResponse := &SignOutResponse{Success: true}
	testRefreshToken := "test_refresh_token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/sessions/current", r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.SignOut(testRefreshToken)

	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestRefreshAccessToken(t *testing.T) {
	expectedResponse := &RefreshAccessTokenResponse{AccessToken: "new_access_token"}
	testRefreshToken := "test_refresh_token"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/sessions/current/refresh", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.RefreshAccessToken(testRefreshToken)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.AccessToken, response.AccessToken)
}
