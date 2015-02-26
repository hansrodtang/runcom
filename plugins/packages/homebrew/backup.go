// +build darwin

package homebrew

import (
	"github.com/hansrodtang/runcom/backends"
	"github.com/hansrodtang/runcom/core"
)

func backup() {
	storage := backends.Get()
	storage.Read()

	casks, caskErr := core.NewCommand(caskCommand, "list").List()
	taps, tapsErr := core.NewCommand(brewCommand, "tap").List()
	formulas, formErr := core.NewCommand(brewCommand, "leaves").List()

	out.Print(caskErr, tapsErr, formErr)

	p := plugin{Casks: casks, Taps: taps, Formulas: formulas}
	storage.Add(pluginName, p)
	storage.Save()
}
