// +build darwin

package homebrew

import "github.com/hansrodtang/runcom/core"

func Restore() {
	var p plugin
	err := core.Get(PluginName, &p)
	if err != nil {
		print(err)
	}

	SetTaps(p.Taps)
	InstallCasks(p.Casks)
	Install(p.Formulas)
}
