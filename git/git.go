package git

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func Checkout(url string) string {
	dir, err := ioutil.TempDir(os.TempDir(), "go-git-sync-")
	if err != nil {
		log.Fatal(err)
	}

	// https://pkg.go.dev/github.com/go-git/go-git/v5?tab=doc#CloneOptions
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:   url,
		Depth: 1,
	})

	if err != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
		os.RemoveAll(dir)
		os.Exit(1)
	}

	ref, err := r.Head()
	fmt.Println("Checked out", ref.Hash())

	return dir
}
