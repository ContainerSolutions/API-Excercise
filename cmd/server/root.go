package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	environment  string
	postgresHost string
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "A REST API server that serves the titanic dataset.",
	Long:  `A REST API server that serves the titanic dataset.`,
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		setupLogger()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&environment, "env", "dev", "Environment to populate.")
	rootCmd.PersistentFlags().StringVar(&postgresHost, "dbHost", "localhost", "Hostname of the database.")
}

func setupLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
}
