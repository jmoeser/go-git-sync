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
package main

import (
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	"github.com/jmoeser/go-git-sync/cmd"
	"github.com/rs/zerolog"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs)

	go func() {
		for {
			// select {
			// case s := <-sigs:
			// 	switch s {
			// 	case os.Interrupt:
			// 		log.Debug().Msgf("RECEIVED SIGNAL: %s", s)
			// 		os.Exit(1)
			// 	}
			// }
			for s := range sigs {
				if s == os.Interrupt {
					log.Debug().Msgf("RECEIVED SIGNAL: %s", s)
					os.Exit(1)
				}

			}
		}
	}()

	cmd.Execute()
}
