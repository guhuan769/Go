package fileOp

import (
	"fmt"
	"io/ioutil"
	"path"
	"sync"
)

func ReadWriteFilesByGoRoutine() {
	//用来计数的
	wg := sync.WaitGroup{}
	list := getAllFIle(sourceFile)
	for _, fileName := range list {
		//子携程
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			fmt.Println(fileName)
			bytes, _ := ioutil.ReadFile(fileName)
			_, name := path.Split(fileName)
			ioutil.WriteFile(destFile+"goroutine/"+name, bytes, 0644)
		}(fileName)
	}
	wg.Wait() //等到wg里面的携程等于0 开始一个+1 完成一个-1
}
