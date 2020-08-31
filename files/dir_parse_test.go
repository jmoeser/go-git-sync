package files_test

import (
	"fmt"
	"testing"

	"github.com/jmoeser/go-git-sync/files"
)

func TestWalkDir(t *testing.T) {
	fileList, err := files.WalkDir("../example")
	if err != nil {
		t.Error(err)
	}

	// bit messy, need to not hardcode the number of files we expect to find
	if len(fileList) != 3 {
		t.Error("More files listed than expected!")
	}

	fileList, err = files.WalkDir("../example/consul/sample-json.json")
	if err != nil {
		t.Error(err)
	}

	// bit messy, need to not hardcode the number of files we expect to find
	if len(fileList) != 1 {
		t.Error("More files listed than expected!")
	}
}

func TestGetFilesAndData(t *testing.T) {
	parsedFiles, err := files.GetFilesAndData("../example")
	if err != nil {
		t.Error(err)
	}

	if len(parsedFiles) != 3 {
		t.Error("More or less parsed files found!")
	}
}

func TestWalkDirPathNotExist(t *testing.T) {
	if _, err := files.WalkDir("doesntexist"); err == nil {
		fmt.Println(err)
		t.Error("Expected an error when trying to parse non-existant path")
	}
}
