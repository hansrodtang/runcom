package directory

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hansrodtang/runcom/backends"
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/viper"
)

func init() {
	n := NewBackend(core.Directory())
	backends.Register("directory", n)
}

type Directory struct {
	location string
	config   *Configuration
}

type Configuration struct {
	Plugins map[string]*json.RawMessage `json:"plugins"`
}

func NewBackend(location string) backends.Backend {
	p := make(map[string]*json.RawMessage)
	config := &Configuration{Plugins: p}
	return &Directory{location, config}
}

func (d *Directory) Init() error {
	err := d.Read()
	if err != nil {
		p := make(map[string]*json.RawMessage)
		d.config = &Configuration{Plugins: p}
	}

	if !core.Exists(d.location) {
		if !createDir(d.location) {
			return errors.New("could not create directory")
		}
	}
	return nil
}

func (d *Directory) Read() error {
	dat, readErr := ioutil.ReadFile(viper.GetString("storage"))
	if readErr != nil {
		return readErr
	}

	parseErr := json.Unmarshal(dat, &d.config)
	if parseErr != nil {
		return parseErr
	}
	return nil
}

func (d *Directory) Save() error {

	b, err := json.MarshalIndent(d.config, " ", "	")
	if err != nil {
		return err
	}

	f, err := os.Create(viper.GetString("storage"))
	if err != nil {
		return err
	}

	defer f.Close()
	f.WriteString(string(b))

	return nil
}

func (d *Directory) Add(plugin string, content interface{}) error {
	b, err := json.Marshal(content)
	if err != nil {
		return err
	}
	d.config.Plugins[plugin] = (*json.RawMessage)(&b)
	return nil
}

func (d *Directory) Get(plugin string, content interface{}) error {
	data, ok := d.config.Plugins[plugin]
	if !ok {
		return core.PluginNotFoundError("plugin " + plugin + " not found")
	}
	parseErr := json.Unmarshal(([]byte)(*data), &content)

	if parseErr != nil {
		return parseErr
	}
	return nil
}
func (d *Directory) path() string {
	return d.location

}

func (d *Directory) Link(file string) error {
	_, filename := filepath.Split(file)
	newfile := filepath.Join(d.path(), filename)

	if !core.IsSymbolic(file) {
		renameErr := os.Rename(file, newfile)
		if renameErr != nil {
			return renameErr
		}
		symlinkErr := os.Symlink(newfile, file)
		if symlinkErr != nil {
			return symlinkErr
		}
		return nil
	}
	return errors.New("file is already a symbolic link")
}

// Restore moves the supplied file from storage.
func (d *Directory) Unlink(location string) error {
	_, filename := filepath.Split(location)
	file := filepath.Join(d.path(), filename)

	if core.IsSymbolic(location) {
		removeErr := os.Remove(location)
		if removeErr != nil {
			return removeErr
		}
		renameErr := os.Rename(file, location)
		if renameErr != nil {
			return renameErr
		}
		return nil
	}
	return errors.New("file is not a symbolic link")
}

func (d *Directory) Restore(string) error {
	return nil

}

func (d *Directory) Remove() {

}

func createDir(directory string) bool {
	err := os.Mkdir(directory, 0755)
	if err != nil {
		return false
	}
	return true
}
