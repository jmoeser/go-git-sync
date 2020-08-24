package consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/rs/zerolog/log"
)

func PublishKV(key string, key_values []byte) error {

	log.Debug().Msgf("Got KV %s", key_values)
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal().Err(err)
		return err
	}

	log.Debug().Msg("Connected to Consul...")

	kv := client.KV()

	p := &api.KVPair{Key: key, Value: key_values}
	_, err = kv.Put(p, nil)
	if err != nil {
		log.Fatal().Err(err)
		return err
	}

	log.Debug().Msg("Published successfully")

	return nil

}
