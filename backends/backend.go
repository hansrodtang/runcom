package backends

import "errors"

var backends map[string]Backend

func init() {
	backends = make(map[string]Backend)
}

var backend Backend

// Make a ignore file, to force files to not be used on a computer.

type Backend interface {
	Init() error                                  // Backend initialization (Create directory etc.)
	Read() error                                  // Read configuration
	Save() error                                  // Save configuration
	Link(string) error                            // Add item to storage
	Unlink(string) error                          // Remove item from storage, put back to original place
	Restore(string) error                         // Create link from storage to location. (Or copy, depends on backend).
	Remove()                                      // Remove item from machine but not storage
	Get(plugin string, content interface{}) error // Get plugin configuration
	Add(string, interface{}) error                // Add new plugin configuration
}

func Register(name string, be Backend) {
	backends[name] = be
}

func Get() Backend {
	return backend
}

func Select(name string) error {
	b, ok := backends[name]
	if !ok {
		return errors.New("no backend found")
	}
	backend = b
	return nil
}

func GetAll() map[string]Backend {
	return backends
}
