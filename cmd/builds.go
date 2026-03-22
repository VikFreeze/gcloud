package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildsCmd = &cobra.Command{
	Use:   "builds",
	Short: "Build related commands",
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a build",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build submitted successfully")
	},
}

func init() {
	buildsCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(buildsCmd)
}
