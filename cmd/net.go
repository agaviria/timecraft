package cmd

import (
	"log"

	"github.com/agaviria/timecraft/modules/configuration"
	"github.com/codegangsta/cli"
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

// Net is the web interface and api command
var Net = cli.Command{
	Name:   "net",
	Usage:  "Initialize echo fasthttp server",
	Action: serveNet,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "path",
			Value: "config.ini",
			Usage: "Configuration filepath",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "8000",
			Usage: "Port number",
		},
	},
}

// serveNet will serve site and api
func serveNet(ctx *cli.Context) {
	// Load configurations
	// TODO: add port to config.ini and bind to log below
	configuration.LoadConf()

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
	log.Println("Listening and serving....")
	e.Run(fasthttp.New(":8000"))
}
