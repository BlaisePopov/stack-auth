package projects

// ProjectConfig содержит конфигурационные настройки проекта
type ProjectConfig struct {
    ClientTeamCreationEnabled bool `json:"client_team_creation_enabled"`
    ClientUserDeletionEnabled bool `json:"client_user_deletion_enabled"`
    CredentialEnabled         bool `json:"credential_enabled"`
    MagicLinkEnabled          bool `json:"magic_link_enabled"`
    PasskeyEnabled            bool `json:"passkey_enabled"`
    SignUpEnabled             bool `json:"sign_up_enabled"`
}

// GetCurrentProjectResponse содержит полный ответ сервера о текущем проекте
type GetCurrentProjectResponse struct {
    Config        ProjectConfig `json:"config"`
    DisplayName   string        `json:"display_name"`
    ID            string        `json:"id"`
}
