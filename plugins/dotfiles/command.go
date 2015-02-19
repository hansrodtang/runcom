package dotfiles

import (
	"github.com/hansrodtang/runcom/plugins"
	"github.com/spf13/cobra"
)

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

var importCommand = &cobra.Command{
	Use:   "import",
	Short: "Import from storage",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Importing stuff")
	},
}

var exportCommand = &cobra.Command{
	Use:   "export",
	Short: "Export to storage",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Exporting stuff")
	},
}

func init() {
	Command.AddCommand(linkCommand)
	Command.AddCommand(unlinkCommand)
	Command.AddCommand(importCommand)
	Command.AddCommand(exportCommand)

	plugins.Register("dotfiles", Import, Export, Command)
}
