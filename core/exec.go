package core

import (
	"io"
	"os/exec"
	"strings"

	"github.com/kr/pty"
)

type Command struct {
	command string
	args    []string
}

func NewCommand(command string, args ...string) Command {
	return Command{command, args}
}

// Install is a command executer for install commands.
func (c Command) Install(input ...string) {

	for _, a := range input {
		args := append(c.args, a)
		cmd := exec.Command(c.command, args...)
		out := NewPrinter("homebrew")

		f, err := pty.Start(cmd)
		if err != nil {
			panic(err)
		}
		io.Copy(out, f)

	}

}

// List is a command executer for collecting command output.
func (c Command) List() ([]string, error) {
	cmd := exec.Command(c.command, c.args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(out)), nil
}

func IsInstalled(command string) bool {
	_, err := exec.LookPath(command)
	if err != nil {
		return false
	}
	return true
}
