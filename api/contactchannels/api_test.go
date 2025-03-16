package contactchannels

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

func TestListContactChannels(t *testing.T) {
	expectedResponse := &ListContactChannelsResponse{
		Items: []ContactChannel{
			{
				ID:          "b3d396b8-c574-4c80-97b3-50031675ceb2",
				IsPrimary:   true,
				IsVerified:  true,
				Type:        "email",
				UsedForAuth: true,
				UserID:      "3241a285-8329-4d69-8f3d-316e08cf140c",
				Value:       "johndoe@example.com",
			},
		},
		Pagination: struct {
			NextCursor string `json:"next_cursor"`
		}{
			NextCursor: "b3d396b8-c574-4c80-97b3-50031675ceb2",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/contact-channels", r.URL.Path)
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "3241a285", r.URL.Query().Get("user_id"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.ListContactChannels("3241a285", "")
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Items[0].ID, response.Items[0].ID)
	assert.Equal(t, expectedResponse.Pagination.NextCursor, response.Pagination.NextCursor)
}

func TestCreateContactChannel(t *testing.T) {
	expectedResponse := &ContactChannelResponse{
		ID:          "b3d396b8-c574-4c80-97b3-50031675ceb2",
		UserID:      "3241a285-8329-4d69-8f3d-316e08cf140c",
		Value:       "johndoe@example.com",
		Type:        "email",
		UsedForAuth: true,
		IsVerified:  true,
		IsPrimary:   true,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/contact-channels", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request CreateContactChannelRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "3241a285", request.UserID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.CreateContactChannel(&CreateContactChannelRequest{
		UserID:      "3241a285",
		Value:       "johndoe@example.com",
		Type:        "email",
		UsedForAuth: true,
	})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
}

func TestVerifyEmail(t *testing.T) {
	expectedResponse := &VerifyResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/contact-channels/verify", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request VerifyRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "123456", request.Code)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.VerifyEmail(&VerifyRequest{Code: "123456"})
	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestCheckEmailVerificationCode(t *testing.T) {
	expectedResponse := &CheckCodeResponse{IsCodeValid: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/contact-channels/verify/check-code", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request CheckCodeRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "654321", request.Code)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.CheckEmailVerificationCode(&CheckCodeRequest{Code: "654321"})
	assert.NoError(t, err)
	assert.True(t, response.IsCodeValid)
}

func TestGetContactChannel(t *testing.T) {
	expectedResponse := &ContactChannelResponse{
		ID:          "test-channel-id",
		UserID:      "test-user-id",
		Value:       "test@example.com",
		Type:        "email",
		UsedForAuth: true,
		IsVerified:  true,
		IsPrimary:   true,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/contact-channels/test-user-id/test-channel-id", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.GetContactChannel("test-user-id", "test-channel-id")
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
}

func TestDeleteContactChannel(t *testing.T) {
	expectedResponse := &DeleteResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/contact-channels/test-user-id/test-channel-id", r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.DeleteContactChannel("test-user-id", "test-channel-id")
	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestUpdateContactChannel(t *testing.T) {
	expectedResponse := &ContactChannelResponse{
		ID:          "test-channel-id",
		UserID:      "test-user-id",
		Value:       "updated@example.com",
		Type:        "email",
		UsedForAuth: true,
		IsVerified:  true,
		IsPrimary:   true,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/contact-channels/test-user-id/test-channel-id", r.URL.Path)
		assert.Equal(t, "PATCH", r.Method)

		var request UpdateContactChannelRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "updated@example.com", *request.Value)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	newValue := "updated@example.com"
	response, err := client.UpdateContactChannel("test-user-id", "test-channel-id", &UpdateContactChannelRequest{
		Value: &newValue,
	})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Value, response.Value)
}

func TestSendVerificationCode(t *testing.T) {
	expectedResponse := &SendCodeResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/contact-channels/test-user-id/test-channel-id/send-verification-code", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request SendCodeRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "https://callback.example", request.CallbackURL)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.SendVerificationCode("test-user-id", "test-channel-id", &SendCodeRequest{
		CallbackURL: "https://callback.example",
	})
	assert.NoError(t, err)
	assert.True(t, response.Success)
}
