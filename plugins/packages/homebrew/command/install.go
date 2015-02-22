package command

import (
	"fmt"
	"os"
	"io"
	"os/exec"

	"github.com/hansrodtang/runcom/core"
	"github.com/kr/pty"
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
		out := core.NewPrinter("homebrew")

		f, err := pty.Start(cmd)
		if err != nil {
			panic(err)
		}
		io.Copy(out, f)

	}

}
