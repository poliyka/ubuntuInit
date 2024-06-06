package commands

import (
	"fmt"
	"ubuntuInit/core"
)

func Sudo() {
	fmt.Println(core.StdGreen("Get Sudo Permission"))
	commands := []string{
		"sudo date",
	}

	core.ExecuteCommands(commands)
}
