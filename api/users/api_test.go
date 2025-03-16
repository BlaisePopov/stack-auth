package users

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

func TestListUsers(t *testing.T) {
	expectedResponse := &ListUsersResponse{
		Items: []User{{
			ID:           "3241a285-8329-4d69-8f3d-316e08cf140c",
			DisplayName:  "John Doe",
			PrimaryEmail: "johndoe@example.com",
		}},
		Pagination: Pagination{
			NextCursor: "b3d396b8-c574-4c80-97b3-50031675ceb2",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.ListUsers("", "", "", "", false, 0)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Items[0].ID, response.Items[0].ID)
	assert.Equal(t, expectedResponse.Pagination.NextCursor, response.Pagination.NextCursor)
}

func TestCreateUser(t *testing.T) {
	expectedResponse := &UserResponse{User: User{
		ID:           "3241a285-8329-4d69-8f3d-316e08cf140c",
		DisplayName:  "John Doe",
		PrimaryEmail: "johndoe@example.com",
	}}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users", r.URL.Path)
		assert.Equal(t, "POST", r.Method)

		var req CreateUserRequest
		json.NewDecoder(r.Body).Decode(&req)
		assert.Equal(t, "John Doe", req.DisplayName)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.CreateUser(&CreateUserRequest{DisplayName: "John Doe"})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.PrimaryEmail, response.PrimaryEmail)
}

func TestGetCurrentUser(t *testing.T) {
	expectedResponse := &UserResponse{User: User{
		ID:           "3241a285-8329-4d69-8f3d-316e08cf140c",
		DisplayName:  "John Doe",
		PrimaryEmail: "johndoe@example.com",
	}}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users/me", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.GetCurrentUser()
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.DisplayName, response.DisplayName)
}

func TestDeleteCurrentUser(t *testing.T) {
	expectedResponse := &SuccessResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users/me", r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.DeleteCurrentUser()
	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestUpdateCurrentUser(t *testing.T) {
	expectedResponse := &UserResponse{User: User{
		ID:          "3241a285-8329-4d69-8f3d-316e08cf140c",
		DisplayName: "Updated Name",
	}}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users/me", r.URL.Path)
		assert.Equal(t, "PATCH", r.Method)

		var req UpdateUserRequest
		json.NewDecoder(r.Body).Decode(&req)
		assert.Equal(t, "Updated Name", req.DisplayName)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.UpdateCurrentUser(&UpdateUserRequest{DisplayName: "Updated Name"})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.DisplayName, response.DisplayName)
}

func TestGetUser(t *testing.T) {
	expectedResponse := &UserResponse{User: User{
		ID:          "3241a285-8329-4d69-8f3d-316e08cf140c",
		DisplayName: "John Doe",
	}}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users/3241a285-8329-4d69-8f3d-316e08cf140c", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.GetUser("3241a285-8329-4d69-8f3d-316e08cf140c")
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.DisplayName, response.DisplayName)
}

func TestDeleteUser(t *testing.T) {
	expectedResponse := &SuccessResponse{Success: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users/test-user-id", r.URL.Path)
		assert.Equal(t, "DELETE", r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.DeleteUser("test-user-id")
	assert.NoError(t, err)
	assert.True(t, response.Success)
}

func TestUpdateUser(t *testing.T) {
	expectedResponse := &UserResponse{User: User{
		ID:          "test-user-id",
		DisplayName: "Updated User",
	}}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/users/test-user-id", r.URL.Path)
		assert.Equal(t, "PATCH", r.Method)

		var req UpdateUserRequest
		json.NewDecoder(r.Body).Decode(&req)
		assert.Equal(t, "Updated User", req.DisplayName)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	client := setupTestClient(server.URL)
	response, err := client.UpdateUser("test-user-id", &UpdateUserRequest{DisplayName: "Updated User"})
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.ID, response.ID)
	assert.Equal(t, expectedResponse.DisplayName, response.DisplayName)
}
