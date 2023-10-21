package files

import (
	"os"

	"github.com/rs/zerolog/log"
)

func ParseJsonOrYamlFile(path string) ([]byte, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	return data, err

}
