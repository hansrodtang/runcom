package plugins

import (
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	plugins = make(map[string]Plugin)
}

type backupFunc func()
type restoreFunc func()
type commandOb *cobra.Command

// Plugin contains the plugin information
type Plugin struct {
	Name    string
	Restore restoreFunc
	Backup  backupFunc
	Command commandOb
}

var plugins map[string]Plugin

// Register adds plugin to plugin list.
func Register(name string, restore restoreFunc, backup backupFunc, command commandOb) {
	plugins[name] = Plugin{name, restore, backup, command}
}

// Get returns the plugin by name, if one exists.
func Get(name string) (Plugin, error) {
	p, ok := plugins[name]
	if !ok {
		return Plugin{}, errors.New("no plugin found")
	}
	return p, nil
}

// GetAll returns all registered plugins
func GetAll() map[string]Plugin {
	return plugins
}
