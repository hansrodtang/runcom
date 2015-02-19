package core

import "os"

func init() {
	path := os.Getenv("DOT_PATH")
	if len(path) > 0 {
		if !Exists(path) {
			answer := Ask("Folder "+path+" does not exist. Use default?", true)
			if answer {
				directory = path
				CreateDir(Directory())
				return
			}
		} else {
			directory = path
			return
		}
	}
	if !Exists(Directory()) {
		answer := Ask("Folder "+directory+" does not exist. Create?", true)
		if answer {
			CreateDir(Directory())
			directory = path
			return
		}
	}
}
