package fileOp

import (
	"io"
	"os"
)

func CopyFIle(srcName, destName string) (int64, error) {
	src, _ := os.Open(srcName)
	defer src.Close() //最后执行
	dst, _ := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY, 0644)
	defer dst.Close()
	return io.Copy(dst, src)
}
