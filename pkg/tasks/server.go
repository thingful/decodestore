package tasks

import (
	"github.com/spf13/cobra"

	"github.com/thingful/decodestore/pkg/server"
)

var bindAddr string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the data store server",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer(bindAddr)
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&bindAddr, "bind", "b", ":8080", "the listen address")
}
