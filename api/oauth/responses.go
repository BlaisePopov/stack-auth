package oauth

// TokenResponse представляет ответ с OAuth токенами
type TokenResponse struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
