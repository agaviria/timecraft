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

	cfg *ini.File = nil
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

// loads all the configurations to config.ini file
func init() {
	c := &configurations{
		&Database{Store: "timecraft.db"},
		&Static{Domain: ":8000", Path: "src", Templates: "src/views"},
	}

	Output = render.New(render.Options{IndentJSON: true})

	// cfg is the path of the config file
	// LooseLoad ignores nonexistent files without error return
	cfg, err := ini.LooseLoad("config.ini")
	err = ini.ReflectFrom(cfg, c)
	Configs = c

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}

// GetConfigKey loads a key from a section of our config.ini
func GetConfigKey(section, key string) *ini.Key {
	val, err := cfg.Section(section).GetKey(key)
	if err != nil {
		return nil
	}
	return val
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
