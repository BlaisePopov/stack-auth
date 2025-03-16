package password

// UpdatePasswordResponse содержит ответ на запрос обновления пароля
type UpdatePasswordResponse struct {
    Success bool `json:"success"`
}

// SignUpResponse содержит ответ на запрос регистрации
type SignUpResponse struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    UserID       string `json:"user_id"`
}

// SignInResponse содержит ответ на запрос входа
type SignInResponse struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    UserID       string `json:"user_id"`
}

// SetPasswordResponse содержит ответ на запрос установки пароля
type SetPasswordResponse struct {
    Success bool `json:"success"`
}

// SendResetCodeResponse содержит ответ на запрос отправки кода сброса
type SendResetCodeResponse struct {
    Success string `json:"success"`
}

// ResetPasswordResponse содержит ответ на запрос сброса пароля
type ResetPasswordResponse struct {
    Success bool `json:"success"`
}

// CheckCodeResponse содержит ответ на проверку кода сброса
type CheckCodeResponse struct {
    IsCodeValid bool `json:"is_code_valid"`
}
