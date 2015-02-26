// +build darwin

package homebrew

import (
	"github.com/hansrodtang/runcom/backends"
	"github.com/hansrodtang/runcom/core"
)

func restore() {
	storage := backends.Get()
	storage.Read()

	var p plugin
	err := storage.Get(pluginName, &p)
	if err != nil {
		out.Print(err)
	}

	tapsCmd := core.NewCommand(brewCommand, "tap")
	tapsCmd.Install(p.Taps...)

	casksCmd := core.NewCommand(caskCommand, "install")
	casksCmd.Install(p.Casks...)

	formulasCmd := core.NewCommand(brewCommand, "install")
	formulasCmd.Install(p.Formulas...)

}
