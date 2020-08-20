package json_parse_test

import (
	"encoding/json"
	"fmt"
	"testing"

	json_parse "github.com/jmoeser/go-git-sync/files"
)

func TestParseJson(t *testing.T) {
	byteValue, err := json_parse.ParseJson("../example/consul/sample.json")
	if err != nil {
		t.Error(err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)
}

func TestParseJsonNonExistantFile(t *testing.T) {
	_, err := json_parse.ParseJson("doesnt_exist.json")
	if err == nil {
		t.Error("Expected an error but none was found")
	}
}
