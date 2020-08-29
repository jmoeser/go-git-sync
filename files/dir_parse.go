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

func GetFilesAndData(path string) (map[string][]byte, error) {

	parsedFiles := make(map[string][]byte)

	fileList, err := WalkDir(path)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	for _, name := range fileList {

		switch extension := filepath.Ext(name); extension {
		case ".json":
			byteData, err := ParseJsonFile(name)
			if err != nil {
				log.Fatal().Err(err)
				return nil, err
			}
			parsedFiles[name] = byteData

		}

	}

	return parsedFiles, nil

}
