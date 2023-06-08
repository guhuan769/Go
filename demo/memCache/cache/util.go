package cache

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

// 2 的10次方意思就是10个2相乘
const (
	B  = 1 << (iota * 10) // iota 是第0个值 0*10 等于0 左移0位还是1
	kb                    //1<<(1*10)
	MB
	GB
	TB
	PB
)
const defaultNum = 100

func ParseSize(size string) (int64, string) {
	//定义政策表达式
	// size 1KB 100KB 1MB 2MB 1GB
	re, _ := regexp.Compile("0-9+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	log.Println("我来了", size, unit, num)
	switch unit { //switch case 默认就会break
	case "1B":
		log.Println("我来了", size)
		byteNum = num
	case "KB":
		byteNum = num * kb
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
	}
	if num == 0 {
		log.Println("PaseSize 仅支持 B、KB、MB、GB、TB、PB")
		num = defaultNum
		byteNum = num * MB
		unit = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unit
	return byteNum, sizeStr
}
func GetValSize(val interface{}) int64 {
	//TODO
	fmt.Println(unsafe.Sizeof(val))
	return 0
}
