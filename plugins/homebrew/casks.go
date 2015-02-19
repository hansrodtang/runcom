package homebrew

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Cask struct {
	identifier string
}

func GetCasks() []Cask {
	var casks []Cask

	brewCmd := exec.Command(CaskCmd, "list")
	installedOut, _ := brewCmd.CombinedOutput()
	installed := strings.Fields(string(installedOut))
	for _, v := range installed {
		if len(v) > 0 {
			casks = append(casks, Cask{v})
		}
	}

	return casks
}

func InstallCasks(casks ...Cask) {
	args := []string{"install"}
	for _, v := range casks {
		args = append(args, v.identifier)
	}

	brewCmd := exec.Command(CaskCmd, args...)
	brewCmd.Stdout = os.Stdout
	brewCmd.Stderr = os.Stderr
	brewCmd.Stdin = os.Stdin

	if err := brewCmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	}
	brewCmd.Wait()

}

//exec.LookPath("brew")
