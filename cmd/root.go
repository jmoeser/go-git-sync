/*
Copyright © 2020 Jordan Moeser <github@defestri.org>

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

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-git-sync",
	Short: "A simple application for syncing from Git via a custom command",
	Long: `To sync the file example/consul/sample-json.json from the repo https://github.com/jmoeser/go-git-sync.git
to the Consul server at address 127.0.0.1:8500:

go-git-sync -c 127.0.0.1:8500 sync -s https://github.com/jmoeser/go-git-sync.git -f example/consul/sample-json.json
`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		if viper.GetBool("debug") {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			log.Debug().Msg("Debug logging enabled")
		}

	},

	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			os.Exit(0)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Set logging to debug level")
	err := viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().StringP("consul", "c", "", "Consul server address")
	err = viper.BindPFlag("consul", rootCmd.PersistentFlags().Lookup("consul"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().IntP("port", "p", 8080, "Listen on given port")
	err = viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.PersistentFlags().StringP("webhook-secret", "w", "", "Webhook Secret to validate webhook requests from Github")
	err = viper.BindPFlag("webhook-secret", rootCmd.PersistentFlags().Lookup("webhook-secret"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	viper.SetEnvPrefix("GGS")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
