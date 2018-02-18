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
