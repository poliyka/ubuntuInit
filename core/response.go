package core

// 定義一個結構來存儲用戶的回答
type Response struct {
	TerminalType     string
	UpdateAndUpgrade bool
	CommonLibs       bool
	InstallChoices   []string
}
