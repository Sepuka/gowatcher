package command

import (
	"bytes"
	"os/exec"
)

type Result struct {
	watcherName string
	text        string
	error       error
}

func NewResult(name string, text string, err error) Result {
	return Result{name, text, err}
}

func (r *Result) GetText() string {
	return r.text
}

func (r *Result) GetError() error {
	return r.error
}

func (r *Result) IsFailure() bool {
	return r.error != nil
}

func (r *Result) GetName() string {
	return r.watcherName
}

type Command interface {
	Command() string
	Args() []string
}

type Cmd struct {
	cmd  string
	args []string
}

func NewCmd(cmd string, args []string) Cmd {
	return Cmd{cmd, args}
}

func (c Cmd) Command() string {
	return c.cmd
}

func (c Cmd) Args() []string {
	return c.args
}

func Run(command Command) Result {
	result, err := execute(command)
	if err != nil {
		return Result{
			command.Command(),
			"command failed",
			err,
		}
	}

	return Result{
		command.Command(),
		result,
		nil,
	}
}

func execute(command Command) (string, error) {
	cmd := exec.Command(command.Command(), command.Args()...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
