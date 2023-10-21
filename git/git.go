package git

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/rs/zerolog/log"
)

func GetTempDir() string {
	dir, err := os.MkdirTemp(os.TempDir(), "-git-sync")
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
		return "", "", err
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

func GetRevisionHash(url string, revision string) (string, error) {
	log.Debug().Msgf("Get hash of revision %s from %s", revision, url)
	r, err := git.Init(memory.NewStorage(), nil)
	if err != nil {
		log.Error().Err(err)
		return "", err
	}

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "source",
		URLs: []string{url},
	})
	if err != nil {
		log.Error().Err(err)
		return "", err
	}

	err = r.Fetch(&git.FetchOptions{
		RemoteName: "source",
	})
	if err != nil {
		log.Error().Err(err)
		return "", err
	}

	log.Debug().Msgf("Resolve source/%s", revision)

	h, err := r.ResolveRevision(plumbing.Revision("source/" + revision))
	if err != nil {
		log.Error().Err(err)
		return "", err
	}

	log.Debug().Msgf("Hash %s", h.String())

	return h.String(), nil
}
