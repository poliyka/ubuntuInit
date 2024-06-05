package commands

import (
	"fmt"
	"ubuntuInit/core"
)

func Pyenv() {
	fmt.Println(core.StdGreen("Installing Pyenv"))

	// 定義 .bashrc 配置內容
	bashrc := `
export PYENV_ROOT="$HOME/.pyenv"
command -v pyenv >/dev/null || export PATH="$PYENV_ROOT/bin:$PATH"
eval "$(pyenv init -)"
eval "$(pyenv virtualenv-init -)"
`

	// 定義 .profile 配置內容
	profile := `
export PYENV_ROOT="$HOME/.pyenv"
command -v pyenv >/dev/null || export PATH="$PYENV_ROOT/bin:$PATH"
eval "$(pyenv init -)"
`

	// 定義 .zshrc 配置內容
	zshrc := `
if ! grep -q 'export PYENV_ROOT="$HOME/.pyenv"' ~/.zshrc; then
	if grep -q 'plugins=(' ~/.zshrc; then
	  sed -i '0,/plugins=(/s//export PYENV_ROOT="$HOME\/.pyenv"\nexport PATH="$PYENV_ROOT\/bin:$PATH"\neval "$(pyenv init --path)"\n&/' ~/.zshrc
	else
	  echo '
export PYENV_ROOT="$HOME/.pyenv"
export PATH="$PYENV_ROOT/bin:$PATH"
eval "$(pyenv init --path)"
' >> ~/.zshrc
	fi
fi
`

	// 定義需要執行的命令
	commands := []string{
		`sudo su $USER -c "curl -L https://raw.github.com/yyuu/pyenv-installer/master/bin/pyenv-installer | bash"`,
	}

	// 根據不同的 shell 配置，追加對應的命令
	switch core.RcPath {
	case "~/.bashrc":
		commands = append(commands, fmt.Sprintf(`echo '%s' >> %s`, bashrc, core.RcPath))
		commands = append(commands, fmt.Sprintf(`echo '%s' >> %s`, profile, "~/.profile"))
	case "~/.zshrc":
		commands = append(commands, zshrc)
	}

	// 執行所有命令
	core.ExecuteCommands(commands)
}
