package web

import (
	"CourseManagement/data/course"
	"encoding/json"
	"net/http"
)

func AddCourse(w http.ResponseWriter, req *http.Request) {
	Name := req.FormValue("name")
	teacherIds := req.FormValue("teacher_ids")
	var teacherIdList []int = make([]int, 0) //反序列化到切片
	json.Unmarshal([]byte(teacherIds), &teacherIdList)

	c := &course.Course{
		Name:        Name,
		StuIDList:   make(map[int]int),
		ClassIDList: make(map[int]int),
		LastClassID: 0,
	}
}
