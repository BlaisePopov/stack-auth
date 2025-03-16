package sessions

// CreateSessionRequest содержит параметры для создания новой сессии
type CreateSessionRequest struct {
    // Идентификатор пользователя
    UserID string `json:"user_id"`
    
    // Время жизни сессии в миллисекундах (опционально)
    ExpiresInMillis *int64 `json:"expires_in_millis,omitempty"`
}
