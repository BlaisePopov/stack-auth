package otp

import (
	"encoding/json"
	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestClient(baseURL string) *Client {
	baseClient := base_http_client.NewClient(base_http_client.Config{
		BaseURL: baseURL,
	})
	return NewClient(baseClient)
}

func TestSignInWithCode(t *testing.T) {
	expectedResponse := &AuthResponse{
		AccessToken:  "eyJhmMiJB2TO...diI4QT",
		IsNewUser:    true,
		RefreshToken: "i8ns3aq2...14y",
		UserID:       "3241a285-8329-4d69-8f3d-316e08cf140c",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/otp/sign-in", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request SignInWithCodeRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "test_code", request.Code)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.SignInWithCode(&SignInWithCodeRequest{Code: "test_code"})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.AccessToken, response.AccessToken)
	assert.Equal(t, expectedResponse.IsNewUser, response.IsNewUser)
}

func TestSendSignInCode(t *testing.T) {
	expectedResponse := &SendSignInCodeResponse{
		Nonce: "test_nonce",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/otp/send-sign-in-code", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request SendSignInCodeRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "test@example.com", request.Email)
		assert.Equal(t, "https://callback.test", request.CallbackURL)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.SendSignInCode(&SendSignInCodeRequest{
		Email:       "test@example.com",
		CallbackURL: "https://callback.test",
	})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Nonce, response.Nonce)
}

func TestMFASignIn(t *testing.T) {
	expectedResponse := &AuthResponse{
		AccessToken:  "mfa_eyJhm...TdiI4QT",
		IsNewUser:    false,
		RefreshToken: "mfa_i8ns3...14y",
		UserID:       "mfa_3241a285-8329-4d69-8f3d-316e08cf140c",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/mfa/sign-in", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request MFASignInRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "test_type", request.Type)
		assert.Equal(t, "test_totp", request.TOTP)
		assert.Equal(t, "test_code", request.Code)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.MFASignIn(&MFASignInRequest{
		Type: "test_type",
		TOTP: "test_totp",
		Code: "test_code",
	})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.RefreshToken, response.RefreshToken)
	assert.Equal(t, expectedResponse.UserID, response.UserID)
}

func TestCheckSignInCode(t *testing.T) {
	expectedResponse := &CheckSignInCodeResponse{
		IsCodeValid: true,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/auth/otp/sign-in/check-code", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request CheckSignInCodeRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "test_check_code", request.Code)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.CheckSignInCode(&CheckSignInCodeRequest{Code: "test_check_code"})
	assert.NoError(t, err)
	assert.True(t, response.IsCodeValid)
}
