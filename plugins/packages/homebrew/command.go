// +build darwin

package homebrew

import (
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins"
	"github.com/hansrodtang/runcom/plugins/packages/homebrew/brew"
	"github.com/spf13/cobra"
)

const (
	brewCmd    = "brew"
	caskCmd    = "brew-cask"
	pluginName = "homebrew"
)

var out core.Printer

func init() {
	out = core.NewPrinter(pluginName)
}

var command = &cobra.Command{
	Use:   "homebrew",
	Short: "Manages your Homebrew packages",
	Long: `Restore and export your Homebrew packages.
Will ask to install Homebrew if not already installed`,

	Run: func(cmd *cobra.Command, args []string) {
		if !core.IsInstalled(brewCmd) {
			answer := out.Ask("Homebrew is not installed, install now?", true)
			if answer == true {
				brew.Setup()
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
		brew.Setup()
		restore()
	},
}

var backupCommand = &cobra.Command{
	Use:   "backup",
	Short: "Backs up packages to storage",
	Long:  `Backs up your Homebrew packages from storage`,

	Run: func(cmd *cobra.Command, args []string) {
		brew.Setup()
		backup()
	},
}

func init() {
	command.AddCommand(restoreCommand)
	command.AddCommand(backupCommand)
	plugins.Register(pluginName, restore, backup, command)
}
