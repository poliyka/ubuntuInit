package commands

import (
	"fmt"
	"ubuntuInit/core"
)

func BashColor() {
	defer core.Wg.Done()

	fmt.Println(core.StdGreen("Installing BashColor"))

	installCmd := `sed -i -e '$a\
\
parse_git_branch() {\
	git branch 2> /dev/null | sed -e \"/^[^*]/d\" -e \"s/* \\(.*\\)/ (\\1)/\"\
}\
export PS1=\"\\[\\033[01;32m\\]\\u@\\h\\[\\e[91m\\]\\$(parse_git_branch) \\[\\e[1;33m\\]\\D{%Y/%m/%d} \\t\\[\\033[00m\\]:\\n\\[\\e[34m\\]\\w\\[\\e[00m\\]\$ \"
' $HOME/.bashrc`

	commands := []string{installCmd}

	core.ExecuteCommands(commands)
}
