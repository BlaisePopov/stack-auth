package otp

// SignInWithCodeRequest содержит параметры запроса для входа с кодом
type SignInWithCodeRequest struct {
    Code string `json:"code"`
}

// SendSignInCodeRequest содержит параметры запроса для отправки кода
type SendSignInCodeRequest struct {
    Email       string `json:"email"`
    CallbackURL string `json:"callback_url"`
}

// MFASignInRequest содержит параметры запроса для MFA аутентификации
type MFASignInRequest struct {
    Type string `json:"type"`
    TOTP string `json:"totp"`
    Code string `json:"code"`
}

// CheckSignInCodeRequest содержит параметры запроса проверки кода
type CheckSignInCodeRequest struct {
    Code string `json:"code"`
}
