package commands

import (
	"fmt"
	"ubuntuInit/core"
)

func Ranger() {
	fmt.Println(core.StdGreen("Installing Ranger"))
	commands := []string{
		"sudo git clone https://github.com/ranger/ranger.git $HOME/ranger",
		`sudo su - -c "cd ${HOME}/ranger; make install"`,
		`sudo su $USER -c "ranger --copy-config=all"`,
		"sudo rm -rf $HOME/ranger",
				`sed -i -e '$a\
\nclass code(Command):\
    def execute(self):\
        dirname = self.fm.thisdir.path\
        codecmd = ["code", dirname]\
        self.fm.execute_command(codecmd)\
' $HOME/.config/ranger/commands.py`,
	}

	core.ExecuteCommands(commands)
}
