package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const configFileNameDefault = ".habit"
const dataFileNameDefault = ".habit_data"

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "habit",
		Short: "Habit is a way of tracking daily progress",
		Long: `Track your long term goals and "dailies"
			   and visualize your progress towards them`,
	}
)

// Execute starts the show!
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		er(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.habit.json)")
	// TODO: load/touch data file here?
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}
		// Search config in home directory with name ".habit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(configFileNameDefault)
	}
	viper.AutomaticEnv()
	viper.SetDefault("datafile", dataFileNameDefault)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

func er(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
