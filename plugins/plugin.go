package plugins

import (
	"errors"

	"github.com/spf13/cobra"
)

type ExporterFunc func()
type ImporterFunc func()
type CommandOb *cobra.Command

type plugin struct {
	Name    string
	Import  ImporterFunc
	Export  ExporterFunc
	Command CommandOb
}

// Formats is the list of registered formats.
var plugins []plugin

func Register(name string, importer ImporterFunc, exporter ExporterFunc, command CommandOb) {
	plugins = append(plugins, plugin{name, importer, exporter, command})
}

func Get(name string) (plugin, error) {
	for _, p := range plugins {
		if p.Name == name {
			return p, nil
		}
	}
	return plugin{}, errors.New("no plugin found")
}

func GetAll() []plugin {
	return plugins
}
