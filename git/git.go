package git

import (
	"io/ioutil"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/rs/zerolog/log"
)

func Checkout(url string) string {
	dir, err := ioutil.TempDir(os.TempDir(), "go-git-sync-")
	if err != nil {
		log.Error().Err(err)
	}

	// https://pkg.go.dev/github.com/go-git/go-git/v5?tab=doc#CloneOptions
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:   url,
		Depth: 1,
	})

	if err != nil {
		log.Error().Err(err)
		os.RemoveAll(dir)
		os.Exit(1)
	}

	ref, err := r.Head()
	log.Debug().Msgf("Checked out %s", ref.Hash())

	return dir
}
