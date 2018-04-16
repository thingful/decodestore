package tasks

import (
	"github.com/renstrom/dedent"
	"github.com/spf13/cobra"

	"github.com/thingful/decodestore/pkg/server"
)

var (
	addr string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the data store server",
	Long: dedent.Dedent(`Starts the data store server running listening on the specified interface.

  The server exposes an RPC interface using a framework called Twirp. This
  uses protobuf to define the interface of the server, then generates client
  and server stubs for implementing and interacting with the service.`),
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer(addr)
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:8080", "the interface on which the server listens")
}
