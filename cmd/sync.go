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
	"github.com/jmoeser/go-git-sync/git"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var source string
var syncCommand string

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync changes from Git using the specified command",
	Long:  `Sync changes from specified Git source repo using the command specified`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msgf("sync called - %s %s", source, syncCommand)
		tempDir := git.GetTempDir()
		checkedOutRepo := git.Checkout(source, tempDir)
		log.Debug().Msg(checkedOutRepo)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().StringVarP(&source, "source", "s", "", "Source Git URL")
	syncCmd.Flags().StringVarP(&syncCommand, "cmd", "c", "", "Command to run when change detected in Git")

	err := syncCmd.MarkFlagRequired("source")
	if err != nil {
		log.Error().Err(err)
	}
	err = syncCmd.MarkFlagRequired("cmd")
	if err != nil {
		log.Error().Err(err)
	}
}
