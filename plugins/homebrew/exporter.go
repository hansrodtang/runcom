// +build darwin

package homebrew

import (
	"fmt"

	"github.com/hansrodtang/runcom/core"
)

func Export() {
	casks, caskErr := GetCasks()
	taps, tapsErr := GetTaps()
	formulas, formErr := GetFormulas()

	fmt.Println(caskErr, tapsErr, formErr)

	p := plugin{Casks: casks, Taps: taps, Formulas: formulas}
	core.Add(PluginName, p)
	core.Save()
}
