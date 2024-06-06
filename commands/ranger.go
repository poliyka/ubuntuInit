package commands

import (
	"fmt"
	"ubuntuInit/core"
)

func Ranger() {
	defer core.Wg.Done()

	fmt.Println(core.StdGreen("Installing Ranger"))

	// 分兩階段執行 ranger 安裝
	// 先嘗試安裝 make 指令, 因為 make 指令安裝會與其他 thread 有衝突這邊要 Lock
	core.Lock.Lock()
	commands := []string{
		"sudo apt-get install make -y",
	}

	core.ExecuteCommands(commands)
	core.Lock.Unlock()

	// 安裝 ranger
	commands = []string{
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

	err := core.ExecuteCommands(commands)

	// 當發生錯誤時因該要刪掉 ranger 資料夾
	if err != nil {
		commands := []string{"sudo rm -rf $HOME/ranger"}
		core.ExecuteCommands(commands)
	}

}
