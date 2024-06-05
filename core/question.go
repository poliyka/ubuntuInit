package core

import (
	"github.com/AlecAivazis/survey/v2"
)

// 定義問題
var QS = []*survey.Question{
	{
		Name: "UpdateAndUpgrade",
		Prompt: &survey.Confirm{
			Message: "Did you want to run 'sudo apt update && sudo apt upgrade'?",
		},
	},
	{
		Name: "CommonLibs",
		Prompt: &survey.Confirm{
			Message: "Did you want to common lib?",
		},
	},
	{
		Name: "InstallChoices",
		Prompt: &survey.MultiSelect{
			Message: "Select the packages you want to install:",
			Options: []string{"Ranger", "Nvm", "Yarn", "Pyenv", "Fzf", "BashColor", "GitAlias"},
		},
	},
	// {
	// 	Name: "ListPath",
	// 	Prompt: &survey.Input{
	// 		Message: "Enter the path to the list file:",
	// 	},
	// 	Validate: survey.ComposeValidators(survey.Required, ValidatePath),
	// },
}
