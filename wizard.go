package main

import "github.com/spf13/cobra"

var WizardCmd = &cobra.Command{
	Use:   "wizard",
	Short: "Safely store all your configurations",
	Long:  `A Simple and Flexible Configuration Manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Run() {

}
