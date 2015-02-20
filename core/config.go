package core

import "github.com/mitchellh/go-homedir"

var directory string = "~/.config/runcom"

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
