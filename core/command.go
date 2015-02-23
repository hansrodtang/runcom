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

var BackupCommand = &cobra.Command{
	Use:   "backup",
	Short: "Backup configuration to storage",
	Long:  `Filler`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

var RestoreCommand = &cobra.Command{
	Use:   "restore",
	Short: "Restore configuration from storage",
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

var PruneCommand = &cobra.Command{
	Use:   "prune",
	Short: "Remove stuff not in storage",
	Long:  `Filler`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}
