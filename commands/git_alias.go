package commands

import (
	"fmt"
	"ubuntuInit/core"
)

func GitAlias() {
	fmt.Println(core.StdGreen("Installing BashColor"))

	GITCONFIG_PATH := "$HOME/.gitconfig"

	installCmd := `
	GITCONFIG_PATH=%s
	if [[ ! -f "$GITCONFIG_PATH" ]]; then
	  sudo su $USER -c "touch ${GITCONFIG_PATH}"
	fi

	if ! sed -n '/\[alias\]/p' $GITCONFIG_PATH | grep '[alias]'; then
	  sudo su $OE_USER -c "printf '[alias]\n' >> $GITCONFIG_PATH"
	fi
	`

	GITCONFIG_PATH = fmt.Sprintf(installCmd, GITCONFIG_PATH)

	insertStr := `
st = status
cm = commit
ch = checkout
sw = switch
br = branch
mg = merge
acm =  "!git add -A && git commit -m"
mgd = "!git mg $1 && git br -d $1; #"
# 查看分支(樹狀圖)
lg = log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit
# 查看reflog(HeighLight)
rl = reflog --pretty=format:\"%Cred%h%Creset %C(auto)%gd%Creset %C(auto)%gs%C(reset) %C(green)(%cr)%C(reset) %C(bold blue)<%an>%Creset\" --abbrev-commit
# 查看stash(HeighLight)
sl = stash list --pretty=format:\"%C(red)%h%C(reset) - %C(dim yellow)(%C(bold magenta)%gd%C(dim yellow))%C(reset) %<(70,trunc)%s %C(green)(%cr) %C(bold blue)<%an>%C(reset)\"
`

	commands := []string{installCmd}
	commands = append(commands, fmt.Sprintf(`echo '%s' >> %s`, insertStr, GITCONFIG_PATH))

	core.ExecuteCommands(commands)
}
