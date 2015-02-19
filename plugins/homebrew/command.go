package homebrew

import (
	"github.com/hansrodtang/runcom/core"
	"github.com/hansrodtang/runcom/plugins"
	"github.com/spf13/cobra"
)

const (
	BrewCmd = "brew"
	CaskCmd = "brew-cask"
)

var Command = &cobra.Command{
	Use:   "homebrew",
	Short: "Manages your Homebrew packages",
	Long: `Import and export your Homebrew packages.
Will ask to install Homebrew if not already installed`,

	Run: func(cmd *cobra.Command, args []string) {
		if !core.IsInstalled(BrewCmd) {
			answer := core.Ask("Homebrew is not installed, install now?", true)
			if answer == true {
				if !installBrew() {
					cmd.Println("Homebrew installation unsuccessful")
				}
			}
		} else {
			cmd.Usage()
		}
	},
}

var ImportCommand = &cobra.Command{
	Use:   "import",
	Short: "Imports packages from storage",
	Long:  `Import your Homebrew packages from storage`,

	Run: func(cmd *cobra.Command, args []string) {
		if core.IsInstalled(BrewCmd) {
			InstallCasks(GetCasks()...)
		} else {
			cmd.Println("Homebrew is not installed. Nothing to import.")
		}
	},
}

var ExportCommand = &cobra.Command{
	Use:   "import",
	Short: "Imports packages from storage",
	Long:  `Import your Homebrew packages from storage`,

	Run: func(cmd *cobra.Command, args []string) {
		if core.IsInstalled(BrewCmd) {
			GetCasks()
		} else {
			cmd.Println("Homebrew is not installed. Nothing to import.")
		}
	},
}

func init() {
	Command.AddCommand(ImportCommand)
	plugins.Register("homebrew", Import, Export, Command)
}
