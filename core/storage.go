package core

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

type storage struct {
	Plugins map[string]*json.RawMessage `json:"plugins"`
}

var store storage

func init() {
	// TODO: Better reading
	err := Read()
	if err != nil {
		p := make(map[string]*json.RawMessage)
		store = storage{Plugins: p}
	}
}

// Add marshals and stores plugin data
func Add(plugin string, content interface{}) error {

	b, err := json.Marshal(content)
	if err != nil {
		return err
	}
	store.Plugins[plugin] = (*json.RawMessage)(&b)
	return nil
}

// Get fetches and unmarshals plugin data
func Get(plugin string, content interface{}) error {
	data, ok := store.Plugins[plugin]
	if !ok {
		return PluginNotFoundError{plugin + " not found"}
	}
	parseErr := json.Unmarshal(([]byte)(*data), &content)

	if parseErr != nil {
		return parseErr
	}
	return nil
}

// Read fetches plugin data from file
func Read() error {
	dat, readErr := ioutil.ReadFile(viper.GetString("storage"))
	if readErr != nil {
		return readErr
	}
	parseErr := json.Unmarshal(dat, &store)
	if parseErr != nil {
		return parseErr
	}
	return nil
}

// Save stores plugin data to file
func Save() error {

	b, err := json.MarshalIndent(store, " ", "	")
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
