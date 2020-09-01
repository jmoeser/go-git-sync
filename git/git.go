package git

import (
	"io/ioutil"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"
)

func GetTempDir() string {
	dir, err := ioutil.TempDir(os.TempDir(), "go-git-sync-")
	if err != nil {
		log.Error().Err(err)
	}

	log.Debug().Msgf("Created temp directory %s", dir)

	return dir
}

func Checkout(url string, revision string, checkoutDir string) (string, string, error) {

	// https://pkg.go.dev/github.com/go-git/go-git/v5?tab=doc#CloneOptions
	r, err := git.PlainClone(checkoutDir, false, &git.CloneOptions{
		URL:   url,
		Depth: 1,
	})

	if err != nil {
		log.Error().Err(err)
		os.RemoveAll(checkoutDir)
		return "", "", nil
	}

	ref, err := r.Head()
	if err != nil {
		log.Error().Err(err)
	}
	log.Debug().Msgf("Checked out %s", ref.Hash())

	w, err := r.Worktree()
	if err != nil {
		log.Error().Err(err)
	}

	log.Debug().Msgf("git checkout %s", revision)
	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(revision),
	})
	if err != nil {
		log.Error().Err(err)
	}

	ref, err = r.Head()
	if err != nil {
		log.Error().Err(err)
	}
	log.Debug().Msgf("Checked out %s", ref.Hash())

	return checkoutDir, ref.Hash().String(), nil
}
