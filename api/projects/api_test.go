package projects

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

func TestGetCurrentProject(t *testing.T) {
	expectedResponse := &GetCurrentProjectResponse{
		Config: ProjectConfig{
			ClientTeamCreationEnabled: true,
			ClientUserDeletionEnabled: true,
			CredentialEnabled:         true,
			MagicLinkEnabled:          true,
			PasskeyEnabled:            true,
			SignUpEnabled:             true,
		},
		DisplayName: "MyMusic",
		ID:          "e0b52f4d-dece-408c-af49-d23061bb0f8d",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/projects/current", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.GetCurrentProject()

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.DisplayName, response.DisplayName)
	assert.Equal(t, expectedResponse.Config, response.Config)
}
