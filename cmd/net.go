package cmd

import (
	"html/template"
	"io"

	"github.com/agaviria/timecraft/modules/configuration"
	"github.com/codegangsta/cli"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

// Template is the struct which will be consumed by Render function.
type Template struct {
	templates *template.Template
}

// Net is the web interface and api command
var Net = cli.Command{
	Name:   "net",
	Usage:  "Initialize echo fasthttp server",
	Action: runNet,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "serve", "s",
			Value: "config/config.ini",
			Usage: "Configuration filepath",
		},
		cli.StringFlag{
			Name:  "port, p",
			Value: "8000",
			Usage: "Port number",
		},
	},
}

// Render implements Template struct.
func (t *Template) Render(writer io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(writer, name, data)
}

// serveNet will serve site and api
func serveNet(ctx *cli.Context) {
	// Load configurations
	configuration.LoadConf()

	e := echo.New()

	// Static assets
	e.Use("static", "src/static")

	// Middleware
	e.Use(m.Logger())
	e.Use(m.Recover())

	// Configure template engine
	e.SetRenderer(&Template{
		templates: template.Must(template.ParseGlob("src/views/*.html")),
	})

	e.GET("/api", func(c echo.Context) error {
		return c.String(200, "1")
	})
	e.Run(fasthttp.New(":8000"))
}
