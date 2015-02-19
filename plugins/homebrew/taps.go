package homebrew

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// http://pythonic.zoomquiet.io/data/20111223160257/index.html

func GetTaps() ([]string, error) {
	var taps []string

	brewCmd := exec.Command(BrewCmd, "tap")
	installedOut, err := brewCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	installed := strings.Fields(string(installedOut))
	for _, v := range installed {
		if len(v) > 0 {
			taps = append(taps, v)
		}
	}

	return taps, nil
}

func SetTaps(taps ...string) {
	args := []string{"tap"}
	for _, v := range taps {
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
