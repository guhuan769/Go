package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"time"
)

var (
	timeout     int64
	count       int64
	size        int64
	sendCount   int64
	sucessCount int
	failCount   int
	minTs       int64 = math.MaxInt32
	maxTs       int64 = 0
	totalTs     int64
)

type nickType int64

var age nickType

type ICMP struct {
	Type     uint8
	Code     uint8
	CheckSum uint16
	ID       uint16
	SeqNum   uint16
}
type myInterface interface {
	Get()
}

// *号就是我们的指针类型
func (icmp *ICMP) Get() {

}

func ExecGet(myInterface2 myInterface) {
}

func main() {
	getArgs()
	desIp := os.Args[len(os.Args)-1]
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	fmt.Printf("正在 Ping %s [%s] 具有 %d 字节的数据:\n", desIp, conn.RemoteAddr())
	for i := 0; i < 10; i++ {
		sendCount++
		//声明了一个ICMP的指针类型 修改指针 全局生效
		//ICMP 头部
		icmp := &ICMP{
			Type:     8,
			Code:     0,
			CheckSum: 0,
			ID:       uint16(i),
			SeqNum:   uint16(i),
		}

		var buffer bytes.Buffer
		//此处用指针来代替io.write
		binary.Write(&buffer, binary.BigEndian, icmp)
		// 切片的初始化只能通过make 如果直接通过 var data1 []byte 来初始化那么它就为[] 如果通过 make进行初始化 那么就不会为[] 而是有自己的长度
		data := make([]byte, size)
		buffer.Write(data)
		data = buffer.Bytes()
		//如果不填写就是完整的切片 data[:]  data[1:2]
		//data1 := data[1:10] //索引为1 到索引为9 但是得写10
		checkSum, err := checkSum(data)
		if err != nil {
			failCount++
			continue
		}
		data[2] = byte(checkSum >> 8)
		data[3] = byte(checkSum)
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond)) //过期时间
		t1 := time.Now()
		_, err = conn.Write(data)
		if err != nil {
			failCount++
			log.Println(err)
			continue
		}
		buf := make([]byte, 1<<16)
		n, err := conn.Read(buf)
		ts := time.Since(t1).Milliseconds()
		totalTs += ts
		if minTs > ts {
			minTs = ts
		}
		if maxTs < ts {
			maxTs = ts
		}
		if err != nil {
			failCount++
			fmt.Println(err)
			continue
		}
		fmt.Printf("来自 %d.%d.%d.%d 的回复: 字节=%d 时间=%dms TTL=%d\n", buf[12], buf[13], buf[14], buf[15], n-28, ts, buf[8])
		sucessCount++
	}
	fmt.Printf("%s 的 Ping 统计信息:\n    数据包: 已发送 = %d，已接收 = %d，丢失 = %d (%.2f%% 丢失)，\n往返行程的估计时间(以毫秒为单位):\n    最短 = %dms，最长 = %dms，平均 = %dms",
		conn.RemoteAddr(), sendCount, sucessCount, failCount, float64(failCount)/float64(sendCount), minTs, maxTs, totalTs/sendCount)
}

// 校验
func checkSum(data []byte) (uint16, error) {
	length := len(data)
	index := 0
	var sum uint32 //无符号取反很简答 如果是有符号取反比较纠结
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1]) //左移8位
		length -= 2
		index += 2
	}
	if length == 1 {
		sum += uint32(data[index])
	}
	/*
		sum 最大值 :0xfffffff
		高16位 0xffff
		低16位 0xffff
	*/
	hi16 := sum >> 16
	for hi16 != 0 {
		sum = hi16 + uint32(uint16(sum))
		hi16 = sum >> 16
	}
	//golang 取反用^
	return uint16(^sum), nil
}

func getName() *string {
	name := "elon"
	return &name
}

func getArgs() {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时间")
	flag.Int64Var(&count, "n", 4, "请求次数")
	flag.Int64Var(&size, "l", 32, "请求次数")
	flag.Parse()
}
