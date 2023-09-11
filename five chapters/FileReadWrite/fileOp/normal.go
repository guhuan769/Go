package fileOp

import (
	"io/ioutil"
	"log"
	"path"
)

func ReadWriteFiles() {
	list := getAllFIle(sourceFile)
	for _, fileName := range list {
		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Println(err)
			return
		}
		_, name := path.Split(fileName)
		err = ioutil.WriteFile(destFile+"normal/"+name, bytes, 0644)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
