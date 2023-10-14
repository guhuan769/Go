package yaml

import (
	"fmt"
	"io/ioutil"
)

func LoadPath(filePath string) {
	data, _ := ioutil.ReadFile(filePath)
	content := string(data)
	fmt.Printf(content)
}
