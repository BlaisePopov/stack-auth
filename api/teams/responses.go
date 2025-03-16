package teams

// TeamResponse содержит информацию о команде
type TeamResponse struct {
	ID                     string                 `json:"id"`
	DisplayName            string                 `json:"display_name"`
	ProfileImageURL        string                 `json:"profile_image_url"`
	CreatedAtMillis        int64                  `json:"created_at_millis"`
	ClientMetadata         map[string]interface{} `json:"client_metadata"`
	ClientReadOnlyMetadata map[string]interface{} `json:"client_read_only_metadata"`
	ServerMetadata         map[string]interface{} `json:"server_metadata"`
}

// ListTeamsResponse содержит список команд с пагинацией
type ListTeamsResponse struct {
	Items      []TeamResponse `json:"items"`
	Pagination Pagination     `json:"pagination"`
}

// TeamMemberProfileResponse содержит профиль участника команды
type TeamMemberProfileResponse struct {
	UserID          string      `json:"user_id"`
	TeamID          string      `json:"team_id"`
	DisplayName     string      `json:"display_name"`
	ProfileImageURL string      `json:"profile_image_url"`
	User            UserDetails `json:"user"`
}

// Pagination содержит информацию о пагинации
type Pagination struct {
	NextCursor string `json:"next_cursor"`
}

// SuccessResponse содержит статус операции
type SuccessResponse struct {
	Success bool   `json:"success"`
	ID      string `json:"id,omitempty"`
}

// UserDetails содержит детальную информацию о пользователе
type UserDetails struct {
	ID                      string                 `json:"id"`
	DisplayName             string                 `json:"display_name"`
	PrimaryEmail            string                 `json:"primary_email"`
	ProfileImageURL         string                 `json:"profile_image_url"`
	SignedUpAtMillis        int64                  `json:"signed_up_at_millis"`
	LastActiveAtMillis      int64                  `json:"last_active_at_millis"`
	PrimaryEmailVerified    bool                   `json:"primary_email_verified"`
	PrimaryEmailAuthEnabled bool                   `json:"primary_email_auth_enabled"`
	ClientMetadata          map[string]interface{} `json:"client_metadata"`
	ClientReadOnlyMetadata  map[string]interface{} `json:"client_read_only_metadata"`
	ServerMetadata          map[string]interface{} `json:"server_metadata"`
}

// ListTeamMembersResponse содержит список профилей участников
type ListTeamMembersResponse struct {
	Items      []TeamMemberProfileResponse `json:"items"`
	Pagination Pagination                  `json:"pagination"`
}

// TeamMembershipResponse содержит информацию о членстве
type TeamMembershipResponse struct {
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

// InvitationDetailsResponse содержит детали приглашения
type InvitationDetailsResponse struct {
	TeamID          string `json:"team_id"`
	TeamDisplayName string `json:"team_display_name"`
}

// CheckCodeResponse содержит результат проверки кода
type CheckCodeResponse struct {
	IsCodeValid bool `json:"is_code_valid"`
}
