// +build !windows

package dotfiles

import "github.com/hansrodtang/runcom/core"

func restore() {
	err := core.Get(pluginName, &storage)
	if err != nil {
		out.Error(err)
	}
	for k, v := range storage {
		out.Print(v, "=>", k)
	}
}
