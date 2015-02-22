package core

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var out Printer

func init() {
	out = NewPrinter(Command)
}

func init() {
	//viper.SetConfigName("config")
	viper.SetDefault("directory", Directory())
	viper.SetDefault("storage", filepath.Join(Directory(), file))

	viper.SetEnvPrefix(Command)
	viper.BindEnv("path")

	setup()
}

func setup() {
	path := viper.GetString("path")
	if len(path) > 0 {
		if !Exists(path) {
			answer := out.Ask("Folder "+path+" does not exist. Use default?", true)
			if answer {
				directory = path
				CreateDir(Directory())
				return
			}
		} else {
			directory = path
			return
		}
	}
	if !Exists(Directory()) {
		answer := out.Ask("Folder "+directory+" does not exist. Create?", true)
		if answer {
			CreateDir(Directory())
		} else {
			out.Printf("%s can not run without storage directory\n", Name)
			os.Exit(1)
		}

	}
}
