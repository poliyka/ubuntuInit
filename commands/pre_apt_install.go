package commands

import (
	"fmt"
	"strings"
	"ubuntuInit/core"
)

func PreAptInstall(libs []string) {
	fmt.Println(core.StdGreen("Pre Apt Installing"))

	// join libs and make commands word like "sudo apt-get install -y lib1 lib2 lib3"
	commands := []string{
		"sudo apt-get update",
		"sudo apt-get install -y " + strings.Join(libs, " "),
	}

	core.ExecuteCommands(commands)
}
