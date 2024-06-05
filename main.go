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

	fmt.Printf("Update && Upgrade: %t\n", config.UpdateAndUpgrade)
	if config.UpdateAndUpgrade {
		commands.UpdateAndUpgrade()
	}
	fmt.Printf("Update && Upgrade: %t\n", config.CommonLibs)
	if config.CommonLibs {
		commands.CommonLibs()
	}
	fmt.Printf("Package Choices: %v\n", config.InstallChoices)
}
