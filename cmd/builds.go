package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsCmd = &cobra.Command{
	Use:   "builds",
	Short: "Build commands",
	Long:  `gcloud builds commands – perform build related actions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("No subcommand for builds. Try 'gcloud builds submit .'")
	},
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a build",
	Long:  `Submit the current directory as a build.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build submitted successfully.")
	},
}

func init() {
	buildsCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(buildsCmd)
}
