package watchers

import (
	"os/exec"
	"bytes"
)

func RunCommand(command string, args...string) WatcherResult {
	result, err := run(command, args...)
	if err != nil {
		return WatcherResult{
			command,
			"command failed",
			err,
			"",
		}
	}

	return WatcherResult{
		command,
		result,
		nil,
		result,
	}
}

func run(command string, arg...string) (string, error) {
	cmd := exec.Command(command, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
