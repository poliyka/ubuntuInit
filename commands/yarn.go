package commands

import (
	"fmt"
	"ubuntuInit/core"
)

func Yarn() {
	defer core.Wg.Done()
	core.Lock.Lock()

	fmt.Println(core.StdGreen("Installing Yarn"))
	commands := []string{
		"curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -",
		`echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list`,
		"sudo apt-get update && sudo apt-get install yarn -y",
	}

	core.ExecuteCommands(commands)
	core.Lock.Unlock()
}
