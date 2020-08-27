package files

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func isParsableFile(extension string) bool {
	switch extension {
	case
		".json",
		".yaml",
		".yml",
		".hcl":
		return true
	}
	return false
}

func WalkDir(path string) ([]string, error) {

	fileList := make([]string, 0)
	log.Debug().Msgf("Will recursively search path %s for files we're looking for", path)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal().Err(err)
		}

		if (!info.IsDir()) && (isParsableFile(filepath.Ext(info.Name()))) {
			fileList = append(fileList, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	log.Debug().Msgf("Found files: %s", fileList)

	return fileList, nil
}
