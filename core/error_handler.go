package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// 驗證 exec.Cmd 的輸出，如果出現錯誤，則將錯誤記錄到日誌中, 如無錯誤，則返回輸出
func HandleError(cmd *exec.Cmd) error {
	// 將標準輸出和錯誤輸出設置為當前進程的輸出
	var stderrBuf bytes.Buffer
	cmd.Stdout = os.Stdout

	// 透過 stderrBuf 紀錄錯誤訊息
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	if err != nil {

		// 讀取 stderrBuf 中的錯誤訊息
		stderrStr := stderrBuf.String()

		// 打印錯誤内容
		os.Stderr.WriteString(StdRed(stderrStr) + "\n")

		// 處理錯誤
		os.Stderr.WriteString(StdRed(err.Error()) + "\n")
		// 記錄錯誤
		os.Stderr.WriteString(StdRed("Failed to execute command: "+cmd.String()) + "\n")
		// 紀錄執行錯誤的位置
		pc, file, line, ok := runtime.Caller(1)
		if ok {
			fn := runtime.FuncForPC(pc)
			os.Stderr.WriteString(StdRed(fmt.Sprintf("Error occurred in %s \n", fn.Name())))
			os.Stderr.WriteString(StdRed(fmt.Sprintf("At %s:%d\n", file, line)))
		}

	}
	return err
}
