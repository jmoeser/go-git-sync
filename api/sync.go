package api

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoeser/go-git-sync/consul"
	"github.com/jmoeser/go-git-sync/files"
	"github.com/jmoeser/go-git-sync/git"
	"github.com/rs/zerolog/log"
)

func RunConsulSync(source string, filePath string, consulServer string, destinationPrefix string) error {
	log.Debug().Msgf("Begin Consul sync with server %s from file path %s in source repo %s to Consul", consulServer, filePath, source)
	if destinationPrefix != "" {
		log.Debug().Msgf("Will sync to prefix %s", destinationPrefix)
	}

	dir := git.GetTempDir()
	defer os.RemoveAll(dir)

	checkedOutDir, headHash, err := git.Checkout(source, dir)
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
