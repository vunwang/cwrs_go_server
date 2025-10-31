package cwrs_utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// 从指定路径filePath读取一个文件的内容
func GetFileContent(filePath string) string {
	bytes, err := os.ReadFile(filePath)
	if nil != err {
		fmt.Errorf(" %s getFileBase64 error: %v", filePath, err)
		return ""
	}
	return string(bytes)
}

// 获取项目根目录
func GetProjectRoot() string {
	var root string
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Unable to get current file path")
	}

	// 获取当前文件的绝对路径
	absPath, err := filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
	}

	// 计算根目录
	root = filepath.Dir(absPath)
	for i := 0; i < 3; i++ { // 假设main.go位于cmd/main/main.go，向上回溯三层找到根目录
		root = filepath.Dir(root)
	}

	return root
}
