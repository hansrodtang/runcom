package core

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func Valid(filename string) bool {
	if Exists(filename) {
		if IsSymbolic(filename) {
			return false
		}
		return true
	}
	return false
}

func IsSymbolic(filename string) bool {
	fi, err := os.Lstat(filename)
	if err != nil {
		return true
	}
	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		print(filename + " is a symbolic link.")
		return true
	} else {
		return false
	}
}

func Exists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(directory string) bool {
	err := os.Mkdir(directory, 0755)
	if err != nil {
		return false
	}

	return true
}

func Ask(message string, value bool) bool {
	g := color.New(color.FgGreen).SprintFunc()
	r := color.New(color.FgRed).SprintFunc()

	if value {
		fmt.Print(message + g(" [Y/n]") + ": ")
	} else {
		fmt.Print(message + r(" [y/N]") + ": ")
	}

	var response string
	fmt.Scanln(&response)

	yesResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	noResponses := []string{"n", "N", "no", "No", "NO"}
	if contains(yesResponses, response) {
		return true
	} else if contains(noResponses, response) {
		return false
	} else {
		return value
	}
}

func IsInstalled(command string) bool {
	_, err := exec.LookPath(command)
	if err != nil {
		return false
	}
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type PrintFunc func(...interface{})

func Printer(plugin string) PrintFunc {
	g := color.New(color.FgGreen).SprintfFunc()
	p := g("[%s] ", plugin)

	print := func(a ...interface{}) {
		fmt.Print(p)
		fmt.Println(a...)
	}
	return print
}
