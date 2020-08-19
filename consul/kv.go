package consul_kv

import (
	"encoding/json"

	"github.com/hashicorp/consul/api"
	"github.com/rs/zerolog/log"
)

func ConsulPublishKV(key string, key_values map[string]string) {

	log.Debug().Msgf("Got KV %s", key_values)
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal().Err(err)
		return
	}

	log.Debug().Msg("Connected to Consul...")

	kv := client.KV()

	marshal_data, err := json.Marshal(key_values)
	if err != nil {
		log.Fatal().Err(err)
	}

	p := &api.KVPair{Key: key, Value: marshal_data}
	_, err = kv.Put(p, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

}
