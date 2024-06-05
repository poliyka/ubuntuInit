package commands

import (
	"os/exec"
	"ubuntuInit/core"
)

func Ranger() {
	cmdStr := "sudo apt update && sudo apt upgrade -y"

	cmd := exec.Command("/bin/bash", "-c", cmdStr)

	core.HandleError(cmd)
}
