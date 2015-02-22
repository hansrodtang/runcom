// +build darwin

package homebrew

type plugin struct {
	Casks    []string `json:"casks"`
	Formulas []string `json:"formulas"`
	Taps     []string `json:"taps"`
}
