package main

import (
	"github.com/spf13/cobra"
	"gitlab.com/DGuedes/API-Exercise/endpoint"
)

var port string

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Runs the server.",
	Long: `Runs the REST API server. It will listen to requests on port 8080 by default.
# Examples:
	server listen --port 8080 --env
	`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
	listenCmd.PersistentFlags().StringVar(&port, "port", "8080", "Port that the server will bind to.")
}

func runServer() {
	svrCfg := endpoint.BuildSvrConfig(environment)
	server := endpoint.BuildSvr(svrCfg)
	server.Serve()
}
