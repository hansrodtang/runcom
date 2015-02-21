// +build !windows

package dotfiles

import "github.com/hansrodtang/runcom/core"

func Restore() {
	err := core.Get(PluginName, &symlink)
	if err != nil {
		print(err)
	}
	for k, v := range symlink {
		print(v, "=>", k)
	}
}
