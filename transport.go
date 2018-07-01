package main

import (
	"github.com/sepuka/gowatcher/watchers"
)

func Transmitter(c <-chan watchers.WatcherResult) {
	for {
		msg := <- c
		watchers.SendMessage(msg, config)
	}
}
