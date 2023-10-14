package course

import (
	"CourseManagement/data"
	"encoding/json"
	"sync"
)

type Course struct {
	ID            int
	Name          string
	TeacherIDList map[int]int
	StuIDList     map[int]int
	ClassIDList   map[int]int
	LastClassID   int
	sync.RWMutex  //枷锁
}

type CourseData struct {
	tableName data.TableName
}

func NewCourseData() *CourseData {
	return &CourseData{
		tableName: data.Course,
	}
}

func (d *CourseData) Add(course *Course) (int, error) {
	id, err := data.Add(d.tableName, course)
	course.ID = id
	data.ShowAllData(d.tableName)
	return course.ID, err
}

func (d *CourseData) Edit(course *Course) error {
	err := data.Edit(d.tableName, course.ID, course)
	data.Edit(d.tableName, course.ID, course)
	return err
}

func (d *CourseData) Del(id int) error {
	err := data.Delete(d.tableName, id)
	return err
}
func (d *CourseData) Get(id ...int) ([]*Course, error) {
	//第二个参数 代表着能有几个切片元素
	list := make([]*Course, 0) //不知道有多少 初始化为0
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
			course, ok := v.(*Course)
			if !ok {
				continue
			}
			course.ID = i
			list = append(list, course) ///...打散集合
		}
	} else {
		//如果是一个集合那么返回的就是 key value
		for k, v := range mp {
			course, ok := v.(*Course)
			if !ok {
				continue
			}
			course.ID = k
			list = append(list, course) //使用append 这种效率比较低
		}
	}
	return list, err
}

func (c *Course) AddStu(stuId ...int) {
	c.Lock()
	defer c.Unlock()
	if c.StuIDList == nil {
		c.StuIDList = make(map[int]int) //初值
	}
	for _, id := range stuId {
		c.StuIDList[id] = id
	}
}

func (c *Course) DelStu(stuId ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range stuId {
		delete(c.StuIDList, id)
		//c.StuIDList[id] = id
	}
}

func (c *Course) AddClass(classId ...int) {
	c.Lock()
	defer c.Unlock()
	if c.ClassIDList == nil {
		c.ClassIDList = make(map[int]int) //初值
	}
	for _, id := range classId {
		c.ClassIDList[id] = id
	}
}

func (c *Course) DelClass(classId ...int) {
	c.Lock()
	defer c.Unlock()
	for _, id := range classId {
		delete(c.ClassIDList, id)
		//c.StuIDList[id] = id
	}
}

func (c *Course) AddTeacher(teacherId ...int) {
	c.Lock()
	defer c.Unlock()
	if c.TeacherIDList == nil {
		c.TeacherIDList = make(map[int]int) //初值
	}
	for _, id := range teacherId {
		c.TeacherIDList[id] = id
	}
}
func (c *Course) String() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}
