package files_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/jmoeser/go-git-sync/files"
)

func TestParseJsonFile(t *testing.T) {
	byteValue, err := files.ParseJsonOrYamlFile("../example/consul/sample-json.json")
	if err != nil {
		t.Error(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}

func TestParseJsonNonExistantFile(t *testing.T) {
	_, err := files.ParseJsonOrYamlFile("doesnt_exist.json")
	if err == nil {
		t.Error("Expected an error but none was found")
	}
}

func TestParseJsonDirectory(t *testing.T) {
	_, err := files.ParseJsonOrYamlFile("../example")
	if err == nil {
		t.Error("Expected an error but none was found")
	}
}

func TestParseYamlFile(t *testing.T) {
	byteValue, err := files.ParseJsonOrYamlFile("../example/consul/sample-yaml.yaml")
	if err != nil {
		t.Error(err)
	}

	var result map[string]interface{}
	err = yaml.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
