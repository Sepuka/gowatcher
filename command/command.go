package command

import (
	"bytes"
	"os/exec"
	"strings"
)

type ContentType string

const (
	PlainTextContent ContentType = "plain_text_content"
	ImageContent     ContentType = "image_content"
)

type Result struct {
	watcherName string
	content     string
	error       error
	contentType ContentType
}

func NewImgResult(name string, content string) Result {
	return Result{name, content, nil, ImageContent}
}

func NewResult(name string, text string, err error) Result {
	return Result{name, text, err, PlainTextContent}
}

func (r Result) GetContent() string {
	return r.content
}

func (r Result) GetError() error {
	return r.error
}

func (r Result) IsFailure() bool {
	return r.error != nil
}

func (r Result) GetName() string {
	return r.watcherName
}

func (r Result) GetType() ContentType {
	return r.contentType
}

type Cmd struct {
	cmd  string
	args []string
	envs []string
}

func NewCmd(cmd, args string) *Cmd {
	return &Cmd{
		cmd:  cmd,
		args: strings.Split(args, " "),
	}
}

func NewEnvedCmd(cmd, args string, env string) *Cmd {
	return &Cmd{
		cmd:  cmd,
		args: strings.Split(args, " "),
		envs: strings.Split(env, " "),
	}
}

func runConsoleCommand(command *Cmd) Result {
	result, err := execute(command)
	if err != nil {
		return NewResult(command.cmd, err.Error(), err)
	}

	return NewResult(command.cmd, result, nil)
}

func execute(command *Cmd) (string, error) {
	cmd := exec.Command(command.cmd, command.args...)
	cmd.Env = command.envs
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return out.String(), err
	}

	return out.String(), nil
}
