package scripts

import (
	"github.com/hansrodtang/runcom/plugins"
	"github.com/spf13/cobra"
)

const (
	pluginName = "scripts"
)

var command = &cobra.Command{
	Use:   "scripts",
	Short: "Store custom scripts",
	Long:  `Store custom scripts`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add script to storage",
	Long:  `Add script to storage`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var removeCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove script from storage",
	Long:  `Remove script from storage`,

	Run: func(cmd *cobra.Command, args []string) {
	},
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run scripts from storage",
	Long:  `Run scripts from storage`,

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	command.AddCommand(addCommand)
	command.AddCommand(removeCommand)
	command.AddCommand(runCommand)
	plugins.Register(pluginName, nil, nil, command)
}
