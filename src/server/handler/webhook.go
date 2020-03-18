package handler

import (
	"dockerWebhook/src/docker"
	"dockerWebhook/src/registry"
	"dockerWebhook/src/server"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const TagKey string = "TAG_KEY"

func NewWebhookHandler(d docker.Docker, reg registry.Registry) func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		var p server.Payload

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if p.PushData.PushedAt == 0 || p.PushData.Tag == "" || p.Repository.RepoName == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		response := buildResponse(p.PushData.PushedAt, p.Repository.RepoName)
		deployed := false

		if p.PushData.Tag == os.Getenv(TagKey) {
			d.Start()
			deployed = true
		} else {
			response.Description = "Tag not deployed"
		}

		reg.Write(registry.LogRegistry{
			p.PushData.PushedAt,
			p.PushData.Tag,
			p.Repository.RepoName,
			response,
			deployed,
		})

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonResponse)
	}
}

func buildHostUrl(url string, port string) string {
	var host = url

	if port != "80" {
		host = fmt.Sprintf("%s:%s", url, port)
	}

	return host
}

func buildResponse(pushedAt int64, repoName string) server.Response {
	targetUrl := fmt.Sprintf(
		"%s/%s/%d",
		buildHostUrl(os.Getenv(server.HostUrlKey), os.Getenv(server.HostPortKey)),
		"result",
		pushedAt,
	)

	return server.Response{
		server.StateSuccess,
		"Service deployed",
		fmt.Sprintf("CD by %s", repoName),
		targetUrl,
	}
}
