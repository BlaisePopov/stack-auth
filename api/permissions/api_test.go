package permissions

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

func TestListTeamPermissions(t *testing.T) {
	expectedResponse := &ListTeamPermissionsResponse{
		Items: []TeamPermission{
			{
				ID:     "read_secret_info",
				TeamID: "ad962777-8244-496a-b6a2-e0c6a449c79e",
				UserID: "3241a285-8329-4d69-8f3d-316e08cf140c",
			},
		},
		Pagination: Pagination{
			NextCursor: "b3d396b8-c574-4c80-97b3-50031675ceb2",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-permissions", r.URL.Path)
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "test_team", r.URL.Query().Get("team_id"))
		assert.Equal(t, "test_user", r.URL.Query().Get("user_id"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.ListTeamPermissions("test_team", "test_user", "", "")

	assert.NoError(t, err)
	assert.Len(t, response.Items, 1)
	assert.Equal(t, "read_secret_info", response.Items[0].ID)
	assert.Equal(t, "b3d396b8-c574-4c80-97b3-50031675ceb2", response.Pagination.NextCursor)
}

func TestGrantTeamPermissionToUser(t *testing.T) {
	expectedResponse := &GrantTeamPermissionResponse{
		ID:     "read_secret_info",
		TeamID: "ad962777-8244-496a-b6a2-e0c6a449c79e",
		UserID: "3241a285-8329-4d69-8f3d-316e08cf140c",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-permissions/team_123/user_456/perm_789", r.URL.Path)
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "true", r.URL.Query().Get("recursive"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.GrantTeamPermissionToUser("team_123", "user_456", "perm_789", "true")

	assert.NoError(t, err)
	assert.Equal(t, "read_secret_info", response.ID)
	assert.Equal(t, "ad962777-8244-496a-b6a2-e0c6a449c79e", response.TeamID)
}

func TestRevokeTeamPermissionFromUser(t *testing.T) {
	expectedResponse := &RevokeTeamPermissionResponse{
		Success: true,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-permissions/team_123/user_456/perm_789", r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)
		assert.Equal(t, "false", r.URL.Query().Get("recursive"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.RevokeTeamPermissionFromUser("team_123", "user_456", "perm_789", "false")

	assert.NoError(t, err)
	assert.True(t, response.Success)
}
