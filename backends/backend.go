package backend

import "errors"

func init() {
	backends = make(map[string]Backend)
}

type ReadFunc func()
type SaveFunc func()

type Backend struct {
	Name string
	Read ReadFunc
	Save SaveFunc
}

var backends map[string]Backend

func Register(name string, read ReadFunc, save SaveFunc) {
	backends[name] = Backend{name, read, save}
}

func Get(name string) (Backend, error) {
	b, ok := backends[name]
	if !ok {
		return Backend{}, errors.New("no backend found")
	}
	return b, nil
}

func GetAll() map[string]Backend {
	return backends
}
