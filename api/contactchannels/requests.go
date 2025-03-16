package contactchannels

// CreateContactChannelRequest представляет запрос на создание контактного канала
type CreateContactChannelRequest struct {
    UserID       string `json:"user_id"`
    Value        string `json:"value"`
    Type         string `json:"type"`
    UsedForAuth  bool   `json:"used_for_auth"`
    IsVerified   *bool  `json:"is_verified,omitempty"`
    IsPrimary    *bool  `json:"is_primary,omitempty"`
}

// VerifyRequest представляет запрос на подтверждение email
type VerifyRequest struct {
    Code string `json:"code"`
}

// CheckCodeRequest представляет запрос на проверку кода подтверждения
type CheckCodeRequest struct {
    Code string `json:"code"`
}

// UpdateContactChannelRequest представляет запрос на обновление контактного канала
type UpdateContactChannelRequest struct {
    Value        *string `json:"value,omitempty"`
    Type         *string `json:"type,omitempty"`
    UsedForAuth  *bool   `json:"used_for_auth,omitempty"`
    IsVerified   *bool   `json:"is_verified,omitempty"`
    IsPrimary    *bool   `json:"is_primary,omitempty"`
}

// SendCodeRequest представляет запрос на отправку кода подтверждения
type SendCodeRequest struct {
    CallbackURL string `json:"callback_url"`
}
