package users

// CreateUserRequest содержит данные для создания пользователя
type CreateUserRequest struct {
	DisplayName             string                 `json:"display_name,omitempty"`
	ProfileImageURL         string                 `json:"profile_image_url,omitempty"`
	ClientMetadata          map[string]interface{} `json:"client_metadata,omitempty"`
	ClientReadOnlyMetadata  map[string]interface{} `json:"client_read_only_metadata,omitempty"`
	ServerMetadata          map[string]interface{} `json:"server_metadata,omitempty"`
	PrimaryEmail            string                 `json:"primary_email,omitempty"`
	PrimaryEmailVerified    *bool                  `json:"primary_email_verified,omitempty"`
	PrimaryEmailAuthEnabled *bool                  `json:"primary_email_auth_enabled,omitempty"`
	Password                string                 `json:"password,omitempty"`
	PasswordHash            string                 `json:"password_hash,omitempty"`
	TOTPSecretBase64        string                 `json:"totp_secret_base64,omitempty"`
}

// UpdateUserRequest содержит данные для обновления пользователя
type UpdateUserRequest struct {
	DisplayName             string                 `json:"display_name,omitempty"`
	ProfileImageURL         string                 `json:"profile_image_url,omitempty"`
	ClientMetadata          map[string]interface{} `json:"client_metadata,omitempty"`
	ClientReadOnlyMetadata  map[string]interface{} `json:"client_read_only_metadata,omitempty"`
	ServerMetadata          map[string]interface{} `json:"server_metadata,omitempty"`
	PrimaryEmail            string                 `json:"primary_email,omitempty"`
	PrimaryEmailVerified    *bool                  `json:"primary_email_verified,omitempty"`
	PrimaryEmailAuthEnabled *bool                  `json:"primary_email_auth_enabled,omitempty"`
	Password                string                 `json:"password,omitempty"`
	PasswordHash            string                 `json:"password_hash,omitempty"`
	TOTPSecretBase64        string                 `json:"totp_secret_base64,omitempty"`
	SelectedTeamID          string                 `json:"selected_team_id,omitempty"`
}
