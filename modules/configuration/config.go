package configuration

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
	"github.com/unrolled/render"
)

var (
	// Configs contains all of the configuration settings
	Configs *configurations

	// Output is the render output
	Output *render.Render
)

type Database struct {
	// Store is the persistent storage filename
	Store string
}

type Static struct {
	Domain    string // i.e. ":8080"
	Path      string // i.e. "src"
	Templates string // i.e. "src/views"
}

type configurations struct {
	*Database
	*Static
}

// LoadConf loads all the configurations from the ini file
func LoadConf() {
	c := &configurations{
		&Database{Store: "timecraft.db"},
		&Static{Domain: ":8000", Path: "src", Templates: "src/views"},
	}

	Output = render.New(render.Options{IndentJSON: true})

	// cfg is the path of the config file
	// LooseLoad ignores nonexistent files without error return
	cfg, err := ini.LooseLoad("config.ini")
	err = ini.ReflectFrom(cfg, c)

	// Map configurations to struct
	// cfg.MapTo(c)
	Configs = c

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}

// SaveConfig saves the current configurations to the ini file
func SaveConfig() {
	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, Configs)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	cfg.SaveTo("config.ini")
}
