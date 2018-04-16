package tasks

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/thingful/decodestore/pkg/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version of dcs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s. Build Date: %s\n", version.Version, version.BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
