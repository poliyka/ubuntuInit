package InstallOptions

type Options int

const (
	Ranger Options = iota
	Nvm
	Yarn
	Pyenv
	Fzf
	Docker
	BashColor
	GitAlias
)

var optionsStr = []string{"Ranger", "Nvm", "Yarn", "Pyenv", "Fzf", "Docker", "BashColor", "GitAlias"}

// ======== 枚舉實現 ========

// String - 返回枚舉項的索引值
func (option Options) String() string {
	return optionsStr[option]
}

// Index - 返回枚舉項的字符值
func (option Options) Index() int {
	return int(option)
}

// ======== 輔助函數 ========

// Values 返回枚舉的所有值
func Values() []string {
	return optionsStr
}

// Contains 判斷某值是否存在枚舉值中
func Contains(str string) bool {
	for _, v := range optionsStr {
		if v == str {
			return true
		}
	}
	return false
}
