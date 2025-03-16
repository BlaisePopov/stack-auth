package others

// TeamInvitation содержит информацию о приглашении в команду
type TeamInvitation struct {
    ExpiresAtMillis float64 `json:"expires_at_millis"`
    ID              string  `json:"id"`
    RecipientEmail  string  `json:"recipient_email"`
    TeamID          string  `json:"team_id"`
}

// Pagination содержит информацию о пагинации
type Pagination struct {
    NextCursor string `json:"next_cursor"`
}

// ListTeamInvitationsResponse представляет ответ на запрос списка приглашений
type ListTeamInvitationsResponse struct {
    Items      []TeamInvitation `json:"items"`
    Pagination Pagination        `json:"pagination"`
}

// DeleteTeamInvitationResponse представляет ответ на запрос удаления приглашения
type DeleteTeamInvitationResponse struct {
    Success bool `json:"success"`
}

// ConfirmNeonTransferCheckResponse представляет ответ на проверку передачи проекта Neon
type ConfirmNeonTransferCheckResponse struct {
    IsCodeValid bool `json:"is_code_valid"`
}
