package core

import (
	"os/exec"
)

func ExecuteCommands(commands []string) {
	for _, cmdStr := range commands {
		cmd := exec.Command("/bin/bash", "-c", cmdStr)

		err := HandleError(cmd)

		if err != nil {
			break
		}
	}
}
