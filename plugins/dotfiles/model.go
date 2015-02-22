// +build !windows

package dotfiles

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	// ConfigFile represents the file from where the configuration will be loaded.
	ConfigFolder = "./plugins/dotfiles/data"
)

type dots struct {
	Name  string
	Files []string
}

var links []string

var storage map[string]string

func init() {
	storage = make(map[string]string)

	files, err := ioutil.ReadDir(ConfigFolder)
	if err != nil {
		out.Error(err)
	}
	for _, file := range files {
		var search dots

		data, readErr := ioutil.ReadFile(filepath.Join(ConfigFolder, file.Name()))
		if readErr != nil {
			out.Error(readErr)
		}

		parseErr := yaml.Unmarshal(data, &search)
		if parseErr != nil {
			out.Error(parseErr)
		}
		links = append(links, search.Files...)
	}
}
