package consul

import (
	consul_api "github.com/hashicorp/consul/api"
	"github.com/rs/zerolog/log"
)

func getClient(config *consul_api.Config) (*consul_api.Client, error) {

	client, err := consul_api.NewClient(config)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	log.Debug().Msg("Got Consul client")

	return client, err
}

func GetKVHandler(config *consul_api.Config) (*consul_api.KV, error) {

	client, err := getClient(config)
	if err != nil {
		log.Fatal().Err(err)
		return nil, err
	}

	kv := client.KV()

	return kv, err
}
