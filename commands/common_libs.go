package commands

import (
	"ubuntuInit/core"
)

func CommonLibs() {
	commands := []string{
		"sudo apt-get install python3-venv python3-wheel libxslt-dev libzip-dev libldap2-dev libsasl2-dev -y",
		"sudo apt-get install python3-setuptools node-less libjpeg-dev gdebi python3-virtualenv -y",
		"sudo apt-get install git python3 python3-pip build-essential wget make vim python3-dev -y",
		"sudo apt-get install libssl-dev zlib1g-dev libbz2-dev libreadline-dev libsqlite3-dev curl llvm libncurses5-dev -y",
		"sudo apt-get install libncursesw5-dev xz-utils tk-dev libffi-dev liblzma-dev python-openssl libxml2-dev -y",
		"sudo apt-get install libxslt1-dev libjpeg62-dev vim-gtk3 pipenv silversearcher-ag exuberant-ctags figlet tox net-tools htop -y",
		"sudo apt-get install nmon -y"}

	core.ExecuteCommands(commands)

}
