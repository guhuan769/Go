package fileOp

import (
	"io"
	"log"
	"os"
	"path"
)

func OneSideReadWriteTodest() {
	list := getAllFIle(sourceFile)
	for _, l := range list {
		_, name := path.Split(l)
		destFileName := destFile + "oneSide/" + name
		OneSideReadWrite(l, destFileName)
	}
}

//边读边写

func OneSideReadWrite(srcName, destName string) {
	src, _ := os.Open(srcName)
	defer src.Close()
	dst, _ := os.OpenFile(destName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer dst.Close()
	//创建一个byte切片
	buf := make([]byte, 1024)
	for {
		//n表示读了多少字节
		n, err := src.Read(buf)
		if err != nil && err != io.EOF {
			log.Println(err)
			return
		}
		if n == 0 {
			break
		}
		dst.Write(buf[:n])
	}
}
