package files_test

import (
	"testing"

	"github.com/jmoeser/go-git-sync/files"
)

func TestWalkDir(t *testing.T) {
	fileList, err := files.WalkDir("../example")
	if err != nil {
		t.Error(err)
	}

	// bit messy, need to not hardcode the number of files we expect to find
	if len(fileList) != 1 {
		t.Error("More files listed than expected!")
	}

	fileList, err = files.WalkDir("../example/consul/sample.json")
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

	if len(parsedFiles) != 1 {
		t.Error("More or less parsed files found!")
	}
}
