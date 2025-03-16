package others

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

func TestListTeamInvitations(t *testing.T) {
	expectedResponse := &ListTeamInvitationsResponse{
		Items: []TeamInvitation{
			{
				ExpiresAtMillis: 1.1,
				ID:              "id",
				RecipientEmail:  "recipient_email",
				TeamID:          "ad962777-8244-496a-b6a2-e0c6a449c79e",
			},
		},
		Pagination: Pagination{
			NextCursor: "b3d396b8-c574-4c80-97b3-50031675ceb2",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-invitations", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.ListTeamInvitations()

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Items[0].ID, response.Items[0].ID)
	assert.Equal(t, expectedResponse.Pagination.NextCursor, response.Pagination.NextCursor)
}

func TestDeleteTeamInvitation(t *testing.T) {
	testID := "test-invitation-id"
	expectedResponse := &DeleteTeamInvitationResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-invitations/"+testID, r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.DeleteTeamInvitation(testID)

	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestConfirmNeonTransferCheck(t *testing.T) {
	expectedResponse := &ConfirmNeonTransferCheckResponse{IsCodeValid: true}
	request := &ConfirmNeonTransferCheckRequest{Code: "test-code"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/integrations/neon/projects/transfer/confirm/check", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var reqBody ConfirmNeonTransferCheckRequest
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		assert.NoError(t, err)
		assert.Equal(t, request.Code, reqBody.Code)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.ConfirmNeonTransferCheck(request)

	assert.NoError(t, err)
	assert.True(t, response.IsCodeValid)
}
