package cmd

import (
	"log"
	"os"

	"github.com/agaviria/timecraft/modules/configuration"
	"github.com/codegangsta/cli"
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	cliui "github.com/mitchellh/cli"
)

// Net is the web interface and api command.
var Net = cli.Command{
	Name:  "net",
	Usage: "./timecraft <Command> <Subcommand>",
	Subcommands: []cli.Command{
		{
			Name:   "run",
			Usage:  "initializes fasthttp server: » ./timecraft net run",
			Action: serveNet,
		},
	},
}

// serveNet will serve site and api
func serveNet(ctx *cli.Context) {
	// Load configuration and save

	ui := &cliui.BasicUi{Writer: os.Stdout, Reader: os.Stdin}

	// input must be prefixed with a colon (:)
	// i.e.  :8080
	domain, _ := ui.Ask("Static Server Port Domain:")

	if domain != "" {
		configuration.Configs.Domain = domain
		configuration.SaveConfig()
	}

	e := echo.New()
	r := pongor.GetRenderer()
	e.SetRenderer(r)

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.Static("/js/", "src/js")
	e.Static("/css/", "src/css")

	// start server
	log.Printf("Listening and serving.... port: %s\n", configuration.Configs.Domain)
	e.Run(fasthttp.New(configuration.Configs.Domain))
}
