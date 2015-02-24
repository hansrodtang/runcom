package core

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

var directory = "~/.config/runcom"
var file = "storage.json"

const Command = "runcom"
const Name = "Runcom"

func Directory() string {
	dir, err := homedir.Expand(directory)
	if err != nil {
		return ""
	}
	return dir
}

func PluginDirectory(plugin string) string {
	dir := Directory()
	return filepath.Join(dir, plugin)
}
