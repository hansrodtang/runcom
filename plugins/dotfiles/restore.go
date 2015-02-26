// +build !windows

package dotfiles

import "github.com/hansrodtang/runcom/backends"

func restore() {
	storage := backends.Get()
	storage.Read()

	err := storage.Get(pluginName, &model)
	if err != nil {
		out.Error(err)
	}
	for k, v := range model {
		out.Print(v, "=>", k)
	}
}
