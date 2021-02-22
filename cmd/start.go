package cmd

import (
	"hquzs/go-web/server"
	"hquzs/go-web/util"
	"hquzs/go-web/version"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	startViper = viper.New()
	log        = util.NewZeroLog("web", "cmd")
)

// rootCmd represents the base command when called without any subcommands
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start web",
}

func init() {
	// config file path
	startCmd.RunE = runStart
	startCmd.Flags().StringP("config", "c", "", "config file path (default is ./browser.yaml)")
	startViper.BindPFlag("config", startCmd.Flags().Lookup("config"))
	rootCmd.AddCommand(startCmd)
}

func getDefaultConfigFilePath() (string, error) {
	gopath := os.Getenv("GOPATH")
	cfgFilePath, err := util.MakeFileAbs("src/web/config.yaml", gopath)
	if err != nil {
		return "", err
	}
	return cfgFilePath, nil
}

func runStart(cmd *cobra.Command, args []string) error {
	log.Info(version.LogVersion())
	var err error
	cfgFile := startViper.GetString("config")
	if cfgFile == "" {
		cfgFile, err = getDefaultConfigFilePath()
		if err != nil {
			log.Fatal("Failed to get default config file path: ", err)
		}
		log.Warn("You are using the default config file ", cfgFile)
	}

	err = server.StartServer(cfgFile)
	if err != nil {
		log.Fatal("runStart WebServer failed: ", err)
	}
	return nil
}
