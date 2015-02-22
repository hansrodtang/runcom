// +build !windows

package dotfiles

import (
	"path/filepath"

	"github.com/hansrodtang/runcom/core"
	"github.com/mitchellh/go-homedir"
)

func Backup() {
	for _, file := range links {
		absolute, _ := homedir.Expand("~/" + file)
		if core.Valid(absolute) {

			_, filename := filepath.Split(absolute)
			storage[file] = filename
		}
	}
	core.Add(PluginName, storage)
	core.Save()
}
