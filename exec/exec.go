package exec

//go:generate mockgen -package mock_exec -destination mocks/mocks.go github.com/MiguelMoll/joycast/exec Executor

import (
	"os/exec"
	"strings"
)

// Executor is an interface wrapper around the std Command type
type Executor interface {
	Run(cmd string, args ...string) (Output, error)
}

// Output is a container type for std output and err
type Output struct {
	StdOut string
	StdErr string
}

// cmdr satisfies the Executor interface using normal system commands/process
type cmdr struct{}

// Creates a new Executor
func New() Executor {
	return &cmdr{}
}

// Run runs a command and the given arguments
// This is a blocking command
func (cr *cmdr) Run(cmd string, args ...string) (Output, error) {
	c := exec.Command(cmd, args...)

	stdOut := &strings.Builder{}
	stdErr := &strings.Builder{}
	c.Stdout = stdOut
	c.Stderr = stdErr

	if err := c.Start(); err != nil {
		return Output{
			StdOut: stdOut.String(),
			StdErr: stdErr.String(),
		}, err
	}

	if err := c.Wait(); err != nil {
		return Output{
			StdOut: stdOut.String(),
			StdErr: stdErr.String(),
		}, err
	}

	return Output{
		StdOut: stdOut.String(),
		StdErr: stdErr.String(),
	}, nil
}
