package password

// UpdatePasswordRequest содержит данные для обновления пароля
type UpdatePasswordRequest struct {
    OldPassword string `json:"old_password"`
    NewPassword string `json:"new_password"`
}

// SignUpWithEmailRequest содержит данные для регистрации с email
type SignUpWithEmailRequest struct {
    Email                 string `json:"email"`
    Password              string `json:"password"`
    VerificationCallbackURL string `json:"verification_callback_url"`
}

// SignInRequest содержит данные для входа с email и паролем
type SignInRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// SetPasswordRequest содержит данные для установки нового пароля
type SetPasswordRequest struct {
    Password string `json:"password"`
}

// SendResetCodeRequest содержит данные для отправки кода сброса
type SendResetCodeRequest struct {
    Email       string `json:"email"`
    CallbackURL string `json:"callback_url"`
}

// ResetPasswordRequest содержит данные для сброса пароля
type ResetPasswordRequest struct {
    Password string `json:"password"`
    Code     string `json:"code"`
}

// CheckCodeRequest содержит данные для проверки кода сброса
type CheckCodeRequest struct {
    Code string `json:"code"`
}
