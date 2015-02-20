// +build darwin

package homebrew

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetCasks() ([]string, error) {
	var casks []string

	brewCmd := exec.Command(CaskCmd, "list")
	installedOut, err := brewCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	installed := strings.Fields(string(installedOut))
	for _, v := range installed {
		if len(v) > 0 {
			casks = append(casks, v)
		}
	}

	return casks, nil
}

func InstallCasks(casks ...string) {
	args := []string{"install"}
	for _, v := range casks {
		args = append(args, v)
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
