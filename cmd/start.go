package cmd

import (
	"github.com/atallahhezbor/habit/habits"
	"github.com/atallahhezbor/habit/output"
	"github.com/spf13/cobra"
	"strings"
)

var shortName string
var tag string
var startCmd = &cobra.Command{
	Use:   "start [a habit to start doing]",
	Short: "Create a new habit",
	Long: `A habit is something you want keep doing to work
			towards a long-term goal`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		habitMap := habits.Load()
		name := strings.Join(args, " ")
		if len(shortName) == 0 {
			// TODO: handle collisions
			shortName = args[0]
		}
		habitMap[shortName] = habits.New(name, shortName, tag)
		habits.Save(habitMap)
		output.List(habitMap)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&shortName, "short", "", "Abbreviated name to refer to this habit")
	startCmd.Flags().StringVar(&tag, "tag", "", "Tag to give this habit to group it under a common goal")
	startCmd.MarkFlagRequired("tag")
}
