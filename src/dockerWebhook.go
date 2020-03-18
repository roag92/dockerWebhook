package dockerWebHook

import (
	"dockerWebhook/src/docker"
	"dockerWebhook/src/registry"
	"dockerWebhook/src/server"
	"dockerWebhook/src/server/handler"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const ConfigDirectory string = "config"
const SecretKey string = "SECRET_KEY"

func loadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func chDir() {
	err := os.Chdir(ConfigDirectory)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(d docker.Docker, r registry.Registry) map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"/webhook/" + os.Getenv(SecretKey): handler.NewWebhookHandler(d, r),
		"/result/": handler.NewResultHandler(r),
	}
}


func Launch() {
	loadEnvironment()
	chDir()

	d := docker.NewDocker(docker.DefaultRunner{exec.Command})
	d.Start()

	s := server.NewServer(routes(d, registry.NewRegistry()))
	s.Serve()
}
