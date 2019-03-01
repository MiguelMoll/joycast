package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func routes(s *server) {
	s.echo.GET("/", s.hello)
	s.echo.GET("/signin", s.signin)

	// youtube endpoints
	s.echo.GET("/youtube/authenticate", s.YoutubeAuthenticate)
	s.echo.GET("/youtube/authorized", s.YoutubeAuthorized)
}

// Handler
func (s *server) hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func (s *server) signin(c echo.Context) error {
	user, err := s.users.Get(1)
	if err != nil {
		return err
	}

	if user.OauthToken == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/youtube/authenticate")
	}

	client, err := newClient(user.OauthToken)
	if err != nil {
		return err
	}

	call := client.Channels.List("snippet,contentDetails,statistics")
	call = call.ForUsername("radioact1ve")
	response, err := call.Do()
	if err != nil {
		c.Error(err)
	}

	output := fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.\n",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount)

	return c.Render(http.StatusOK, "index.html", output)
}
