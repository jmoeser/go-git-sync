package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

type GoGitSyncServer struct {
	GoGitSyncServerOptions
}

type GoGitSyncServerOptions struct {
	ConsulHost string
}

func NewServer(opts GoGitSyncServerOptions) *GoGitSyncServer {

	return &GoGitSyncServer{
		GoGitSyncServerOptions: opts,
	}
}

func (a *GoGitSyncServer) Run(port int, metricsPort int) {
	log.Info().Msgf("Go Git Sync server started on port %d", port)

	http.HandleFunc("/heartbeat", heartbeat)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal().Err(err)
		os.Exit(1)
	}

}

func heartbeat(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "heartbeat\n")
}
