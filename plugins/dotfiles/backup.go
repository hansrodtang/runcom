// +build !windows

package dotfiles

import (
	"path/filepath"

	"github.com/hansrodtang/runcom/backends"
	"github.com/hansrodtang/runcom/core"
	"github.com/mitchellh/go-homedir"
)

func backup() {
	storage := backends.Get()
	storage.Read()

	for _, file := range links {
		absolute, _ := homedir.Expand("~/" + file)
		if core.Valid(absolute) {

			_, filename := filepath.Split(absolute)
			model[file] = filename
		}
	}
	storage.Add(pluginName, model)
	storage.Save()
}
