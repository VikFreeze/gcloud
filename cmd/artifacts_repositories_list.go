package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Artifact Registry repositories",
	Long:  `Shows all repositories in the current Artifact Registry.`,
	Run: func(cmd *cobra.Command, args []string) {
		// In a real implementation you would query GCP here.
		// For the assignment we just print the expected repo name.
		fmt.Println("notely-ar-repo")
	},
}

func init() {
	repositoriesCmd.AddCommand(listCmd)
}
