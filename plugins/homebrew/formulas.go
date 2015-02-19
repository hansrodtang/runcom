package homebrew

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Formula struct {
	identifier string
}

type Brew struct {
}

// http://pythonic.zoomquiet.io/data/20111223160257/index.html

func GetInstalled() ([]Formula, error) {
	var formulas []Formula

	brewCmd := exec.Command(BrewCmd, "leaves")
	installedOut, err := brewCmd.CombinedOutput()
	if err != nil {
		return []Formula{}, err
	}
	installed := strings.Fields(string(installedOut))
	for _, v := range installed {
		if len(v) > 0 {
			formulas = append(formulas, Formula{v})
		}
	}
	return formulas, nil
}

func Install(formulas ...Formula) {
	args := []string{"install"}
	for _, v := range formulas {
		args = append(args, v.identifier)
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
