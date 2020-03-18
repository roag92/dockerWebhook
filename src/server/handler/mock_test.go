package handler_test

import (
	"dockerWebhook/src/docker"
	"dockerWebhook/src/registry"
	"errors"
	"fmt"
)

type dockerMock struct {}

func (d dockerMock) Start() {
}

func NewDockerMock() docker.Docker {
	return dockerMock{}
}

type registryMock struct {}

func (r registryMock) Write(lr registry.LogRegistry) bool {
	return lr.Id != 0
}

func (r registryMock) Read(id int64) (string, error) {
	if id == 0 {
		return "registry.LogRegistry{}", errors.New("error reading file")
	}

	return fmt.Sprintf(`{"id": %d}}`, id), nil
}

func NewRegistryMock() registry.Registry {
	return registryMock{}
}
