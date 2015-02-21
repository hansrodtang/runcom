// +build darwin

package homebrew

import "github.com/hansrodtang/runcom/core"

func Backup() {
	casks, caskErr := GetCasks()
	taps, tapsErr := GetTaps()
	formulas, formErr := GetFormulas()

	print(caskErr, tapsErr, formErr)

	p := plugin{Casks: casks, Taps: taps, Formulas: formulas}
	core.Add(PluginName, p)
	core.Save()
}
