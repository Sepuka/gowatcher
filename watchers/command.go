package watchers

import (
	"os/exec"
	"bytes"
)

func Run(command string, arg...string) (string, error) {
	cmd := exec.Command(command, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}


func RunCommand(command string) WatcherResult {
	result, err := Run(command)
	if err != nil {
		return WatcherResult{
			command,
			"",
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