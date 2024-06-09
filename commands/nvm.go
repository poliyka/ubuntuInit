package commands

import (
	"fmt"
	"ubuntuInit/core"
	"ubuntuInit/enum/InstallOptions"
)

func init() {
	core.InstallAptCollection(InstallOptions.Nvm, []string{"curl"})
}

func Nvm() {
	defer core.Wg.Done()

	fmt.Println(core.StdGreen("Installing Nvm"))
	version := core.GetGithubRepLatestRelease("nvm-sh", "nvm")

	commands := []string{
		fmt.Sprintf(`sudo su $USER -c "curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/%s/install.sh" | /bin/bash`, version),
	}

	core.ExecuteCommands(commands)
}
