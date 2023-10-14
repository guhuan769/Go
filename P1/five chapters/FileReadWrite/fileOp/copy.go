package fileOp

import (
	"io"
	"os"
	"path"
)

func CopyAllFileToDest() {
	list := getAllFIle(sourceFile)
	for _, l := range list {
		_, name := path.Split(l)
		destFileName := destFile + "copy/" + name
		CopyFIle(l, destFileName)
	}
}

func CopyFIle(srcName, destName string) (int64, error) {
	//fmt.Println(srcName, destName)
	//打开源文件
	src, _ := os.Open(srcName)
	defer src.Close() //最后执行
	dst, _ := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
	defer dst.Close()
	return io.Copy(dst, src)
}
