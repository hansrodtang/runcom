// +build !windows

package dotfiles

import (
	"github.com/hansrodtang/runcom/plugins"
	"github.com/spf13/cobra"
)

const PluginName = "dotfiles"

var Command = &cobra.Command{
	Use:   "dotfiles",
	Short: "Manage dotfiles",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var linkCommand = &cobra.Command{
	Use:   "link",
	Short: "Link a file to storage",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Linking stuff")
	},
}

var unlinkCommand = &cobra.Command{
	Use:   "unlink",
	Short: "Unlink a file from storage",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Unlinking stuff")
	},
}

var restoreCommand = &cobra.Command{
	Use:   "restore",
	Short: "Restore dotfiles from storage",
	Long:  `Restore dotfiles from storage`,
	Run: func(cmd *cobra.Command, args []string) {
		Restore()
	},
}

var backupCommand = &cobra.Command{
	Use:   "backup",
	Short: "Backup dotfiles to storage",
	Long:  `Backup dotfiles to storage`,
	Run: func(cmd *cobra.Command, args []string) {
		Backup()
	},
}

func init() {
	Command.AddCommand(linkCommand)
	Command.AddCommand(unlinkCommand)
	Command.AddCommand(restoreCommand)
	Command.AddCommand(backupCommand)

	plugins.Register(PluginName, Restore, Backup, Command)
}
