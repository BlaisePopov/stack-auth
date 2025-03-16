package users

// ListUsersResponse содержит ответ со списком пользователей
type ListUsersResponse struct {
    Items      []User       `json:"items"`
    Pagination Pagination   `json:"pagination"`
}

// UserResponse содержит данные пользователя
type UserResponse struct {
    User
}

// SuccessResponse содержит статус успешной операции
type SuccessResponse struct {
    Success bool `json:"success"`
}

// User представляет модель пользователя
type User struct {
    ID                      string                 `json:"id"`
    DisplayName             string                 `json:"display_name"`
    ProfileImageURL         string                 `json:"profile_image_url"`
    ClientMetadata          map[string]interface{} `json:"client_metadata"`
    ClientReadOnlyMetadata  map[string]interface{} `json:"client_read_only_metadata"`
    ServerMetadata          map[string]interface{} `json:"server_metadata"`
    PrimaryEmail            string                 `json:"primary_email"`
    PrimaryEmailVerified    bool                   `json:"primary_email_verified"`
    PrimaryEmailAuthEnabled bool                   `json:"primary_email_auth_enabled"`
    LastActiveAtMillis      int64                  `json:"last_active_at_millis"`
    SignedUpAtMillis        int64                  `json:"signed_up_at_millis"`
    SelectedTeamID          string                 `json:"selected_team_id"`
    SelectedTeam            *SelectedTeam          `json:"selected_team,omitempty"`
}

// SelectedTeam содержит информацию о выбранной команде
type SelectedTeam struct {
    ID                     string                 `json:"id"`
    DisplayName            string                 `json:"display_name"`
    ProfileImageURL        string                 `json:"profile_image_url"`
    ClientMetadata         map[string]interface{} `json:"client_metadata"`
    ClientReadOnlyMetadata map[string]interface{} `json:"client_read_only_metadata"`
    ServerMetadata         map[string]interface{} `json:"server_metadata"`
    CreatedAtMillis        int64                  `json:"created_at_millis"`
}

// Pagination содержит информацию о пагинации
type Pagination struct {
    NextCursor string `json:"next_cursor"`
}
