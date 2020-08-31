package api_test

import (
	"path/filepath"
	"strings"
	"testing"

	consul_api "github.com/hashicorp/consul/api"
	"github.com/jmoeser/go-git-sync/api"
)

func TestRunConsulSync(t *testing.T) {
	source := "https://github.com/jmoeser/go-git-sync.git"
	file := "example/consul/sample-json.json"
	consul := "127.0.0.1:8500"

	err := api.RunConsulSync(source, file, consul, "")
	if err != nil {
		t.Error(err)
	}

	client, err := consul_api.NewClient(consul_api.DefaultConfig())
	if err != nil {
		t.Error(err)
	}

	kv := client.KV()
	_, err = kv.Delete(strings.TrimSuffix(file, filepath.Ext(file)), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestRunConsulSyncWithPrefix(t *testing.T) {
	source := "https://github.com/jmoeser/go-git-sync.git"
	file := "example/consul/sample-json.json"
	consul := "127.0.0.1:8500"
	prefix := "test"

	err := api.RunConsulSync(source, file, consul, prefix)
	if err != nil {
		t.Error(err)
	}

	client, err := consul_api.NewClient(consul_api.DefaultConfig())
	if err != nil {
		t.Error(err)
	}

	kv := client.KV()
	_, err = kv.Delete(prefix+strings.TrimSuffix(file, filepath.Ext(file)), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestRunConsulSyncWithDirectory(t *testing.T) {
	source := "https://github.com/jmoeser/go-git-sync.git"
	file := "example/consul"
	consul := "127.0.0.1:8500"
	prefix := "test"

	err := api.RunConsulSync(source, file, consul, prefix)
	if err != nil {
		t.Error(err)
	}

	client, err := consul_api.NewClient(consul_api.DefaultConfig())
	if err != nil {
		t.Error(err)
	}

	kv := client.KV()
	_, err = kv.Delete(prefix+file, nil)
	if err != nil {
		t.Error(err)
	}
}
