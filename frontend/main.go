package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"

	"github.com/MiguelMoll/joycast/audio"
)

var config *oauth2.Config

var secret = []byte(`{"web":{"client_id":"1003910230744-gqvdgs38d2kvllslba3jvddsgvo5pckq.apps.googleusercontent.com","project_id":"projectcast","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"4WqhzHGa9_eJnVovcHQsS_CY","redirect_uris":["http://www.sire.ninja/auth"],"javascript_origins":["http://www.sire.ninja"]}}`)

func signin(c echo.Context) error {
	var err error
	config, err = google.ConfigFromJSON(secret, youtube.YoutubeReadonlyScope)
	if err != nil {
		c.Error(err)
	}
	fmt.Println(err)
	fmt.Println(config)

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, authURL)
}

var code = ""

func auth(c echo.Context) error {
	code = c.QueryParam("code")
	return c.Redirect(http.StatusTemporaryRedirect, "/ytinfo")
}

func ytinfo(c echo.Context) error {
	if code == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/signin")
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		c.Error(err)
	}

	fmt.Println("token:", tok)

	client := config.Client(context.Background(), tok)
	service, err := youtube.New(client)
	if err != nil {
		c.Error(err)
	}
	call := service.Channels.List("snippet,contentDetails,statistics")
	call = call.ForUsername("radioact1ve")
	response, err := call.Do()
	if err != nil {
		c.Error(err)
	}

	fmt.Println("Response:", response)
	fmt.Println("Len:", len(response.Items))
	output := fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.\n",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount)

	return c.String(http.StatusOK, output)
}

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	as := audio.New(audio.Config{})
	if err := as.FromVideoFile(dst.Name(), fmt.Sprintf("%s.mp3", name)); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s.</p>", file.Filename, name))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e := echo.New()

	box := packr.NewBox("./templates")
	t := &Renderer{
		box: box,
	}

	e.Renderer = t
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error { return c.Render(http.StatusOK, "index.html", nil) })
	e.GET("/signin", func(c echo.Context) error { return c.Render(http.StatusOK, "signin.html", nil) })
	e.GET("/auth", auth)
	e.GET("/ytinfo", ytinfo)
	e.POST("/signin", signin)
	e.GET("/google8151a84b9af0aeb1.html", func(c echo.Context) error { return c.Render(http.StatusOK, "google8151a84b9af0aeb1.html", nil) })
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

type Renderer struct {
	box packr.Box
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t, err := template.New(name).Parse(r.box.String(name))
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}
