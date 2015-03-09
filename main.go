package main

import (
	"fmt"
	"os"

	"github.com/hansrodtang/runcom/backends"
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins"
	// Backend
	_ "github.com/hansrodtang/runcom/backends/directory"
	_ "github.com/hansrodtang/runcom/backends/git"
	// Plugins
	_ "github.com/hansrodtang/runcom/plugins/dotfiles"
	_ "github.com/hansrodtang/runcom/plugins/packages/homebrew"
	_ "github.com/hansrodtang/runcom/plugins/scripts"
	// Other dependencies
	"github.com/hansrodtang/viper"
	"github.com/spf13/cobra"
)

const defaultBackend = "directory"

func main() {
	viper.SetDefault("backend", defaultBackend)

	var mainCmd = &cobra.Command{
		Use:   core.BinaryName,
		Short: core.Name + " is a simple and flexible configuration manager",
		Long: `A Simple and Flexible Configuration Manager..
Complete documentation is available at http:/runcom.github.io`,
	}

	mainCmd.PersistentFlags().String("backend", defaultBackend, "Storage backend to use for saving configuration")
	viper.BindPFlag("backend", mainCmd.PersistentFlags().Lookup("backend"))

	print := core.NewPrinter(core.BinaryName)

	backendErr := backends.Select(viper.GetString("backend"))
	if backendErr != nil {
		answer := print.Ask("Backend '"+viper.GetString("backend")+"' not found. Use default? ["+defaultBackend+"]", true)
		if !answer {
			fmt.Println("No storage backend chosen. Exiting")
			os.Exit(0)
		}
		backendErr = backends.Select(defaultBackend)
		if backendErr != nil {
			panic("Could not load default backend.")
		}

	}

	plugs := plugins.GetAll()

	for _, plugin := range plugs {
		mainCmd.AddCommand(plugin.Command)
	}

	mainCmd.AddCommand(
		core.ConnectCommand,
		core.PushCommand,
		core.GetCommand,
		core.UnpackCommand,
		core.PackCommand)

	mainCmd.Execute()
}
