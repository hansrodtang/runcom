package command

import (
	"os/exec"
	"strings"
)

func List() ([]string, error) {
	return list(brewCommand, "leaves")
}

func ListTaps() ([]string, error) {
	return list(brewCommand, "tap")
}

func ListCasks() ([]string, error) {
	return list(caskCommand, "list")
}

func list(command, argument string) ([]string, error) {
	cmd := exec.Command(command, argument)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(out)), nil
}
