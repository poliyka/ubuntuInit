package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"ubuntuInit/core"
)

func GitAlias() {
	defer core.Wg.Done()

	fmt.Println(core.StdGreen("Installing GitAlias"))

	GITCONFIG_PATH := os.Getenv("HOME") + "/.gitconfig"

	// Check if .gitconfig file exists, create if it doesn't
	if _, err := os.Stat(GITCONFIG_PATH); os.IsNotExist(err) {
		file, err := os.Create(GITCONFIG_PATH)
		if err != nil {
			fmt.Println("Error creating .gitconfig file:", err)
			return
		}
		defer file.Close()
	}

	file, err := os.OpenFile(GITCONFIG_PATH, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening .gitconfig file:", err)
		return
	}
	defer file.Close()

	aliasZone := []string{
		"st = status",
		"cm = commit",
		"ch = checkout",
		"sw = switch",
		"br = branch",
		"mg = merge",
		"acm = \"!git add -A && git commit -m\"",
		"mgd = \"!git mg $1 && git br -d $1; #\"",
		"# 查看分支(樹狀圖)",
		"lg = log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit",
		"# 查看reflog(HeighLight)",
		"rl = reflog --pretty=format:\"%Cred%h%Creset %C(auto)%gd%Creset %C(auto)%gs%C(reset) %C(green)(%cr)%C(reset) %C(bold blue)<%an>%Creset\" --abbrev-commit",
		"# 查看stash(HeighLight)",
		"sl = stash list --pretty=format:\"%C(red)%h%(reset) - %C(dim yellow)(%C(bold magenta)%gd%C(dim yellow))%(reset) %%(70,trunc)%s %C(green)(%cr) %C(bold blue)<%an>%(reset)\"",
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	var AliasLines []string
	aliasFound := false
	insideAlias := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "[alias]" {
			aliasFound = true
			insideAlias = true
		} else if strings.HasPrefix(strings.TrimSpace(line), "[") {
			insideAlias = false
		}

		if insideAlias {
			AliasLines = append(AliasLines, line)
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading .gitconfig file:", err)
		return
	}

	if !aliasFound {
		// Append [alias] section and the aliasZone at the end
		AliasLines = append(AliasLines, "[alias]")
		AliasLines = append(AliasLines, aliasZone...)
	} else {
		AliasLines = append(AliasLines, aliasZone...)
	}

	// Write back the updated content to .gitconfig file
	file, err = os.OpenFile(GITCONFIG_PATH, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening .gitconfig file for writing:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	for _, line := range AliasLines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

}
