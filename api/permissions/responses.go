package permissions

// ListTeamPermissionsResponse представляет ответ на запрос списка командных разрешений
type ListTeamPermissionsResponse struct {
    Items      []TeamPermission `json:"items"`
    Pagination Pagination        `json:"pagination"`
}

// TeamPermission содержит информацию о командном разрешении пользователя
type TeamPermission struct {
    ID           string `json:"id"`
    TeamID       string `json:"team_id"`
    UserID       string `json:"user_id"`
}

// Pagination содержит информацию о пагинации
type Pagination struct {
    NextCursor string `json:"next_cursor"`
}

// GrantTeamPermissionResponse представляет ответ на запрос выдачи разрешения
type GrantTeamPermissionResponse struct {
    ID     string `json:"id"`
    TeamID string `json:"team_id"`
    UserID string `json:"user_id"`
}

// RevokeTeamPermissionResponse представляет ответ на запрос отзыва разрешения
type RevokeTeamPermissionResponse struct {
    Success bool `json:"success"`
}
