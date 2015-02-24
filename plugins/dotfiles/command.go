// +build !windows

package dotfiles

import (
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins"
	"github.com/spf13/cobra"
)

const pluginName = "dotfiles"

var out core.Printer

func init() {
	out = core.NewPrinter(pluginName)
}

var command = &cobra.Command{
	Use:   "dotfiles",
	Short: "Manage dotfiles",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		out.Print("Default command")
	},
}

var linkCommand = &cobra.Command{
	Use:   "link",
	Short: "Link a file to storage",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		out.Print("Linking stuff")
	},
}

var unlinkCommand = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink a file from storage",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		out.Print("Unlinking stuff")
	},
}

var restoreCommand = &cobra.Command{
	Use:   "restore",
	Short: "Restore dotfiles from storage",
	Long:  `Restore dotfiles from storage`,
	Run: func(cmd *cobra.Command, args []string) {
		restore()
	},
}

var backupCommand = &cobra.Command{
	Use:   "backup",
	Short: "Backup dotfiles to storage",
	Long:  `Backup dotfiles to storage`,
	Run: func(cmd *cobra.Command, args []string) {
		backup()
	},
}

func init() {
	command.AddCommand(linkCommand)
	command.AddCommand(unlinkCommand)
	command.AddCommand(restoreCommand)
	command.AddCommand(backupCommand)

	plugins.Register(pluginName, restore, backup, command)
}
