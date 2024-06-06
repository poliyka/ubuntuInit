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
			go func() {
				commands.Nvm()
			}()
		case "Yarn":
			go func() {
				commands.Yarn()
			}()
		case "Pyenv":
			commands.Pyenv()
		case "Fzf":
			go func() {
				commands.Fzf()
			}()
		case "BashColor":
			commands.BashColor()
		case "GitAlias":
			commands.GitAlias()
		}
	}

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
