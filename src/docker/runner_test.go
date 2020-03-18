package docker_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const TestStdoutValue = "test fun value!"

func TestCommandSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	_, _ = fmt.Fprintf(os.Stdout, TestStdoutValue)
	os.Exit(0)
}

func Command(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestCommandSuccess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_TEST_PROCESS=1"}
	return cmd
}
