package class

import (
	"CourseManagement/data"
	"encoding/json"
	"sync"
)

type Class struct {
	ID           int
	Name         string
	StuIDList    map[int]int
	HeadmasterID int
	sync.RWMutex
	CourseID int
}

type ClassData struct {
	tableName data.TableName
}

func NewCourseData() *ClassData {
	return &ClassData{
		tableName: data.Class,
	}
}

func (d *ClassData) Add(class *Class) (int, error) {
	id, err := data.Add(d.tableName, class)
	class.ID = id
	data.ShowAllData(d.tableName)
	return class.ID, err
}

func (d *ClassData) Edit(class *Class) error {
	err := data.Edit(d.tableName, class.ID, class)
	data.Edit(d.tableName, class.ID, class)
	return err
}

func (d *ClassData) Del(id int) error {
	err := data.Delete(d.tableName, id)
	return err
}
func (d *ClassData) Get(id ...int) ([]*Class, error) {
	//第二个参数 代表着能有几个切片元素
	list := make([]*Class, 0) //不知道有多少 初始化为0
	mp, err := data.Get(d.tableName, id...)
	if err != nil {
		return nil, err
	}
	if len(id) > 0 {
		// _索引 i id值
		for _, i := range id {
			v, ok := mp[i]
			if !ok {
				continue
			}
			//i1, ok := v.(int)
			class, ok := v.(*Class)
			if !ok {
				continue
			}
			class.ID = i
			list = append(list, class) ///...打散集合
		}
	} else {
		//如果是一个集合那么返回的就是 key value
		for k, v := range mp {
			class, ok := v.(*Class)
			if !ok {
				continue
			}
			class.ID = k
			list = append(list, class) //使用append 这种效率比较低
		}
	}
	return list, err
}

func (c *Class) AddStu(classId ...int) {
	c.Lock()
	defer c.Unlock()
	if c.StuIDList == nil {
		c.StuIDList = make(map[int]int) //初值
	}
	for _, id := range classId {
		c.StuIDList[id] = id
	}
}

func (c *Class) DelStu(classId ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range classId {
		delete(c.StuIDList, id)
		//c.StuIDList[id] = id
	}
}
func (c *Class) String() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}
