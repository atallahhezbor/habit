package cmd

import (
	"github.com/atallahhezbor/habit/habits"
	"github.com/atallahhezbor/habit/output"
	"github.com/spf13/cobra"
)

var histCmd = &cobra.Command{
	Use:   "hist",
	Short: "Visualize habit progress",
	Long:  "Display a histogram of progress towards habits",
	Run: func(cmd *cobra.Command, args []string) {
		habitMap := habits.Load()
		output.Hist(habitMap)
	},
}

func init() {
	rootCmd.AddCommand(histCmd)
	// TODO: flag for category?
}
