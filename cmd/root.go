/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

const (
	cfgName = "config"
	cfgType = "yaml"
)

var (
	//Logger xlog.Logger
	cfgPath string
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "tasker",
	Long:  `tasker`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initialize)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgPath, "config", "", "config file (default is $HOME/.integration_vok_callback.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgPath != "" {
		viper.SetConfigFile(cfgPath) // path to look for the config file in
	} else {
		viper.SetConfigName(cfgName) // name of config file (without extension)
		viper.SetConfigType(cfgType) // REQUIRED if the config file does not have the extension in the name
		//viper.AddConfigPath(fmt.Sprintf("/etc/%s/", appName)) // path to look for the config file in
		viper.AddConfigPath(".") // optionally look for config in the working directory
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()        // read in environment variables that match
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

//func initLogger() {
//	Logger = xlogrus.NewXLogrus(xlog.LoggerCfg{
//		Level: "trace",
//		Out:   os.Stdout,
//	})
//}

func initialize() {
	//initConfig()
	//initLogger()
}
