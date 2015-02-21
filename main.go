package main

import (
	"bufio"
	"io"
	"os"

	// Plugins
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins"
	_ "github.com/hansrodtang/runcom/plugins/dotfiles"
	_ "github.com/hansrodtang/runcom/plugins/packages/homebrew"
	// Other dependencies
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

const defaultBackend = "directory"

func main() {
	viper.SetDefault("backend", defaultBackend)

	var mainCmd = &cobra.Command{
		Use:   core.Command,
		Short: core.Name + " is a simple and flexible configuration manager",
		Long: `A Simple and Flexible Configuration Manager..
Complete documentation is available at http:/dots.github.io`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here

		},
	}

	mainCmd.PersistentFlags().String("backend", defaultBackend, "Storage backend to use for saving configuration")
	viper.BindPFlag("backend", mainCmd.PersistentFlags().Lookup("backend"))

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

	if !terminal.IsTerminal(int(os.Stdout.Fd())) {
		var test io.Writer
		test = bufio.NewWriter(test)
		mainCmd.SetOutput(test)
	}

	mainCmd.Execute()
}
