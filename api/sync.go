package api

import (
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jmoeser/go-git-sync/consul"
	"github.com/jmoeser/go-git-sync/files"
	"github.com/jmoeser/go-git-sync/git"
	"github.com/rs/zerolog/log"
)

func StartSyncLoop(source string, filePath string, consulServer string, destinationPrefix string, revision string) {

	for {
		log.Debug().Msgf("Started sync loop")

		if err := RunConsulSync(source, filePath, consulServer, destinationPrefix, revision); err != nil {
			log.Error().Err(err)
			os.Exit(1)
		}

		Nsecs := rand.Intn(30000)
		log.Debug().Msgf("About to sleep %dms before looping again", Nsecs)
		time.Sleep(time.Millisecond * time.Duration(Nsecs))

	}

}

func RunConsulSync(source string, filePath string, consulServer string, destinationPrefix string, revision string) error {
	log.Debug().Msgf("Begin Consul sync with server %s from file path %s in source repo %s to Consul", consulServer, filePath, source)
	if destinationPrefix != "" {
		log.Debug().Msgf("Will sync to prefix %s", destinationPrefix)
	}

	dir := git.GetTempDir()
	defer os.RemoveAll(dir)

	checkedOutDir, headHash, err := git.Checkout(source, revision, dir)
	if err != nil {
		log.Fatal().Err(err)
		return err
	}

	log.Debug().Msgf("Checked out repo at commit %s", headHash)

	parsedData, err := files.GetFilesAndData(filepath.Join(checkedOutDir, filePath))
	if err != nil {
		log.Fatal().Err(err)
		return err
	}

	for key, value := range parsedData {
		destinationKey := strings.Replace(key, checkedOutDir+"/", "", -1)
		destinationKey = strings.TrimSuffix(destinationKey, filepath.Ext(destinationKey))

		if destinationPrefix != "" {
			destinationKey = destinationPrefix + "/" + destinationKey
		}

		log.Debug().Msgf("Publishing to %s", destinationKey)

		err = consul.PublishKV(destinationKey, value)
		if err != nil {
			log.Fatal().Err(err)
			return err
		}
	}

	log.Debug().Msg("Sync operation completed successfully")

	return nil
}
