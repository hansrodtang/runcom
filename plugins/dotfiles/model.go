// +build !windows

package dotfiles

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	// ConfigFile represents the file from where the configuration will be loaded.
	configFolder = "./plugins/dotfiles/data"
)

type dots struct {
	Name  string
	Files []string
}

var links []string

var model map[string]string

func init() {
	model = make(map[string]string)

	files, err := ioutil.ReadDir(configFolder)
	if err != nil {
		out.Error(err)
	}
	for _, file := range files {
		data, readErr := ioutil.ReadFile(filepath.Join(configFolder, file.Name()))
		if readErr != nil {
			out.Error(readErr)
		}

		var search dots
		parseErr := yaml.Unmarshal(data, &search)
		if parseErr != nil {
			out.Error(parseErr)
		}
		links = append(links, search.Files...)
	}
}
