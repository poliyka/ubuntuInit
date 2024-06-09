package commands

import (
	"fmt"
	"ubuntuInit/core"
	"ubuntuInit/enum/InstallOptions"
)

func init() {
	core.InstallAptCollection(InstallOptions.Fzf, []string{"git"})
}

func Fzf() {
	defer core.Wg.Done()

	fmt.Println(core.StdGreen("Installing Fzf"))

	installCmd := `sudo su $USER -c ${HOME}/.fzf/install << 'EOF'
y
y
y
EOF`

	commands := []string{
		`sudo su $USER -c "git clone --depth 1 https://github.com/junegunn/fzf.git $HOME/.fzf"`,
		installCmd,
	}

	core.ExecuteCommands(commands)
}
