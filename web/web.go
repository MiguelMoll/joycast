package web

import (
	"html/template"
	"io"

	"github.com/MiguelMoll/joycast/storage"
	"github.com/gobuffalo/packr"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Options struct {
	Address string
	Store   storage.Store
}

var Store storage.Store

func Run(opts Options) {
	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Renderer = newRenderer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes(e)

	Store = opts.Store

	// Start server
	e.Logger.Fatal(e.Start(opts.Address))
}

func newRenderer() echo.Renderer {
	return &renderer{
		box:      packr.NewBox("./templates"),
		rootTmpl: template.New("web"),
		subTmpl:  map[string]*template.Template{},
	}
}

type renderer struct {
	box      packr.Box
	rootTmpl *template.Template
	subTmpl  map[string]*template.Template
}

func (r *renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	t, ok := r.subTmpl[name]
	if !ok {
		var err error
		t, err = r.rootTmpl.Parse(r.box.String(name))
		if err != nil {
			return err
		}
		r.subTmpl[name] = t
	}

	return t.Execute(w, data)
}
