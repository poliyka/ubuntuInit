package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/deckarep/golang-set/v2"
	"ubuntuInit/commands"
	"ubuntuInit/core"
	"ubuntuInit/enum/InstallOptions"
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

	// 將選項添加到 CurrentChoices 方便後面計算
	core.Ic.CurrentChoices = resp.InstallChoices

	// get sudo permission first
	commands.Sudo()
	wgCount := 0

	// 預先安裝 Apt 依賴
	preAptLibs := mapset.NewSet[string]()
	installChoicesMap := core.Ic.CompareChoices()
	for option, value := range installChoicesMap {
		if value && core.MapContains(core.Ic.AptPreInstallMap, option) {
			preAptLibs = preAptLibs.Union(core.Ic.AptPreInstallMap[option])
		}
	}
	if preAptLibs.Cardinality() > 0 {
		commands.PreAptInstall(preAptLibs.ToSlice())
	}

	// if update and upgrade is selected, run update and upgrade
	if resp.UpdateAndUpgrade {
		go commands.UpdateAndUpgrade()
		wgCount++
	}

	// if common libs is selected, install common libs
	if resp.CommonLibs {
		go commands.CommonLibs()
		wgCount++
	}

	// if ranger is selected, install ranger
	for _, choice := range resp.InstallChoices {
		switch choice {
		case InstallOptions.Ranger.String():
			go commands.Ranger()
			wgCount++
		case InstallOptions.Nvm.String():
			go commands.Nvm()
			wgCount++
		case InstallOptions.Yarn.String():
			go commands.Yarn()
			wgCount++
		case InstallOptions.Pyenv.String():
			go commands.Pyenv()
			wgCount++
		case InstallOptions.Fzf.String():
			go commands.Fzf()
			wgCount++
		case InstallOptions.BashColor.String():
			go commands.BashColor()
			wgCount++
		case InstallOptions.GitAlias.String():
			go commands.GitAlias()
			wgCount++
		}
	}

	core.Wg.Add(wgCount)
	core.Wg.Wait()

	fmt.Println(core.Ic.AptPreInstallMap)

	// 顯示最終消息
	fmt.Println(core.StdGreen("===================================="))
	fmt.Println("Done! Ubuntu Initialize Dependencies:")
	// print TerminalType
	fmt.Println("TerminalType: " + core.StdGreen(resp.TerminalType))

	// fmt.Println(ic.AptPreInstallMap)
	fmt.Println(core.Ic.Options)
	fmt.Println(core.Ic.CurrentChoices)
	fmt.Println(core.Ic.CompareChoices())

	// print InstallChoices
	for choice, value := range installChoicesMap {
		if value {
			fmt.Println(choice, "=", core.StdGreen(value))
		} else {
			fmt.Println(choice, "=", core.StdRed(value))
		}
	}

	// 直接比對的方式，留做參考用
	// for _, qs := range core.QS {
	// 	if qs.Name == "InstallChoices" {
	// 		// 進行型別斷言
	// 		if mp, ok := qs.Prompt.(*survey.MultiSelect); ok {
	// 			for _, choice := range mp.Options {
	// 				compare := core.MapContains(resp.InstallChoices, choice)
	// 				if compare {
	// 					fmt.Println(choice, "=", core.StdGreen(compare))
	// 				} else {
	// 					fmt.Println(choice, "=", core.StdRed(compare))
	// 				}
	// 			}

	// 		}
	// 		break
	// 	}
	// }

	fmt.Println("Finished. Restart your shell or reload config file.")
	fmt.Println(core.StdYellow("source ~/.bashrc"))
	fmt.Println(core.StdGreen("===================================="))

}
