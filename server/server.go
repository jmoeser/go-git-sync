package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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

	now := time.Now()
	unixTimestamp := strconv.FormatInt(now.Unix(), 10)
	heartbeatData := map[string]string{"timestamp": unixTimestamp, "readableTime": now.Format("2006-01-02 15:04:05")}

	jsonData, err := json.Marshal(heartbeatData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
