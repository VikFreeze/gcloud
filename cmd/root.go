package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: "Yes this is totally the legitimate gcloud binary :)",
	Use:     "gcloud",
	Short:   "A tiny replacement for gcloud",
	Long:    `gcloud is a minimal CLI that mimics the gcloud command for teaching purposes.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gcloud: no subcommand provided. See 'gcloud --help'")
	},
}

// Execute is called from main.go to run the CLI
func Execute() {
	rootCmd.AddCommand(buildsCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
