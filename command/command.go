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

type Command interface {
	Command() string
	GetArgs() []string
}

type Cmd struct {
	Cmd       string
	Args string
	Env []string
}

func (c Cmd) Command() string {
	return c.Cmd
}

func (c Cmd) GetArgs() []string {
	return strings.Split(c.Args, " ")
}

func Run(command Command) Result {
	result, err := execute(command)
	if err != nil {
		return NewResult(command.Command(), "command failed", err)
	}

	return NewResult(command.Command(), result, nil)
}

func execute(command Command) (string, error) {
	cmd := exec.Command(command.Command(), command.GetArgs()...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
