package transports

import (
	"github.com/sepuka/gowatcher/command"
	"net/http"
)

const (
	contentTypeJson = "application/json"
)

type Transport interface {
	Send(msg command.Result) (resp *http.Response, err error)
}