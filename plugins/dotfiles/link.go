// +build !windows

package dotfiles

import (
	"os"
	"path/filepath"

	"github.com/hansrodtang/runcom/core"
)

func link(file string) {
	_, filename := filepath.Split(file)
	newfile := filepath.Join(core.Directory(), filename)

	if !core.IsSymbolic(file) {
		/*err :=*/ os.Rename(file, newfile)
		/*err2 :=*/ os.Symlink(newfile, file)
	}
	// Git Commit
}

func unlink(location string) {
	_, filename := filepath.Split(location)
	file := filepath.Join(core.Directory(), filename)

	if core.IsSymbolic(location) {
		/*err :=*/ os.Remove(location)
		/*err2 :=*/ os.Rename(file, location)
	}

	// Git Commit
}
