package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"ubuntuInit/commands"
	"ubuntuInit/core"
)

func main() {
	// 創建一個新的配置實例
	config := &core.Config{}

	// 啟動問題並存儲回答
	err := survey.Ask(core.QS, config)
	if err != nil {
		fmt.Println("Failed to execute survey:", err)
		return
	}

	// 輸出收集到的配置信息
	// if update and upgrade is selected, run update and upgrade
	if config.UpdateAndUpgrade {
		commands.UpdateAndUpgrade()
	}

	// if common libs is selected, install common libs
	if config.CommonLibs {
		commands.CommonLibs()
	}

	// if ranger is selected, install ranger
	for _, choice := range config.InstallChoices {
		switch choice {
		case "Ranger":
			commands.Ranger()
			case "Nvm":
				commands.Nvm()
			case "Yarn":
				commands.Yarn()
			// case "Pyenv":
			// 	commands.Pyenv()
			// case "Fzf":
			// 	commands.Fzf()
			// case "BashColor":
			// 	commands.BashColor()
			// case "GitAlias":
			// 	commands.GitAlias()
		}
	}
}
