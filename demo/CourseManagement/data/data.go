package data

import (
	"encoding/json"
	"fmt"
	"sync"
)

type TableName string //做别名

// 创建常量
const (
	Stu     TableName = "stu"
	Course  TableName = "course"
	Teacher TableName = "teacher"
	Class   TableName = "class"
)

type item struct {
	// 在golang里面interface有点类似与其它语言中的泛型
	data  map[int]interface{}
	count int
	maxId int
	//
	sync.RWMutex //采用值类型 不用作初始化
}

type it struct {
}

// 结构体的名称作为变量名
func (i *it) Get() {

}

var dataStorage map[TableName]*item

// 首先执行init方法
func init() {
	//mp := make(map[TableName]*item)
	//mp[Stu] = &item{}
	dataStorage = map[TableName]*item{
		Stu: &item{
			data:  make(map[int]interface{}, 0),
			count: 0,
			maxId: 0,
		},
		Teacher: &item{
			data:  make(map[int]interface{}, 0),
			count: 0,
			maxId: 0,
		},
		Class: &item{
			data:  make(map[int]interface{}, 0),
			count: 0,
			maxId: 0,
		},
		Course: &item{
			data:  make(map[int]interface{}, 0),
			count: 0,
			maxId: 0,
		},
	}
}

func Add(tableName TableName, obj interface{}) (id int, err error) {
	//使用读写锁
	dataStorage[tableName].Lock() //加锁
	//defer 在方法结束后执行操作
	defer dataStorage[tableName].Unlock() //解锁
	//dataStorage[tableName].RLock() //读锁
	////上面方法开始加锁结束就解锁
	//defer dataStorage[tableName].RUnlock()
	id = dataStorage[tableName].maxId + 1
	dataStorage[tableName].count += 1
	dataStorage[tableName].maxId = id
	dataStorage[tableName].data[id] = obj
	return
}

func Edit(tableName TableName, id int, obj interface{}) (err error) {
	//使用读写锁
	dataStorage[tableName].Lock() //加锁
	//defer 在方法结束后执行操作
	defer dataStorage[tableName].Unlock() //解锁
	dataStorage[tableName].data[id] = obj
	return
}

func Delete(tableName TableName, id int) (err error) {
	dataStorage[tableName].Lock()         //加锁
	defer dataStorage[tableName].Unlock() //解锁
	delete(dataStorage[tableName].data, id)
	dataStorage[tableName].count -= 1
	return
}
func Get(tableName TableName, id ...int) (mp map[int]interface{}, err error) {
	dataStorage[tableName].RLock() //读锁
	defer dataStorage[tableName].RUnlock()
	//初始化mp
	count := 100
	if len(id) > 0 {
		mp = make(map[int]interface{}, len(id))
		//golnag中如果定义了变量必须使用 否则error 否则就用_接收进行忽略
		for _, i := range id {
			item, ok := dataStorage[tableName].data[i]
			if ok {
				mp[i] = item
			}
		}
	} else {
		mp = make(map[int]interface{}, count)
		for i, item := range dataStorage[tableName].data {
			if i >= count {
				break
			}
			mp[i] = item
		}
	}
	return
}

func ShowAllData(tableName TableName) {
	item := dataStorage[tableName]
	str, _ := json.Marshal(item.data)
	fmt.Println("--------------------------start")
	fmt.Println("tableName:", tableName)
	fmt.Println("count:", item.count)
	fmt.Println("maxId:", item.maxId)
	fmt.Println("data:", string(str))
	fmt.Println("----------------------------End")
}
