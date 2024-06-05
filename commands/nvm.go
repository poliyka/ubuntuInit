package commands

import (
	"fmt"
	"ubuntuInit/core"
)

func Nvm() {
	fmt.Println(core.StdGreen("Installing Nvm"))
	version := core.GetGithubRepLatestRelease("nvm-sh", "nvm")
	fmt.Println(core.StdGreen(version))

	commands := []string{
		fmt.Sprintf(`sudo su $USER -c "curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/%s/install.sh" | /bin/bash`, version),
	}

	core.ExecuteCommands(commands)
}
