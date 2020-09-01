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
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var source string
var filePath string
var destinationPrefix string
var revision string

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

		if err := api.RunConsulSync(source, filePath, consulServer, destinationPrefix, revision); err != nil {
			log.Error().Err(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	defaultRevision := "master"

	syncCmd.Flags().StringVarP(&source, "source", "s", "", "Source Git URL")
	syncCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path in the Git repo")
	syncCmd.Flags().StringVarP(&destinationPrefix, "prefix", "p", "", "Prefix of the path to sync to in Consul")
	syncCmd.Flags().StringVarP(&revision, "revision", "r", defaultRevision, "Revision to check out, defaults to `master`")

	err := syncCmd.MarkFlagRequired("source")
	if err != nil {
		log.Error().Err(err)
	}
	err = syncCmd.MarkFlagRequired("file")
	if err != nil {
		log.Error().Err(err)
	}

}
