// +build !windows

package dotfiles

import (
	"path/filepath"

	"github.com/hansrodtang/runcom/core"
	"github.com/mitchellh/go-homedir"
)

func Backup() {
	for _, v := range search {
		for _, file := range v {
			absolute, _ := homedir.Expand(file)
			if core.Valid(absolute) {
				_, filename := filepath.Split(absolute)
				symlink[file] = filename
			}
		}
	}
	core.Add(PluginName, symlink)
	core.Save()
}
