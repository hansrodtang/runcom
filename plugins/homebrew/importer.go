// +build darwin

package homebrew

import (
	"fmt"

	"github.com/hansrodtang/runcom/core"
)

func Import() {
	var p plugin
	err := core.Get(PluginName, &p)
	if err != nil {
		fmt.Println(err)
	}

	SetTaps(p.Taps)
	InstallCasks(p.Casks)
	Install(p.Formulas)
}
