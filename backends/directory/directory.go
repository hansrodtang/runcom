package directory

import "os"

func Read() {

}

func Save() {

}

func CreateDir(directory string) bool {
	err := os.Mkdir(directory, 0755)
	if err != nil {
		return false
	}

	return true
}
