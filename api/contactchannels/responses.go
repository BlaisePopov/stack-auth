package contactchannels

// ListContactChannelsResponse представляет ответ со списком контактных каналов
type ListContactChannelsResponse struct {
    Items      []ContactChannel `json:"items"`
    Pagination struct {
        NextCursor string `json:"next_cursor"`
    } `json:"pagination"`
}

// ContactChannelResponse представляет ответ с данными контактного канала
type ContactChannelResponse struct {
    ID           string `json:"id"`
    UserID       string `json:"user_id"`
    Value        string `json:"value"`
    Type         string `json:"type"`
    UsedForAuth  bool   `json:"used_for_auth"`
    IsVerified   bool   `json:"is_verified"`
    IsPrimary    bool   `json:"is_primary"`
}

// VerifyResponse представляет ответ на подтверждение email
type VerifyResponse struct {
    Success bool `json:"success"`
}

// CheckCodeResponse представляет ответ на проверку кода подтверждения
type CheckCodeResponse struct {
    IsCodeValid bool `json:"is_code_valid"`
}

// DeleteResponse представляет ответ на удаление контактного канала
type DeleteResponse struct {
    Success bool `json:"success"`
}

// SendCodeResponse представляет ответ на отправку кода подтверждения
type SendCodeResponse struct {
    Success bool `json:"success"`
}

// ContactChannel представляет структуру контактного канала
type ContactChannel struct {
    ID           string `json:"id"`
    UserID       string `json:"user_id"`
    Value        string `json:"value"`
    Type         string `json:"type"`
    UsedForAuth  bool   `json:"used_for_auth"`
    IsVerified   bool   `json:"is_verified"`
    IsPrimary    bool   `json:"is_primary"`
}
