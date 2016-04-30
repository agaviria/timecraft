package configuration

import (
	log "github.com/Sirupsen/logrus"
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

type Logger struct {
	FileName   string
	ConfigMode uint32
}

type Static struct {
	Domain    string // i.e. ":8080"
	Path      string // i.e. "src"
	Templates string // i.e. "src/views"
}

type ConfigPath struct {
	LogPath string // i.e. "./timecraft/log"
	DBPath  string
	TmpPath string
}

type configurations struct {
	*Database
	*Logger
	*Static
	*ConfigPath
}

// loads all the configurations to config.ini file
func LoadConfig() error {
	// TODO: log FileName and db Store are being loaded from utils.  Fix this!
	c := &configurations{
		&Database{Store: "timecraft.db"},
		&Logger{FileName: "timecraft.log", ConfigMode: 0640},
		&Static{Domain: ":8000", Path: "src", Templates: "src/views"},
		&ConfigPath{LogPath: "./timecraft/log", DBPath: "./timecraft/db", TmpPath: "./timecraft/tmp"},
	}

	Output = render.New(render.Options{IndentJSON: true})

	// cfg is the path of the config file
	// LooseLoad ignores nonexistent files without error return
	cfg, err := ini.LooseLoad("config.ini")
	log.Infoln("Applying configuration file settings...")
	err = ini.ReflectFrom(cfg, c)
	Configs = c

	if err != nil {
		log.Fatalln("Cannot Load configurations to config.ini file")
	}
	return err
}

// GetLogDir get key log path from config.ini
func GetLogDir() string {
	valueDir := cfg.Section("ConfigPath").Key("LogPath").String()
	return valueDir
}

func GetLogName() string {
	valueName := cfg.Section("Logger").Key("FileName").String()
	return valueName
}

// Gets Log Config mode 0755 used to create FileLog.Out
func GetLogMode() uint {
	valueMode, _ := cfg.Section("Logger").Key("ConfigMode").Uint()
	return valueMode
}

// SaveConfig saves the current configurations to the ini file
func SaveConfig() {
	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, Configs)

	if err != nil {
		log.Fatalln("Cannot save to config.ini file: ", os.Stderr, err.Error()+"\n")
	}
	cfg.SaveTo("config.ini")
	log.Infoln("Succesfully saved configurations into config file")
}
