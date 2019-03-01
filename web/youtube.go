package web

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

// Authenticate is the first step in the process to get a
// users permission to access their YouTube account
// This will redirect to Google sign-in page for permission
func YoutubeAuthenticate(c echo.Context) error {
	config, err := newConfig()
	if err != nil {
		return err
	}

	// TODO: Handle state token here for security
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, authURL)
}

// Authorized handles the callback from YouTube with
// a state and code to create a token
func YoutubeAuthorized(c echo.Context) error {
	config, err := newConfig()
	if err != nil {
		return err
	}

	// TODO: check state value too!
	// https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps

	tok, err := config.Exchange(context.Background(), c.QueryParam("code"))
	if err != nil {
		return err
	}

	u, err := Store.GetUser(1)
	if err != nil {
		return err
	}

	u.OauthToken = tok

	if err := Store.SaveUser(u); err != nil {
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/signin")
}

// newClient creates a new YouTube service client
// for making API calls
func newClient(token *oauth2.Token) (*youtube.Service, error) {
	config, err := newConfig()
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), token)
	return youtube.New(client)
}

// newConfig creates a new OAuth config
func newConfig() (*oauth2.Config, error) {
	secret := []byte(`{"web":{"client_id":"1003910230744-gqvdgs38d2kvllslba3jvddsgvo5pckq.apps.googleusercontent.com","project_id":"projectcast","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"4WqhzHGa9_eJnVovcHQsS_CY","redirect_uris":["http://localhost:8000/youtube/authorized"],"javascript_origins":["http://www.sire.ninja"]}}`)
	return google.ConfigFromJSON(secret, youtube.YoutubeReadonlyScope, youtube.YoutubeUploadScope)
}
