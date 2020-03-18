package registry

import (
	"dockerWebhook/src/server"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Registry interface {
	Write(logRegistry LogRegistry) bool
	Read(id int64) (string, error)
}

type LogRegistry struct {
	Id int64 `json:"id"`
	Tag string `json:"tag"`
	Repository string `json:"repository"`
	Response server.Response `json:"response"`
	Deployed bool `json:"deployed"`
}

type registry struct {}

func (r registry) Write(lr LogRegistry) bool {
	jsonContent, err := json.Marshal(lr)

	if err != nil {
		log.Fatalln(err)

		return false
	}

	err = ioutil.WriteFile(fmt.Sprintf("../log/%d", lr.Id), jsonContent, 0600)
	if err != nil {
		log.Fatalln(err)

		return false
	}

	return true
}

func (r registry) Read(id int64) (string, error) {
	body, err := ioutil.ReadFile(fmt.Sprintf("../log/%d", id))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func NewRegistry() Registry {
	return registry{}
}
