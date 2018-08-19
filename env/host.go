package env

import (
	"github.com/gravitational/log"
	"os"
)

func GetCurrentHost() string {
	host, err := os.Hostname()
	if err != nil {
		log.Warningf("Cannot detect current host: %v", err)
		host = "unknown host"
	}

	return host
}
