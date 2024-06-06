package commands

import (
	"os/exec"
	"ubuntuInit/core"
)

func UpdateAndUpgrade() {
	defer core.Wg.Done()
	core.Lock.Lock()

	cmdStr := "sudo apt update && sudo apt upgrade -y"

	cmd := exec.Command("/bin/bash", "-c", cmdStr)

	core.HandleError(cmd)

	core.Lock.Unlock()
}
