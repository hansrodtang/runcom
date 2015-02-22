package scripts

import (
	"github.com/hansrodtang/runcom/plugins"
	"github.com/spf13/cobra"
)

const (
	PluginName = "scripts"
)

var Command = &cobra.Command{
	Use:   "scripts",
	Short: "Store custom scripts",
	Long:  `Store custom scripts`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var AddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add script to storage",
	Long:  `Add script to storage`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var RemoveCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove script from storage",
	Long:  `Remove script from storage`,

	Run: func(cmd *cobra.Command, args []string) {
	},
}

var RunCommand = &cobra.Command{
	Use:   "run",
	Short: "Run scripts from storage",
	Long:  `Run scripts from storage`,

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Command.AddCommand(AddCommand)
	Command.AddCommand(RemoveCommand)
	Command.AddCommand(RunCommand)
	plugins.Register(PluginName, nil, nil, Command)
}
