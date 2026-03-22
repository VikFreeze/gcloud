package cmd

import (
	"github.com/spf13/cobra"
)

var artifactsCmd = &cobra.Command{
	Use:   "artifacts",
	Short: "Manage Artifact Registry resources",
	Long:  `Commands that operate on Artifact Registry.`,
}

func init() {
	rootCmd.AddCommand(artifactsCmd)
}
