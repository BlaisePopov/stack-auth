package password

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

func TestUpdatePassword(t *testing.T) {
	expectedResponse := &UpdatePasswordResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/password/update", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.UpdatePassword(&UpdatePasswordRequest{
		OldPassword: "old",
		NewPassword: "new",
	})

	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestSignUpWithEmail(t *testing.T) {
	expectedResponse := &SignUpResponse{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		UserID:       "user_id",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/password/sign-up", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.SignUpWithEmail(&SignUpWithEmailRequest{
		Email:                   "test@example.com",
		Password:                "pass",
		VerificationCallbackURL: "https://example.com",
	})

	assert.NoError(t, err)
	assert.Equal(t, "access_token", response.AccessToken)
	assert.Equal(t, "user_id", response.UserID)
}

func TestSignInWithEmail(t *testing.T) {
	expectedResponse := &SignInResponse{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		UserID:       "user_id",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/password/sign-in", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.SignInWithEmail(&SignInRequest{
		Email:    "test@example.com",
		Password: "pass",
	})

	assert.NoError(t, err)
	assert.Equal(t, "access_token", response.AccessToken)
	assert.Equal(t, "user_id", response.UserID)
}

func TestSetPassword(t *testing.T) {
	expectedResponse := &SetPasswordResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/password/set", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.SetPassword(&SetPasswordRequest{Password: "newpass"})

	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestSendResetPasswordCode(t *testing.T) {
	expectedResponse := &SendResetCodeResponse{Success: "success"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/password/send-reset-code", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.SendResetPasswordCode(&SendResetCodeRequest{
		Email:       "test@example.com",
		CallbackURL: "https://example.com",
	})

	assert.NoError(t, err)
	assert.Equal(t, "success", response.Success)
}

func TestResetPasswordWithCode(t *testing.T) {
	expectedResponse := &ResetPasswordResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/password/reset", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.ResetPasswordWithCode(&ResetPasswordRequest{
		Password: "newpass",
		Code:     "123456",
	})

	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestCheckResetPasswordCode(t *testing.T) {
	expectedResponse := &CheckCodeResponse{IsCodeValid: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/password/reset/check-code", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.CheckResetPasswordCode(&CheckCodeRequest{Code: "123456"})

	assert.NoError(t, err)
	assert.True(t, response.IsCodeValid)
}
