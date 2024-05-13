package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"log"
	"os/exec"
)

// 定義一個結構來存儲用戶的回答
type cliConfig struct {
	ProjectName string
	Frontend    string
	Backend     bool
}

func validatePath(val interface{}) error {
	str, ok := val.(string)
	if str == "" || !ok {
		return fmt.Errorf("invalid input")
	}
	cmd := exec.Command("ls", str)

	output, err := cmd.Output()

	if err != nil {
		log.Fatal(fmt.Errorf("failed to execute command: '%w'", err))
	}

	fmt.Println(string(output))

	return nil
}

func main() {
	// 創建一個新的配置實例
	config := &cliConfig{}

	// 定義問題
	var qs = []*survey.Question{
		{
			Name: "ProjectName",
			Prompt: &survey.Input{
				Message: "What is the name of the project?",
			},
			Validate: survey.Required,
		},
		{
			Name: "Frontend",
			Prompt: &survey.Select{
				Message: "Choose a frontend framework:",
				Options: []string{"Vue.js", "React", "Angular", "None"},
			},
		},
		{
			Name: "Backend",
			Prompt: &survey.Confirm{
				Message: "Do you need a backend setup?",
			},
		},
		{
			Name: "ListPath",
			Prompt: &survey.Input{
				Message: "Enter the path to the list file:",
			},
			Validate: survey.ComposeValidators(survey.Required, validatePath),
		},
	}

	// 啟動問題並存儲回答
	err := survey.Ask(qs, config)
	if err != nil {
		fmt.Println("Failed to execute survey:", err)
		return
	}

	// 輸出收集到的配置信息
	fmt.Printf("Project Name: %s\n", config.ProjectName)
	fmt.Printf("Frontend Framework: %s\n", config.Frontend)
	fmt.Printf("Include Backend: %t\n", config.Backend)
}
