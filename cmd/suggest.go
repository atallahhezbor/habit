package cmd

import (
	"github.com/atallahhezbor/habit/habits"
	"github.com/atallahhezbor/habit/output"
	"github.com/spf13/cobra"
)

var suggestCommand = &cobra.Command{
	Use:   "suggest",
	Short: "Get a random suggestion from your habit list",
	Long: `Use this command when you want
	        one of your habits to be recommended to you
			randomly`,
	Run: func(cmd *cobra.Command, args []string) {
		habitMap := habits.Load()
		output.Suggest(habitMap)
	},
}

func init() {
	rootCmd.AddCommand(suggestCommand)
}
