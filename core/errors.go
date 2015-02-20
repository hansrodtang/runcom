package core

type PluginNotFoundError struct {
	msg string
}

func (e PluginNotFoundError) Error() string {
	return e.msg
}
