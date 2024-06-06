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

	// get sudo permission first
	commands.Sudo()

	// if update and upgrade is selected, run update and upgrade
	if resp.UpdateAndUpgrade {
		go commands.UpdateAndUpgrade()
	}

	// if common libs is selected, install common libs
	if resp.CommonLibs {
		go commands.CommonLibs()
	}

	// if ranger is selected, install ranger
	for _, choice := range resp.InstallChoices {
		switch choice {
		case "Ranger":
			go commands.Ranger()
		case "Nvm":
			go commands.Nvm()
		case "Yarn":
			go commands.Yarn()
		case "Pyenv":
			go commands.Pyenv()
		case "Fzf":
			go commands.Fzf()
		case "BashColor":
			go commands.BashColor()
		case "GitAlias":
			go commands.GitAlias()
		}
	}

	core.Wg.Add(9)
	core.Wg.Wait()

	// print the final message
	fmt.Println(core.StdGreen("===================================="))
	fmt.Println("Done! Ubuntu Initialize Dependencies:")
	// print TerminalType
	fmt.Println("TerminalType: " + core.StdGreen(resp.TerminalType))
	// print InstallChoices
	for _, qs := range core.QS {
		if qs.Name == "InstallChoices" {
			// 進行型別斷言
			if mp, ok := qs.Prompt.(*survey.MultiSelect); ok {
				for _, choice := range mp.Options {
					compare := core.MapContains(resp.InstallChoices, choice)
					if compare {
						fmt.Println(choice, "=", core.StdGreen(compare))
					} else {
						fmt.Println(choice, "=", core.StdRed(compare))
					}
				}

			}
			break
		}
	}
	fmt.Println("Finished. Restart your shell or reload config file.")
	fmt.Println(core.StdYellow("source ~/.bashrc"))
	fmt.Println(core.StdGreen("===================================="))
}
