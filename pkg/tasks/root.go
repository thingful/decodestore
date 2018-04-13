package tasks

import (
	"fmt"
	"os"

	"github.com/renstrom/dedent"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dcs",
	Short: "dcs is an implementation of an encrypted data store for DECODE",
	Long: dedent.Dedent(`An implementation of a simple encrypted data store created for DECODE.

  DECODE (https://decodeproject.eu) is an EU funded project that attempts to
  provide tools that put individuals in control of their personal
  information, allowing them to decide when and who to share it with.
		`),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dcs")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serverCmd)
}
