package core

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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
		out.Print(filename + " is a symbolic link.")
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

type Printer struct {
	name    string
	errname string
}

func NewPrinter(name string) Printer {
	g := color.New(color.FgGreen).SprintfFunc()
	r := color.New(color.FgGreen).SprintfFunc()

	gn := g("[%s] ", name)
	rn := r("[%s] ", name)
	return Printer{gn, rn}
}

func (p Printer) Print(a ...interface{}) {
	fmt.Print(p.name)
	fmt.Println(a...)
}

func (p Printer) Printf(format string, a ...interface{}) {
	fmt.Print(p.name)
	fmt.Printf(format, a...)
}

func (p Printer) Error(a ...interface{}) {
	fmt.Print(p.errname)
	fmt.Println(a...)
}

func (p Printer) Write(b []byte) (int, error) {
	//reader, writer := io.Pipe()
	s := string(b)
	a := strings.Split(s, "\n")
	a = a[:len(a)-1]

	for _, v := range a {
		fmt.Println(p.name, v)
	}

	return 0, nil
}

func (p Printer) Ask(message string, value bool) bool {
	g := color.New(color.FgGreen).SprintFunc()
	r := color.New(color.FgRed).SprintFunc()

	if value {
		p.Print(message + g(" [Y/n]") + ": ")
	} else {
		p.Print(message + r(" [y/N]") + ": ")
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
