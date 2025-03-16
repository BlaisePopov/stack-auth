package sessions

// CreateSessionResponse содержит ответ сервера при создании сессии
type CreateSessionResponse struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}

// SignOutResponse содержит ответ сервера при выходе из сессии
type SignOutResponse struct {
    Success bool `json:"success"`
}

// RefreshAccessTokenResponse содержит ответ сервера при обновлении токена доступа
type RefreshAccessTokenResponse struct {
    AccessToken string `json:"access_token"`
}
