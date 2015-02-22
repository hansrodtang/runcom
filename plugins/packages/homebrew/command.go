// +build darwin

package homebrew

import (
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins"
	"github.com/hansrodtang/runcom/plugins/packages/homebrew/command"
	"github.com/spf13/cobra"
)

const (
	BrewCmd    = "brew"
	CaskCmd    = "brew-cask"
	PluginName = "homebrew"
)

var out core.Printer

func init() {
	out = core.NewPrinter(PluginName)
}

var Command = &cobra.Command{
	Use:   "homebrew",
	Short: "Manages your Homebrew packages",
	Long: `Restore and export your Homebrew packages.
Will ask to install Homebrew if not already installed`,

	Run: func(cmd *cobra.Command, args []string) {
		if !core.IsInstalled(BrewCmd) {
			answer := out.Ask("Homebrew is not installed, install now?", true)
			if answer == true {
				command.Setup()
			}
		} else {
			cmd.Usage()
		}
	},
}

var restoreCommand = &cobra.Command{
	Use:   "restore",
	Short: "Restores packages from storage",
	Long:  `Restore your Homebrew packages from storage`,

	Run: func(cmd *cobra.Command, args []string) {
		command.Setup()
		Restore()
	},
}

var backupCommand = &cobra.Command{
	Use:   "backup",
	Short: "Backs up packages to storage",
	Long:  `Backs up your Homebrew packages from storage`,

	Run: func(cmd *cobra.Command, args []string) {
		command.Setup()
		Backup()
	},
}

	},
}

func init() {
	Command.AddCommand(restoreCommand)
	Command.AddCommand(backupCommand)
	plugins.Register(PluginName, Restore, Backup, Command)
}
