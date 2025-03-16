package api

import (
	"github.com/BlaisePopov/stack-auth/api/contactchannels"
	"github.com/BlaisePopov/stack-auth/api/oauth"
	"github.com/BlaisePopov/stack-auth/api/others"
	"github.com/BlaisePopov/stack-auth/api/otp"
	"github.com/BlaisePopov/stack-auth/api/password"
	"github.com/BlaisePopov/stack-auth/api/permissions"
	"github.com/BlaisePopov/stack-auth/api/projects"
	"github.com/BlaisePopov/stack-auth/api/sessions"
	"github.com/BlaisePopov/stack-auth/api/teams"
	"github.com/BlaisePopov/stack-auth/api/users"
	base_http_client "github.com/BlaisePopov/stack-auth/base-http-client"
)

type Client struct {
	ContactChannels *contactchannels.Client
	Oauth           *oauth.Client
	Others          *others.Client
	OTP             *otp.Client
	Password        *password.Client
	Permissions     *permissions.Client
	Projects        *projects.Client
	Sessions        *sessions.Client
	Teams           *teams.Client
	Users           *users.Client
}

func NewClient(config base_http_client.Config) *Client {
	baseHTTPClient := base_http_client.NewClient(config)

	return &Client{
		ContactChannels: contactchannels.NewClient(baseHTTPClient),
		Oauth:           oauth.NewClient(baseHTTPClient),
		Others:          others.NewClient(baseHTTPClient),
		OTP:             otp.NewClient(baseHTTPClient),
		Password:        password.NewClient(baseHTTPClient),
		Permissions:     permissions.NewClient(baseHTTPClient),
		Projects:        projects.NewClient(baseHTTPClient),
		Sessions:        sessions.NewClient(baseHTTPClient),
		Teams:           teams.NewClient(baseHTTPClient),
		Users:           users.NewClient(baseHTTPClient),
	}
}
