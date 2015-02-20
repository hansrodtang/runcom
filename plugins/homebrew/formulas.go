// +build darwin

package homebrew

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// http://pythonic.zoomquiet.io/data/20111223160257/index.html

func GetFormulas() ([]string, error) {
	var formulas []string

	brewCmd := exec.Command(BrewCmd, "leaves")
	installedOut, err := brewCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	formulas = strings.Fields(string(installedOut))
	return formulas, nil
}

func Install(formulas []string) {
	args := []string{"install"}
	for _, v := range formulas {
		args = append(args, v)
	}

	brewCmd := exec.Command(BrewCmd, args...)
	brewCmd.Stdout = os.Stdout
	brewCmd.Stderr = os.Stderr
	brewCmd.Stdin = os.Stdin

	if err := brewCmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	}
	brewCmd.Wait()

}

//exec.LookPath("brew")
