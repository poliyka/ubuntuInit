package commands

import (
	"fmt"
	"ubuntuInit/core"
	"ubuntuInit/enum/InstallOptions"
)

func init() {
	core.InstallAptCollection(InstallOptions.Docker, []string{"ca-certificates", "curl"})
}

func Docker() {
	defer core.Wg.Done()

	fmt.Println(core.StdGreen("Installing Docker"))

	GPGCmd := `
# Add Docker's official GPG key:
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
	`

	commands := []string{
		// remove old docker
		`for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done`,
		// Add Gpg key && repository
		GPGCmd,
		// Install Docker
		"sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin",
		// Add docker group
		"sudo groupadd docker | echo 'docker group already exists'",
		// Add user to docker group
		"sudo usermod -aG docker $USER",
	}

	core.ExecuteCommands(commands)
}
