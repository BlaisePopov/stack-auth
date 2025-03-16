package otp

// AuthResponse содержит данные аутентификационного ответа
type AuthResponse struct {
    AccessToken  string `json:"access_token"`
    IsNewUser    bool   `json:"is_new_user"`
    RefreshToken string `json:"refresh_token"`
    UserID       string `json:"user_id"`
}

// SendSignInCodeResponse содержит ответ на запрос отправки кода
type SendSignInCodeResponse struct {
    Nonce string `json:"nonce"`
}

// CheckSignInCodeResponse содержит результат проверки кода
type CheckSignInCodeResponse struct {
    IsCodeValid bool `json:"is_code_valid"`
}
