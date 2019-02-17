package transports

import (
	"github.com/sepuka/gowatcher/command"
)

const (
	contentTypeJson = "application/json"
)

type Transport interface {
	Send(msg command.Result) (err error)
}
