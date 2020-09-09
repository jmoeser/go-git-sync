package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/go-github/v32/github"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
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

	http.HandleFunc("/heartbeat", a.heartbeatHandler)

	webhookSecret := viper.GetString("webhook-secret")
	if webhookSecret != "" {
		log.Info().Msg("Got webhook secret, enabling webhook handler")
		http.HandleFunc("/webhook", a.webhookHandler)
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal().Err(err)
		os.Exit(1)
	}

}

func (a *GoGitSyncServer) heartbeatHandler(w http.ResponseWriter, req *http.Request) {

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

func (a *GoGitSyncServer) webhookHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		payload, err := github.ValidatePayload(req, []byte(viper.GetString("webhook-secret")))
		if err != nil {
			log.Error().Msgf("Error validating webhook: %s", err)
			return
		}
		defer req.Body.Close()

		event, err := github.ParseWebHook(github.WebHookType(req), payload)
		if err != nil {
			log.Error().Msgf("Could not parse webhook: %s", err)
			return
		}

		switch e := event.(type) {
		case *github.PushEvent:
			revision := *e.After
			repoUrl := *e.Repo.HTMLURL
			pusher := *e.Pusher.Name
			log.Info().Msgf("Got push event from webhook, repo %s, hash %s, push by %s", repoUrl, revision, pusher)
			log.Debug().Msgf("Will sync to %s", a.GoGitSyncServerOptions.ConsulHost)
			// api.RunConsulSync(repoUrl, filePath, consulServer, destinationPrefix, revision)
		default:
			log.Printf("unknown event type %s\n", github.WebHookType(req))
			return
		}

	} else {
		http.Error(w, "Invalid request method.", 405)
	}
}
