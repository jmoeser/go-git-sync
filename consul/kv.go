package consul

import (
	consul_api "github.com/hashicorp/consul/api"
	"github.com/rs/zerolog/log"
)

func PublishKV(key string, key_values []byte) error {

	log.Debug().Msgf("Got KV %s", key_values)

	kv, err := GetKVHandler(consul_api.DefaultConfig())
	if err != nil {
		log.Fatal().Err(err)
		return err
	}

	p := &consul_api.KVPair{Key: key, Value: key_values}
	_, err = kv.Put(p, nil)
	if err != nil {
		log.Fatal().Err(err)
		return err
	}

	log.Debug().Msg("Published successfully")

	return nil

}
