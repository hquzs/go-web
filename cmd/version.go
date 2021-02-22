package cmd

import (
	"fmt"
	"hquzs/go-web/version"

	"github.com/spf13/cobra"
)

var versionCMD = &cobra.Command{
	Use:   "version",
	Short: "print web version",
	RunE:  printVersion,
}

func init() {
	rootCmd.AddCommand(versionCMD)
}

func printVersion(cmd *cobra.Command, args []string) error {
	fmt.Println(version.LogVersion())
	return nil
}
