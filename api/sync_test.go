package api_test

import (
	"testing"

	consul_api "github.com/hashicorp/consul/api"
	"github.com/jmoeser/go-git-sync/api"
)

func TestRunConsulSync(t *testing.T) {
	source := "https://github.com/jmoeser/go-git-sync.git"
	file := "example/consul/sample.json"
	consul := "127.0.0.1:8500"
	destination := "animals/data"

	err := api.RunConsulSync(source, file, consul, destination)
	if err != nil {
		t.Error(err)
	}

	client, err := consul_api.NewClient(consul_api.DefaultConfig())
	if err != nil {
		t.Error(err)
	}

	kv := client.KV()
	_, err = kv.Delete(destination, nil)
	if err != nil {
		t.Error(err)
	}
}
