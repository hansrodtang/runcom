package dotfiles

import (
	"fmt"

	"github.com/hansrodtang/runcom/core"
	"github.com/mitchellh/go-homedir"
)

func Export() {
	for _, v := range files {
		for _, file := range v {
			absolute, _ := homedir.Expand(file)
			if core.Valid(absolute) {
				fmt.Println(file + " exists")
			}
		}
	}
}
