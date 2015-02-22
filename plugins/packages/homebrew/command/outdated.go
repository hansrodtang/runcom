package command

import (
	"fmt"
	"os/exec"
	"regexp"
)

func Outdated() []string {

	re := regexp.MustCompile(`^(d.*) \(\d.*\)$`)

	cmd := exec.Command(brewCommand, "outdated")
	output, err := cmd.Output()
	if err != nil {
		print(err)
	}
	s := re.FindSubmatch(output)
	for _, v := range s {
		fmt.Println(v)
	}
	return nil

}
