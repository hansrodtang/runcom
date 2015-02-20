package core

import "github.com/spf13/cobra"

var ConnectCommand = &cobra.Command{
	Use:   "connect",
	Short: "Connect to GitHub",
	Long:  `Filler`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var PushCommand = &cobra.Command{
	Use:   "push",
	Short: "Upload repository",
	Long:  `Filler`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var GetCommand = &cobra.Command{
	Use:   "get",
	Short: "Fetch repository",
	Long:  `Filler`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Add to plugins?
var UnpackCommand = &cobra.Command{
	Use:   "unpack",
	Short: "Open compressed storage",
	Long:  `Filler`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var PackCommand = &cobra.Command{
	Use:   "pack",
	Short: "Pack storage to compressed file",
	Long:  `Filler`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}
