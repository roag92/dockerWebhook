package docker_test

import (
	"dockerWebhook/src/docker"
	"testing"
)

func TestNewDocker(t *testing.T) {
	d := docker.NewDocker(docker.DefaultRunner{Command})

	assertDockerInterface(d, t)
}

func TestDocker_Start(t *testing.T) {
	d := docker.NewDocker(docker.DefaultRunner{Command})

	assertDockerInterface(d, t)

	d.Start()
}

func assertDockerInterface(d docker.Docker, t *testing.T)  {
	_, ok := d.(docker.Docker)

	if ok != true {
		t.Error("expecting Docker interface")
	}
}
