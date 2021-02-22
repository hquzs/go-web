package cmd

import (
	"fmt"
	"hquzs/go-web/config"
	"hquzs/go-web/util"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	initCmd = &cobra.Command{
		Use: "init",
	}
	initViper    = viper.New()
	workerDir, _ = os.Getwd()
)

func init() {
	initCmd.RunE = runinit
	initCmd.Flags().StringP("config", "c", "config.yaml", "The config file of web")
	initViper.BindPFlag("config", initCmd.Flags().Lookup("config"))
	rootCmd.AddCommand(initCmd)
}

func runinit(cmd *cobra.Command, args []string) error {
	//log.Info(version.LogVersion())
	cfgFile := initViper.GetString("config")
	if cfgFile == "" {
		cfgFile = "config.yaml"
	}
	err := createConfigFile(cfgFile)
	if err != nil {
		log.Error("Create config file failed: ", err)
		return err
	}
	log.Infof("Create config file %s succeed.", cfgFile)
	return nil
}

func createConfigFile(cfgFile string) error {
	log.Info("dir:", workerDir)
	cfgAbsPath, err := util.MakeFileAbs(cfgFile, workerDir)
	if err != nil {
		return err
	}

	if util.FileExists(cfgAbsPath) {
		return fmt.Errorf("Config file %s exist, please delete it first", cfgAbsPath)
	}

	cfg := config.Config{
		ConnectionLimit: config.DefaultConnections,
		LogLevel:        "Info",
		HTTPConfig: config.HTTPConfig{
			Port: 9090,
			Host: "localhost",
		},
	}
	d, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatal("Failed to marshal yaml config: ", err)
	}
	return ioutil.WriteFile(cfgAbsPath, []byte(d), 0755)
}
