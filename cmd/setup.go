package cmd

import (
	"fmt"
	"os"

	"github.com/agaviria/timecraft/modules/configuration"
	"github.com/codegangsta/cli"
	cliui "github.com/mitchellh/cli"
)

// Setup command to install the configuration and reset the database
var Setup = cli.Command{
	Name:  "setup",
	Usage: "timecraft setup tools",
	Subcommands: []cli.Command{
		{
			Name:   "install",
			Usage:  "Install timecraft",
			Action: install,
		},
		{
			Name:   "reset",
			Usage:  "Resets the database",
			Action: reset,
		},
	},
}

// install will create the database and run all migrations
func install(ctx *cli.Context) {
	configuration.LoadConf()
	ui := &cliui.BasicUi{Writer: os.Stdout, Reader: os.Stdin}

	ps, _ := ui.Ask("Database filename:")

	if ps != "" {
		// Save configurations to persistent storage
		configuration.Configs.Store = ps
		configuration.SaveConfig()
	}
}

// reset will drop the database schema and run all migrations again
func reset(ctx *cli.Context) {
	configuration.LoadConf()

	err := os.Remove(configuration.Configs.Store)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Database %s reset", configuration.Configs.Store)
}
