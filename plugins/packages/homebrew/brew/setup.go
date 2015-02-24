package brew

import "github.com/hansrodtang/runcom/core"

const brewCommand = "brew"
const caskCommand = "brew-cask"

func Setup() bool {
	if !core.IsInstalled(brewCommand) {
		install("ruby", "-e", []string{"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"})
	}
	if !core.IsInstalled(caskCommand) {
		Install("caskroom/cask/brew-cask")
	}
	return true
}
