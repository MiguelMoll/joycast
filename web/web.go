package web

import (
	"html/template"
	"io"

	"github.com/MiguelMoll/joycast/realm"
	"github.com/gobuffalo/packr"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type option func(s *server) error

type server struct {
	address string
	echo    *echo.Echo
	users   *realm.UserService
}

func New(opts ...option) (*server, error) {
	s := &server{}

	s.echo = echo.New()
	s.echo.HideBanner = true
	s.echo.Renderer = newRenderer()

	// Middleware
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	routes(s)

	return s, nil
}

func (s *server) Start(address string) error {
	return s.echo.Start(address)
}

func UserService(u *realm.UserService) option {
	return func(s *server) error {
		s.users = u
		return nil
	}
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
