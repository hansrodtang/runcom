package homebrew

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Tap struct {
	identifier string
}

// http://pythonic.zoomquiet.io/data/20111223160257/index.html

func GetTaps() ([]Tap, error) {
	var taps []Tap

	brewCmd := exec.Command(BrewCmd, "tap")
	installedOut, err := brewCmd.CombinedOutput()
	if err != nil {
		return []Tap{}, err
	}
	installed := strings.Fields(string(installedOut))
	for _, v := range installed {
		if len(v) > 0 {
			taps = append(taps, Tap{v})
		}
	}

	return taps, nil
}

func SetTaps(taps ...Tap) {
	args := []string{"tap"}
	for _, v := range taps {
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
