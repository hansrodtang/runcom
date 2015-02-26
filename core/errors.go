package core

type PluginNotFoundError string

func (e PluginNotFoundError) Error() string {
	return string(e)
}
