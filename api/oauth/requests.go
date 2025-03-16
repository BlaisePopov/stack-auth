package oauth

// TokenRequest представляет запрос для получения OAuth токена
type TokenRequest struct {
	// Тип grant для OAuth-потока (authorization_code, refresh_token и т.д.)
	GrantType string `json:"grant_type"`
}

// AuthorizeQuery представляет запрос для инициализации OAuth-потока
type AuthorizeQuery struct {
	Type                     string `json:"type,omitempty"`
	Token                    string `json:"token,omitempty"`
	ClientID                 string `json:"client_id"`
	ClientSecret             string `json:"client_secret"`
	RedirectURI              string `json:"redirect_uri"`
	Scope                    string `json:"scope"`
	State                    string `json:"state"`
	GrantType                string `json:"grant_type"`
	CodeChallenge            string `json:"code_challenge"`
	CodeChallengeMethod      string `json:"code_challenge_method"`
	ResponseType             string `json:"response_type"`
	ProviderScope            string `json:"provider_scope,omitempty"`
	ErrorRedirectURI         string `json:"error_redirect_uri,omitempty"`
	AfterCallbackRedirectURL string `json:"after_callback_redirect_url,omitempty"`
}
