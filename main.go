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
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	var MainCmd = &cobra.Command{
		Use:   core.Command,
		Short: core.Name + " is a simple and flexible configuration manager",
		Long: `A Simple and Flexible Configuration Manager..
Complete documentation is available at http:/dots.github.io`,
		/*Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},*/
	}

	plugs := plugins.GetAll()

	for _, plugin := range plugs {
		MainCmd.AddCommand(plugin.Command)
	}

	MainCmd.AddCommand(
		core.ConnectCommand,
		core.PushCommand,
		core.GetCommand,
		core.UnpackCommand,
		core.PackCommand)

	if !terminal.IsTerminal(int(os.Stdout.Fd())) {
		var test io.Writer
		test = bufio.NewWriter(test)
		MainCmd.SetOutput(test)
	}

	MainCmd.Execute()
}
