// +build !windows

package core

import (
	"os"
	"path/filepath"
)

// Link moves the supplied file to storage and creates a symbolic
// link to it.
func Link(file string) {
	_, filename := filepath.Split(file)
	newfile := filepath.Join(Directory(), filename)

	if !IsSymbolic(file) {
		/*err :=*/ os.Rename(file, newfile)
		/*err2 :=*/ os.Symlink(newfile, file)
	}
}

// Unlink moves the supplied file from storage.
func Unlink(location string) {
	_, filename := filepath.Split(location)
	file := filepath.Join(Directory(), filename)

	if IsSymbolic(location) {
		/*err :=*/ os.Remove(location)
		/*err2 :=*/ os.Rename(file, location)
	}

	// Git Commit
}
