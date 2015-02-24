// +build darwin

package homebrew

import (
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins/packages/homebrew/brew"
)

func backup() {
	casks, caskErr := brew.ListCasks()
	taps, tapsErr := brew.ListTaps()
	formulas, formErr := brew.List()

	out.Print(caskErr, tapsErr, formErr)

	p := plugin{Casks: casks, Taps: taps, Formulas: formulas}
	core.Add(pluginName, p)
	core.Save()
}
