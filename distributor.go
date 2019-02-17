package main

import (
	"github.com/sepuka/gowatcher/definition/transport"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/transports"
)

func transmitter() {
	for {
		msg := <-watcherResult

		for _, sender := range services.GetByTag(transport.DefTransportTag) {
			go sender.(transports.Transport).Send(msg)
		}
	}
}
