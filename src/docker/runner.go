package docker

import (
	"fmt"
	"os/exec"
)

type Runner interface {
	run(string, ...string)
}

type DefaultRunner struct {
	Command func(name string, arg ...string) *exec.Cmd
}

func (dr DefaultRunner) run(command string, args ...string) {
	out, err := dr.Command(command, args...).CombinedOutput()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
