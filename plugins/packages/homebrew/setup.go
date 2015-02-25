package homebrew

import "github.com/hansrodtang/runcom/core"

const brewCommand = "brew"
const caskCommand = "brew-cask"

func setup() bool {
	if !core.IsInstalled(brewCommand) {
		ruby := core.NewCommand("ruby", "-e")
		ruby.Install("$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)")
	}
	if !core.IsInstalled(caskCommand) {
		brew := core.NewCommand(brewCommand, "install")
		brew.Install("caskroom/cask/brew-cask")
	}
	return true
}
