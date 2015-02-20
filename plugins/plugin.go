package plugins

import (
	"errors"

	"github.com/spf13/cobra"
)

type BackupFunc func()
type RestoreFunc func()
type CommandOb *cobra.Command

type plugin struct {
	Name    string
	Restore RestoreFunc
	Backup  BackupFunc
	Command CommandOb
}

// Formats is the list of registered formats.
var plugins []plugin

func Register(name string, restore RestoreFunc, backup BackupFunc, command CommandOb) {
	plugins = append(plugins, plugin{name, restore, backup, command})
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
