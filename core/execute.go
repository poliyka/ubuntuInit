package core

import (
	"os/exec"
)

func ExecuteCommands(commands []string) error {
	var err error
	for _, cmdStr := range commands {
		cmd := exec.Command("/bin/bash", "-c", cmdStr)

		err = HandleError(cmd)

		if err != nil {
			break
		}
	}
	return err
}
