package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"ubuntuInit/commands"
	"ubuntuInit/core"
)

func main() {
	// 創建一個新的配置實例
	resp := &core.Response{}

	// 啟動問題並存儲回答
	err := survey.Ask(core.QS, resp)
	if err != nil {
		fmt.Println("Failed to execute survey:", err)
		return
	}

	// 輸出收集到的配置信息
	switch resp.TerminalType {
	case "Bash":
		core.RcPath = "~/.bashrc"
	case "Zsh":
		core.RcPath = "~/.zshrc"
	}

	// if update and upgrade is selected, run update and upgrade
	if resp.UpdateAndUpgrade {
		commands.UpdateAndUpgrade()
	}

	// if common libs is selected, install common libs
	if resp.CommonLibs {
		commands.CommonLibs()
	}

	// if ranger is selected, install ranger
	for _, choice := range resp.InstallChoices {
		switch choice {
		case "Ranger":
			commands.Ranger()
		case "Nvm":
			commands.Nvm()
		case "Yarn":
			commands.Yarn()
		case "Pyenv":
			commands.Pyenv()
		case "Fzf":
			commands.Fzf()
		case "BashColor":
			commands.BashColor()
			// case "GitAlias":
			// 	commands.GitAlias()
		}
	}
}
