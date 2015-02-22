// +build darwin

package homebrew

import (
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins/packages/homebrew/command"
)

func Backup() {
	casks, caskErr := command.ListCasks()
	taps, tapsErr := command.ListTaps()
	formulas, formErr := command.List()

	out.Print(caskErr, tapsErr, formErr)

	p := plugin{Casks: casks, Taps: taps, Formulas: formulas}
	core.Add(PluginName, p)
	core.Save()
}
