package cmd

import (
	"fmt"
	"github.com/atallahhezbor/habit/habits"
	"github.com/atallahhezbor/habit/output"
	"github.com/spf13/cobra"
)

var tickCmd = &cobra.Command{
	Use:   "tick [short-name of a habit]",
	Short: "Tick an existing habit",
	Long:  "Mark Progress towards a habit",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		habitMap := habits.Load()
		if habitMap[name] == nil {
			return fmt.Errorf("Invalid short name %s. Use `list` to view habit short names", name)
		}
		habitMap[name].Tick()
		habits.Save(habitMap)
		output.Hist(habitMap)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tickCmd)
}
