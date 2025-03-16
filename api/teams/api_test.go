package teams

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

func TestListTeams(t *testing.T) {
	expectedResponse := &ListTeamsResponse{
		Items: []TeamResponse{
			{
				ID:              "ad962777-8244-496a-b6a2-e0c6a449c79e",
				DisplayName:     "My Team",
				CreatedAtMillis: 1630000000000,
			},
		},
		Pagination: Pagination{
			NextCursor: "b3d396b8-c574-4c80-97b3-50031675ceb2",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/teams", r.URL.Path)
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "123", r.URL.Query().Get("user_id"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.ListTeams("123")
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Items[0].ID, response.Items[0].ID)
}

func TestCreateTeam(t *testing.T) {
	expectedResponse := &TeamResponse{
		ID:              "ad962777-8244-496a-b6a2-e0c6a449c79e",
		DisplayName:     "My Team",
		CreatedAtMillis: 1630000000000,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/teams", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var request CreateTeamRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "My Team", request.DisplayName)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.CreateTeam(&CreateTeamRequest{DisplayName: "My Team"})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
}

func TestGetTeam(t *testing.T) {
	expectedResponse := &TeamResponse{
		ID:          "ad962777-8244-496a-b6a2-e0c6a449c79e",
		DisplayName: "My Team",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/teams/test-team", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.GetTeam("test-team")
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
}

func TestDeleteTeam(t *testing.T) {
	expectedResponse := &SuccessResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/teams/test-team", r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.DeleteTeam("test-team")
	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestUpdateTeam(t *testing.T) {
	expectedResponse := &TeamResponse{
		ID:          "ad962777-8244-496a-b6a2-e0c6a449c79e",
		DisplayName: "Updated Team",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/teams/test-team", r.URL.Path)
		assert.Equal(t, "PATCH", r.Method)

		var request UpdateTeamRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "Updated Team", *request.DisplayName)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	name := "Updated Team"
	response, err := client.UpdateTeam("test-team", &UpdateTeamRequest{DisplayName: &name})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.DisplayName, response.DisplayName)
}

func TestListTeamMembersProfiles(t *testing.T) {
	expectedResponse := &ListTeamMembersResponse{
		Items: []TeamMemberProfileResponse{
			{
				UserID: "3241a285-8329-4d69-8f3d-316e08cf140c",
				TeamID: "ad962777-8244-496a-b6a2-e0c6a449c79e",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-member-profiles", r.URL.Path)
		assert.Equal(t, "team123", r.URL.Query().Get("team_id"))
		assert.Equal(t, "user456", r.URL.Query().Get("user_id"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.ListTeamMembersProfiles("team123", "user456")
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Items[0].UserID, response.Items[0].UserID)
}

func TestSendInviteEmail(t *testing.T) {
	expectedResponse := &SuccessResponse{Success: true, ID: "invite123"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-invitations/send-code", r.URL.Path)

		var request SendInviteEmailRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "team123", request.TeamID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.SendInviteEmail(&SendInviteEmailRequest{TeamID: "team123"})
	assert.NoError(t, err)
	assert.Equal(t, "invite123", response.ID)
}

func TestAcceptInvite(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-invitations/accept", r.URL.Path)

		var request AcceptInviteRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "testcode", request.Code)

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	err := client.AcceptInvite(&AcceptInviteRequest{Code: "testcode"})
	assert.NoError(t, err)
}

func TestAddTeamMember(t *testing.T) {
	expectedResponse := &TeamMembershipResponse{
		TeamID: "team123",
		UserID: "user456",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-memberships/team123/user456", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.AddTeamMember("team123", "user456")
	assert.NoError(t, err)
	assert.Equal(t, "user456", response.UserID)
}

func TestRemoveTeamMember(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-memberships/team123/user456", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(SuccessResponse{Success: true})
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.RemoveTeamMember("team123", "user456")
	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestGetTeamMemberProfile(t *testing.T) {
	expectedResponse := &TeamMemberProfileResponse{
		UserID: "user456",
		TeamID: "team123",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-member-profiles/team123/user456", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.GetTeamMemberProfile("team123", "user456")
	assert.NoError(t, err)
	assert.Equal(t, "user456", response.UserID)
}

func TestUpdateTeamMemberProfile(t *testing.T) {
	expectedResponse := &TeamMemberProfileResponse{
		DisplayName: "New Name",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-member-profiles/team123/user456", r.URL.Path)

		var request UpdateTeamMemberProfileRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "New Name", *request.DisplayName)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	name := "New Name"
	response, err := client.UpdateTeamMemberProfile("team123", "user456", &UpdateTeamMemberProfileRequest{DisplayName: &name})
	assert.NoError(t, err)
	assert.Equal(t, "New Name", response.DisplayName)
}

func TestGetInvitationDetails(t *testing.T) {
	expectedResponse := &InvitationDetailsResponse{
		TeamID:          "team123",
		TeamDisplayName: "Best Team",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-invitations/accept/details", r.URL.Path)

		var request AcceptInviteRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "testcode", request.Code)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.GetInvitationDetails("testcode")
	assert.NoError(t, err)
	assert.Equal(t, "Best Team", response.TeamDisplayName)
}

func TestCheckInviteCode(t *testing.T) {
	expectedResponse := &CheckCodeResponse{IsCodeValid: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/team-invitations/accept/check-code", r.URL.Path)

		var request AcceptInviteRequest
		json.NewDecoder(r.Body).Decode(&request)
		assert.Equal(t, "testcode", request.Code)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.CheckInviteCode("testcode")
	assert.NoError(t, err)
	assert.True(t, response.IsCodeValid)
}
