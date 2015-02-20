// +build !windows

package dotfiles

import (
	"fmt"

	"github.com/hansrodtang/runcom/core"
)

func Import() {
	err := core.Get(PluginName, &symlink)
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range symlink {
		fmt.Println(k, v)
	}
}
