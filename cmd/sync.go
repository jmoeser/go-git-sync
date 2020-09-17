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

	"github.com/jmoeser/go-git-sync/api"
	"github.com/jmoeser/go-git-sync/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync changes from Git using the specified command",
	Long:  `Sync changes from specified Git source repo using the command specified`,
	Run: func(cmd *cobra.Command, args []string) {

		consulServer := viper.GetString("consul")
		if consulServer == "" {
			log.Error().Msg("Consul server address missing")

			err := rootCmd.Help()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			os.Exit(1)
		}

		log.Debug().Msgf("Consul server: %s", consulServer)

		httpPort := viper.GetInt("port")
		log.Debug().Msgf("Starting on port %d", httpPort)

		serverOpts := server.GoGitSyncServerOptions{
			ConsulHost: consulServer,
		}

		source := viper.GetString("source")
		if source == "" {
			log.Error().Msg("Source repo missing")

			err := rootCmd.Help()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			os.Exit(1)
		}

		filePath := viper.GetString("file")
		if filePath == "" {
			log.Error().Msg("File path in source repo missing")

			err := rootCmd.Help()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			os.Exit(1)
		}

		destinationPrefix := viper.GetString("prefix")
		revision := viper.GetString("revision")

		if viper.GetString("webhook-secret") == "" {
			log.Info().Msg("Didn't get a webhook secret, starting polling server")
			go api.StartSyncLoop(source, filePath, consulServer, destinationPrefix, revision)
		}

		goGitSyncServer := server.NewServer(serverOpts)
		goGitSyncServer.Run(httpPort, 8181)

	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	defaultRevision := "master"

	syncCmd.Flags().StringP("source", "s", "", "Source Git URL")
	err := viper.BindPFlag("source", syncCmd.Flags().Lookup("source"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	syncCmd.Flags().StringP("file", "f", "", "File path in the Git repo")
	err = viper.BindPFlag("file", syncCmd.Flags().Lookup("file"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	syncCmd.Flags().StringP("prefix", "", "", "Prefix of the path to sync to in Consul")
	err = viper.BindPFlag("prefix", syncCmd.Flags().Lookup("prefix"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	syncCmd.Flags().StringP("revision", "r", defaultRevision, "Revision to check out, defaults to `master`")
	err = viper.BindPFlag("revision", syncCmd.Flags().Lookup("revision"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
