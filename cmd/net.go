package cmd

import (
	log "github.com/Sirupsen/logrus"

	"github.com/agaviria/timecraft/modules/configuration"
	"github.com/codegangsta/cli"
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
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
	configuration.LoadConfig()

	log.Info("Initializing server, middleware, renderer and routes...")

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
	log.Infof("Listening and serving.... port: %s\n", configuration.Configs.Domain)
	e.Run(fasthttp.New(configuration.Configs.Domain))
}
