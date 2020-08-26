package consul_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/consul/api"

	"github.com/jmoeser/go-git-sync/consul"
)

func TestPublishKV(t *testing.T) {

	var test_data = map[string]string{"Pink": "Flamingo", "Yellow": "Elephant"}
	var test_key = "animals"

	marshal_data, err := json.Marshal(test_data)
	if err != nil {
		t.Error(err)
	}

	err = consul.PublishKV(test_key, marshal_data)
	if err != nil {
		t.Error(err)
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		t.Error(err)
	}

	kv := client.KV()

	pair, _, err := kv.Get(test_key, nil)
	if err != nil {
		t.Error(err)
	}

	data_from_consul := make(map[string]string)
	err = json.Unmarshal(pair.Value, &data_from_consul)
	if err != nil {
		t.Error(err)
	}

	eq := reflect.DeepEqual(data_from_consul, test_data)
	if !eq {
		t.Error("Data in Consul does not much the data we sent the publish function!")
	}

	_, err = kv.Delete(test_key, nil)
	if err != nil {
		t.Error(err)
	}
}
