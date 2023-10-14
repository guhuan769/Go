package fileOp

import (
	"log"
	"os"
	"strings"
)

// 源文件
const sourceFile = "sourceFile/"

// 目标文件
const destFile = "destFile/"

func getAllFIle(dir string) []string {
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return nil
	}
	list := make([]string, 0)
	//range 对切片的操作
	for _, fi := range fs {
		if !fi.IsDir() { //如果不是目录
			fullName := strings.Trim(dir, "/") + "/" + fi.Name()
			list = append(list, fullName)
		}
	}
	return list
}
