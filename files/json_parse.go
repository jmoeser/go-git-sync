package files

import (
	"io/ioutil"

	"github.com/rs/zerolog/log"
)

func ParseJson(path string) ([]byte, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	return data, err

}
