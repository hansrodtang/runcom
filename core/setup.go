package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var print PrintFunc

func init() {
	print = Printer(Command)
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
			answer := Ask("Folder "+path+" does not exist. Use default?", true)
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
		answer := Ask("Folder "+directory+" does not exist. Create?", true)
		if answer {
			CreateDir(Directory())
		} else {
			fmt.Printf("%s can not run without storage directory\n", Name)
			os.Exit(1)
		}

	}
}
