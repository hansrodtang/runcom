// +build darwin

package homebrew

import (
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
	taps = strings.Fields(string(installedOut))
	return taps, nil
}

func SetTaps(taps []string) {
	args := []string{"tap"}
	for _, v := range taps {
		args = append(args, v)
	}

	brewCmd := exec.Command(BrewCmd, args...)
	brewCmd.Stdout = os.Stdout
	brewCmd.Stderr = os.Stderr
	brewCmd.Stdin = os.Stdin

	if err := brewCmd.Start(); err != nil {
		print(err)
	}
	brewCmd.Wait()

}

//exec.LookPath("brew")
