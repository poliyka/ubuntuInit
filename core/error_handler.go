package core

import (
	"os"
	"os/exec"
)

// 驗證 exec.Cmd 的輸出，如果出現錯誤，則將錯誤記錄到日誌中, 如無錯誤，則返回輸出
func HandleError(cmd *exec.Cmd) error {
	// 將標準輸出和錯誤輸出設置為當前進程的輸出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		// 處理錯誤
		os.Stderr.WriteString(err.Error())
	}
	return err
}
