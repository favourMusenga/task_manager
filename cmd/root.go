/*
Copyright © 2026 favour Musenga favourmusenga@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	categorycmd "github.com/favourmusenga/task-manager/cmd/categoryCmd"
	todocmd "github.com/favourmusenga/task-manager/cmd/todoCmd"
	"github.com/favourmusenga/task-manager/internals/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taskmar",
	Short: "Priority-based task manager for the terminal",

	Long: `TaskMar is a lightweight and efficient CLI task manager built for organizing todos with priorities directly from your terminal.

It helps you create, manage, and track tasks using a simple command-line workflow designed for speed and productivity. With support for priority levels, TaskMar makes it easy to focus on important work while keeping your daily tasks organized.

Whether you are managing personal goals, study plans, or development tasks, TaskMar provides a clean and fast terminal experience for staying productive.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task-manager.yaml)")

	// Add subcommands
	rootCmd.AddCommand(todocmd.NewCommand())
	rootCmd.AddCommand(categorycmd.NewCommand())

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Initialize the database
	err := db.InitDB("task_manager.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to initialize database:", err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".task-manager" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".task-manager")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
