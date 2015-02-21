// +build !windows

package dotfiles

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	// ConfigFile represents the file from where the configuration will be loaded.
	ConfigFile = "./plugins/dotfiles/dotfiles.yml"
	Name       = "[dotfiles]"
)

var search map[string][]string

var symlink map[string]string

func init() {
	symlink = make(map[string]string)

	logger := log.New(os.Stdout, Name+" ", 0)

	file, readErr := ioutil.ReadFile(ConfigFile)
	if readErr != nil {
		logger.Fatalf("Configuration error: %v \n", readErr)
	}

	parseErr := yaml.Unmarshal(file, &search)
	if parseErr != nil {
		logger.Fatalf("Parse error: %v \n", parseErr)
	}
}
