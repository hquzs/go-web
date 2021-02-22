package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "webServer",
	Short: "webServer, restful api used to hand business",
	Long: `WebServer is a api server.
You can use get/post method to get Genkey or to verify issuing/transfer tx.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error("RootCmd exec failed: ", err)
		os.Exit(1)
	}
}
