package stu

import (
	"CourseManagement/common"
	"CourseManagement/data"
	"encoding/json"
	"sync"
	"time"
)

type Stu struct {
	ID           int
	Name         string        //姓名
	Gender       common.Gender //性别
	Birthday     time.Time     //生日日期
	CourseIDList map[int]int   //int切片
	ClassIDList  map[int]int   //班级信息
	sync.RWMutex               //加一个读写锁
}

type StuData struct {
	tableName data.TableName
}

func NewStuData() *StuData {
	return &StuData{
		tableName: data.Stu,
	}
}

func (d *StuData) Add(stu *Stu) (int, error) {
	id, err := data.Add(d.tableName, stu)
	stu.ID = id
	data.ShowAllData(d.tableName)
	return stu.ID, err
}

func (d *StuData) Edit(stu *Stu) error {
	err := data.Edit(d.tableName, stu.ID, stu)
	data.Edit(d.tableName, stu.ID, stu)
	return err
}

func (d *StuData) Del(id int) error {
	err := data.Delete(d.tableName, id)
	return err
}
func (d *StuData) Get(id ...int) ([]*Stu, error) {
	//第二个参数 代表着能有几个切片元素
	list := make([]*Stu, 0) //不知道有多少 初始化为0
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
			stu, ok := v.(*Stu)
			if !ok {
				continue
			}
			stu.ID = i
			list = append(list, stu) ///...打散集合
		}
	} else {
		//如果是一个集合那么返回的就是 key value
		for k, v := range mp {
			stu, ok := v.(*Stu)
			if !ok {
				continue
			}
			stu.ID = k
			list = append(list, stu) //使用append 这种效率比较低
		}
	}
	return list, err
}

func (c *Stu) AddCourse(courseId ...int) {
	c.Lock()
	defer c.Unlock()
	if c.CourseIDList == nil {
		c.CourseIDList = make(map[int]int) //初值
	}
	for _, id := range courseId {
		c.CourseIDList[id] = id
	}
}

func (c *Stu) DelCourse(courseId ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range courseId {
		delete(c.CourseIDList, id)
		//c.StuIDList[id] = id
	}
}

func (c *Stu) AddClass(classId ...int) {
	c.Lock()
	defer c.Unlock()
	if c.ClassIDList == nil {
		c.ClassIDList = make(map[int]int) //初值
	}
	for _, id := range classId {
		c.ClassIDList[id] = id
	}
}

func (c *Stu) DelClass(classId ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range classId {
		delete(c.ClassIDList, id)
		//c.StuIDList[id] = id
	}
}

func (s *Stu) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

//func (d StuData) Add(stu *Stu) (int, error) {
//
//	return 0, nil
//}
