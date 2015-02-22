package plugins

import (
	"errors"

	"github.com/spf13/cobra"
)

func init() {
	plugins = make(map[string]Plugin)
}

type BackupFunc func()
type RestoreFunc func()
type CommandOb *cobra.Command

type Plugin struct {
	Name    string
	Restore RestoreFunc
	Backup  BackupFunc
	Command CommandOb
}

var plugins map[string]Plugin

func Register(name string, restore RestoreFunc, backup BackupFunc, command CommandOb) {
	plugins[name] = Plugin{name, restore, backup, command}
}

func Get(name string) (Plugin, error) {
	p, ok := plugins[name]
	if !ok {
		return Plugin{}, errors.New("no plugin found")
	}
	return p, nil
}

func GetAll() map[string]Plugin {
	return plugins
}
