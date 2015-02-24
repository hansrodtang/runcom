// +build darwin

package homebrew

import (
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins/packages/homebrew/brew"
)

func restore() {
	var p plugin
	err := core.Get(pluginName, &p)
	if err != nil {
		out.Print(err)
	}

	brew.InstallTaps(p.Taps...)
	brew.InstallCasks(p.Casks...)
	brew.Install(p.Formulas...)
}
