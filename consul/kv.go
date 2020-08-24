package consul

import (
	"encoding/json"

	"github.com/hashicorp/consul/api"
	"github.com/rs/zerolog/log"
)

func PublishKV(key string, key_values map[string]string) error {

	log.Debug().Msgf("Got KV %s", key_values)
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal().Err(err)
		return err
	}

	log.Debug().Msg("Connected to Consul...")

	kv := client.KV()

	marshal_data, err := json.Marshal(key_values)
	if err != nil {
		log.Fatal().Err(err)
		return err
	}

	p := &api.KVPair{Key: key, Value: marshal_data}
	_, err = kv.Put(p, nil)
	if err != nil {
		log.Fatal().Err(err)
		return err
	}

	return nil

}
