package teacher

import (
	"CourseManagement/common"
	"CourseManagement/data"
	"time"
)

type Teacher struct {
	ID       int
	Name     string
	Gender   common.Gender
	Birthday time.Time
}

type TeacherData struct {
	tableName data.TableName
}

func NewCourseData() *TeacherData {
	return &TeacherData{
		tableName: data.Class,
	}
}

func (d *TeacherData) Add(teacher *Teacher) (int, error) {
	id, err := data.Add(d.tableName, teacher)
	teacher.ID = id
	data.ShowAllData(d.tableName)
	return teacher.ID, err
}

func (d *TeacherData) Edit(teacher *Teacher) error {
	err := data.Edit(d.tableName, teacher.ID, teacher)
	data.Edit(d.tableName, teacher.ID, teacher)
	return err
}

func (d *TeacherData) Del(id int) error {
	err := data.Delete(d.tableName, id)
	return err
}
func (d *TeacherData) Get(id ...int) ([]*Teacher, error) {
	//第二个参数 代表着能有几个切片元素
	list := make([]*Teacher, 0) //不知道有多少 初始化为0
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
			teacher, ok := v.(*Teacher)
			if !ok {
				continue
			}
			teacher.ID = i
			list = append(list, teacher) ///...打散集合
		}
	} else {
		//如果是一个集合那么返回的就是 key value
		for k, v := range mp {
			teacher, ok := v.(*Teacher)
			if !ok {
				continue
			}
			teacher.ID = k
			list = append(list, teacher) //使用append 这种效率比较低
		}
	}
	return list, err
}
