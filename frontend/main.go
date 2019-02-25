package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/MiguelMoll/joycast/audio"
)

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
