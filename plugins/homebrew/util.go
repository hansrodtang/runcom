// +build darwin

package homebrew

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hansrodtang/runcom/core"
)

type plugin struct {
	Casks    []string `json:"casks"`
	Formulas []string `json:"formulas"`
	Taps     []string `json:"taps"`
}

func installBrew() bool {
	installCmd := exec.Command("ruby", "-e", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	installCmd.Stdin = os.Stdin

	if err := installCmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
		return false
	}
	installCmd.Wait()
	return true

}

func installCask() bool {
	if !core.IsInstalled(BrewCmd) {
		if !installBrew() {
			return false
		}
	}

	installCmd := exec.Command(BrewCmd, "install", "caskroom/cask/brew-cask")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	installCmd.Stdin = os.Stdin

	if err := installCmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
		return false
	}
	installCmd.Wait()
	return true

}
