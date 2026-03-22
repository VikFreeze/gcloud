package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fake configuration – replace with real logic as needed
var (
	fakeProject = "my-project"
	fakeAccount = "user@example.com"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration commands",
	Long:  `gcloud config commands – read, write, or delete configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("No subcommand for config. Try 'gcloud config list'")
	},
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the current configuration",
	Long:  `Shows the active GCP project and account.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("project: %s\n", fakeProject)
		fmt.Printf("account: %s\n", fakeAccount)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configListCmd)
}
