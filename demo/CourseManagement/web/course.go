package web

import (
	"CourseManagement/data/class"
	"CourseManagement/data/course"
	"CourseManagement/data/stu"
	"CourseManagement/data/teacher"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	c.AddTeacher(teacherIdList...)
	_, err := courseData.Add(c)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	fmt.Fprintf(w, c.String())
}

func GetCourseList(w http.ResponseWriter, req *http.Request) {
	list, err := courseData.Get()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	str, err := json.Marshal(list)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(str))
}

// 课程开班
func NewClassByCourse(w http.ResponseWriter, req *http.Request) {
	courseId, _ := strconv.Atoi(req.FormValue("course_Id"))
	className := req.FormValue("class_name")
	headmasterID, _ := strconv.Atoi(req.FormValue("head_master_id"))
	if courseId == 0 {
		err := errors.New("课程ID为0")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	courselist, err := courseData.Get(courseId)
	if err != nil || len(courselist) == 0 {
		err := errors.New("课程信息未找到")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	courseObj := courselist[0]

	//班级信息
	c := &class.Class{
		Name:         className,
		StuIDList:    make(map[int]int, 0),
		HeadmasterID: headmasterID,
		CourseID:     courseObj.ID,
	}
	classId, err := classData.Add(c)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	courseObj.LastClassID = classId
	courseObj.AddClass(classId)
	err = courseData.Edit(courseObj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	var t *teacher.Teacher
	teacherlist, err := teacherData.Get(headmasterID) //
	if err == nil {
		if len(teacherlist) > 0 {
			if len(teacherlist) > 0 {
				t = teacherlist[0]
			} else {
				t = &teacher.Teacher{}
			}
		} else {
			log.Println(err)
		}
	}
	mp := make(map[string]interface{})
	mp["CourseID"] = courseObj.ID
	mp["CourseName"] = courseObj.Name
	mp["ClassID"] = c.ID
	mp["ClassName"] = c.Name
	mp["HeadMasterID"] = c.HeadmasterID
	mp["HeadMasterName"] = t.Name
	res, _ := json.Marshal(mp)
	fmt.Fprintf(w, string(res))
}

// 课程报名
func SignUpCourse(w http.ResponseWriter, req *http.Request) {
	stuId, _ := strconv.Atoi(req.FormValue("stu_id"))
	courseId, _ := strconv.Atoi(req.FormValue("course_id"))
	if stuId == 0 {
		err := errors.New("输入学员信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	if courseId == 0 {
		err := errors.New("请输入课程信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}

	stuObjList, err := stuData.Get(stuId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}

	var stuObj *stu.Stu
	if len(stuObjList) > 0 {
		stuObj = stuObjList[0]
	}
	if stuObj == nil {
		err := errors.New("未找到学院信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	courseObjList, err := courseData.Get(courseId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	var courseObj *course.Course
	if len(courseObjList) > 0 {
		courseObj = courseObjList[0]
	}
	if courseObj == nil {
		err := errors.New("未找到课程信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	classlist, err := classData.Get(courseObj.LastClassID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	var classObj *class.Class
	if len(classlist) > 0 {
		classObj = classlist[0]

	}
	if classObj == nil {
		err := errors.New("课程需先开班")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	stuObj.AddCourse(courseObj.ID)
	stuObj.AddClass(classObj.ID)
	classObj.AddStu(stuObj.ID)
	courseObj.AddStu(stuObj.ID)
	err = stuData.Edit(stuObj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	err = classData.Edit(classObj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}

	err = classData.Edit(classObj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}

	err = courseData.Edit(courseObj)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}

	mp := map[string]interface{}{
		"StuName":    stuObj.Name,
		"CourseName": courseObj.Name,
		"ClassName":  classObj.Name,
	}
	res, _ := json.Marshal(mp)
	fmt.Fprintf(w, string(res))
}

// 退出课程
func SignOutCourse(w http.ResponseWriter, req *http.Request) {
	stuId, _ := strconv.Atoi(req.FormValue("stu_id"))
	courseId, _ := strconv.Atoi(req.FormValue("course_id"))
	if stuId == 0 {
		err := errors.New("输入学员信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	if courseId == 0 {
		err := errors.New("请输入课程信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}

	stuObjList, err := stuData.Get(stuId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}

	var stuObj *stu.Stu
	if len(stuObjList) > 0 {
		stuObj = stuObjList[0]
	}
	if stuObj == nil {
		err := errors.New("未找到学院信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	courseObjList, err := courseData.Get(courseId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	var courseObj *course.Course
	if len(courseObjList) > 0 {
		courseObj = courseObjList[0]
	}
	if courseObj == nil {
		err := errors.New("未找到课程信息")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	stuObj.DelCourse(courseObj.ID)
	courseObj.DelStu(stuObj.ID)
	classIdList := make([]int, len(courseObj.ClassIDList))
	i := 0
	for _, classId := range courseObj.ClassIDList {
		classIdList[i] = classId
		i++
	}
	classNameList := make([]string, len(courseObj.ClassIDList))
	classObjList, err := classData.Get(classIdList...)
	if err == nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, err.Error())
		return
	}
	j := 0
	for _, classObj := range classObjList {
		stuObj.DelClass(classObj.ID)
		classObj.DelStu(stuObj.ID)
		classData.Edit(classObj)
		classNameList[j] = classObj.Name
		j++
	}
	stuData.Edit(stuObj)
	courseData.Edit(courseObj)
	mp := make(map[string]interface{})
	mp["StuName"] = stuObj.Name
	mp["CourseName"] = courseObj.Name
	mp["class"] = strings.Join(classNameList, ",")
	res, _ := json.Marshal(mp)
	fmt.Fprintf(w, string(res))
}

type CourseResp struct {
	CourseID    int                `json:"courseID"`
	CourseName  string             `json:"courseName"`
	StuCount    int                `json:"stuCount"`
	TeacherList []*teacher.Teacher `json:"teacherList"`
	StuList     []*stu.Stu         `json:"stuList"`
}

func GetHotCourseLisst(w http.ResponseWriter, req *http.Request) {
	list, err := courseData.Get()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	teacherIds := make([]int, 0)
	stuIds := make([]int, 0)
	for _, l := range list {
		for key, _ := range l.TeacherIDList {
			teacherIds = append(teacherIds, key)
		}
		for key, _ := range l.StuIDList {
			stuIds = append(stuIds, key)
		}
	}
	teacherList, err := teacherData.Get(teacherIds...)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	teacherMap := make(map[int]*teacher.Teacher)
	for _, t := range teacherList {
		teacherMap[t.ID] = t
	}

	stuList, err := stuData.Get(stuIds...)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	stuMap := make(map[int]*stu.Stu)
	for _, s := range stuList {
		stuMap[s.ID] = s
	}

	res := make([]*CourseResp, 0)
	for _, l := range list {
		cr := &CourseResp{
			CourseID:    l.ID,
			CourseName:  l.Name,
			StuCount:    len(l.StuIDList),
			TeacherList: make([]*teacher.Teacher, len(l.TeacherIDList)),
			StuList:     make([]*stu.Stu, len(l.StuIDList)),
		}
		i := 0
		for _, tid := range l.TeacherIDList {
			cr.TeacherList[i] = teacherMap[tid]
			i++
		}
		j := 0
		for _, sid := range l.StuIDList {
			cr.StuList[j] = stuMap[sid]
			j++
		}
		res = append(res, cr)

	}
	res = BubbleSortCompare(res, compare)
	str, _ := json.Marshal(res)
	fmt.Fprintf(w, string(str))
}

// 冒泡排序
func BubbleSortCompare(list []*CourseResp, compareFun func(i, j *CourseResp) int) []*CourseResp {
	length := len(list)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if compareFun(list[j], list[j+1]) == 1 {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

// 判断大小
func compare(i, j *CourseResp) int {
	if i.StuCount < j.StuCount {
		return 1
	} else {
		return 0
	}
}
