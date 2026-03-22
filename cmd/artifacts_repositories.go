package cmd

import (
	"github.com/spf13/cobra"
)

var repositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "Manage Artifact Registry repositories",
	Long:  `Commands that operate on Artifact Registry repositories.`,
}

func init() {
	artifactsCmd.AddCommand(repositoriesCmd)
}
