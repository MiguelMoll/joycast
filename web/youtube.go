package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

// Authenticate is the first step in the process to get a
// users permission to access their YouTube account
// This will redirect to Google sign-in page for permission
func Authenticate(c echo.Context) error {
	config, err := newConfig()
	if err != nil {
		return err
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, authURL)
}

// Authorized handles the callback from YouTube with
// a state and code to create a token
func Authorized(c echo.Context) error {
	config, err := newConfig()
	if err != nil {
		return err
	}

	// TODO: check state value too!
	// https://developers.google.com/youtube/v3/guides/auth/server-side-web-apps

	tok, err := config.Exchange(context.Background(), c.QueryParam("code"))
	if err != nil {
		c.Error(err)
	}

	// TODO: Save token!
	fmt.Println(tok)
	return nil
}

// newClient creates a new YouTube service client
// for making API calls
func newClient() (*youtube.Service, error) {
	config, err := newConfig()
	if err != nil {
		return nil, err
	}

	// should get token from storage
	token := &oauth2.Token{}

	client := config.Client(context.Background(), token)
	return youtube.New(client)
}

// newConfig creates a new OAuth config
func newConfig() (*oauth2.Config, error) {
	secret := []byte("TODO")
	return google.ConfigFromJSON(secret, youtube.YoutubeReadonlyScope)
}
