// +build !windows

package dotfiles

import "github.com/hansrodtang/runcom/core"

func Restore() {
	err := core.Get(PluginName, &storage)
	if err != nil {
		out.Error(err)
	}
	for k, v := range storage {
		out.Print(v, "=>", k)
	}
}
