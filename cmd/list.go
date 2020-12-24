package cmd

import (
	"github.com/atallahhezbor/habit/habits"
	"github.com/atallahhezbor/habit/output"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all habits that have been created",
	Long:  `fill in`,
	Run: func(cmd *cobra.Command, args []string) {
		habitMap := habits.Load()
		output.List(habitMap)
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
