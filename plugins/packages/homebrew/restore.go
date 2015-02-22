// +build darwin

package homebrew

import (
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins/packages/homebrew/command"
)

func Restore() {
	var p plugin
	err := core.Get(PluginName, &p)
	if err != nil {
		print(err)
	}

	command.InstallTaps(p.Taps...)
	command.InstallCasks(p.Casks...)
	command.Install(p.Formulas...)
}
