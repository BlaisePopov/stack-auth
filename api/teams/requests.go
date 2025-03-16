package teams

// CreateTeamRequest содержит данные для создания команды
type CreateTeamRequest struct {
	DisplayName            string                 `json:"display_name"`
	CreatorUserID          *string                `json:"creator_user_id,omitempty"`
	ClientReadOnlyMetadata map[string]interface{} `json:"client_read_only_metadata,omitempty"`
	ServerMetadata         map[string]interface{} `json:"server_metadata,omitempty"`
	ProfileImageURL        *string                `json:"profile_image_url,omitempty"`
	ClientMetadata         map[string]interface{} `json:"client_metadata,omitempty"`
}

// UpdateTeamRequest содержит данные для обновления команды
type UpdateTeamRequest struct {
	DisplayName            *string                `json:"display_name,omitempty"`
	ProfileImageURL        *string                `json:"profile_image_url,omitempty"`
	ClientMetadata         map[string]interface{} `json:"client_metadata,omitempty"`
	ClientReadOnlyMetadata map[string]interface{} `json:"client_read_only_metadata,omitempty"`
	ServerMetadata         map[string]interface{} `json:"server_metadata,omitempty"`
}

// UpdateTeamMemberProfileRequest содержит данные для обновления команды
type UpdateTeamMemberProfileRequest struct {
	DisplayName     *string `json:"display_name,omitempty"`
	ProfileImageURL *string `json:"profile_image_url,omitempty"`
}

// SendInviteEmailRequest содержит данные для отправки приглашения
type SendInviteEmailRequest struct {
	TeamID      string `json:"team_id"`
	Email       string `json:"email"`
	CallbackURL string `json:"callback_url"`
}

// AcceptInviteRequest содержит код приглашения
type AcceptInviteRequest struct {
	Code string `json:"code"`
}
