// +build darwin

package homebrew

import (
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins"
	"github.com/spf13/cobra"
)

const (
	BrewCmd    = "brew"
	CaskCmd    = "brew-cask"
	PluginName = "homebrew"
)

var print core.PrintFunc

func init() {
	print = core.Printer(PluginName)
}

var Command = &cobra.Command{
	Use:   "homebrew",
	Short: "Manages your Homebrew packages",
	Long: `Restore and export your Homebrew packages.
Will ask to install Homebrew if not already installed`,

	Run: func(cmd *cobra.Command, args []string) {
		if !core.IsInstalled(BrewCmd) {
			answer := core.Ask("Homebrew is not installed, install now?", true)
			if answer == true {
				if !installBrew() {
					print("Homebrew installation unsuccessful")
				}
			}
		} else {
			err := cmd.Usage()
			if err != nil {
				print(err)
			}
		}
	},
}

var restoreCommand = &cobra.Command{
	Use:   "restore",
	Short: "Restores packages from storage",
	Long:  `Restore your Homebrew packages from storage`,

	Run: func(cmd *cobra.Command, args []string) {
		if core.IsInstalled(BrewCmd) {
			Restore()
		} else {
			print("Homebrew is not installed. Nothing to import.")
		}
	},
}

var backupCommand = &cobra.Command{
	Use:   "backup",
	Short: "Backs up packages to storage",
	Long:  `Backs up your Homebrew packages from storage`,

	Run: func(cmd *cobra.Command, args []string) {
		if core.IsInstalled(BrewCmd) {
			Backup()
		} else {
			print("Homebrew is not installed. Nothing to import.")
		}
	},
}

func init() {
	Command.AddCommand(restoreCommand)
	Command.AddCommand(backupCommand)
	plugins.Register(PluginName, Restore, Backup, Command)
}
