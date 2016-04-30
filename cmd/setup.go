package cmd

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/agaviria/timecraft/modules/configuration"
	"github.com/codegangsta/cli"
)

// Setup command to install the configuration and reset the database
var Setup = cli.Command{
	Name:  "setup",
	Usage: "./timecraft <Command> <Subcommand>",
	Subcommands: []cli.Command{
		{
			Name:   "install",
			Usage:  "installs configurations: » ./timecraft setup install",
			Action: install,
		},
		{
			Name:   "reset",
			Usage:  "resets database: » ./timecraft setup reset",
			Action: reset,
		},
	},
}

// install will create the database and run all migrations
func install(ctx *cli.Context) {
	configuration.LoadConfig()
	log.Infoln("Loading configuration setup...")
	// Save configurations to persistent storage
	configuration.SaveConfig()
	log.Infoln("Configuration setup has finished succesfully.")

	// store.NewStormUserStore()
}

// reset will drop the database schema and run all migrations again
func reset(ctx *cli.Context) {
	configuration.LoadConfig()
	err := os.Remove(configuration.Configs.Store)

	if err != nil {
		log.Fatalf("Error Removing File: %s\n", err)
		return
	}
	fmt.Printf("Database %s has executed a reset", configuration.Configs.Store)
}
