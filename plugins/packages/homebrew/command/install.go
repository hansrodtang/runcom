package command

import (
	"fmt"
	"os"
	"os/exec"
)

func Install(formulas []string) {
	install(brewCommand, "install", formulas)
}

func InstallTaps(taps []string) {
	install(brewCommand, "tap", taps)
}

func InstallCasks(casks []string) {
	install(caskCommand, "install", casks)
}

func install(command, args string, input []string) {
	for _, a := range input {

		cmd := exec.Command(command, args, a)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			fmt.Println(err)
		}

		cmd.Wait()
	}

}
