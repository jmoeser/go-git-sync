package files_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jmoeser/go-git-sync/files"
)

func TestParseJsonFile(t *testing.T) {
	byteValue, err := files.ParseJsonFile("../example/consul/sample.json")
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
	_, err := files.ParseJsonFile("doesnt_exist.json")
	if err == nil {
		t.Error("Expected an error but none was found")
	}
}

func TestParseJsonDirectory(t *testing.T) {
	_, err := files.ParseJsonFile("../example")
	if err == nil {
		t.Error("Expected an error but none was found")
	}
}
