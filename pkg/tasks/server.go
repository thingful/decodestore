package tasks

import (
	"github.com/spf13/cobra"

	"github.com/thingful/decodestore/pkg/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the data store server",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer(":8080")
		server.Start()
	},
}
