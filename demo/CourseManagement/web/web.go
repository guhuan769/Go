package web

import (
	"CourseManagement/data/class"
	"CourseManagement/data/course"
	"CourseManagement/data/stu"
	"CourseManagement/data/teacher"
)

var (
	// 1
	stuData     = stu.NewStuData()
	teacherData = teacher.NewTeacherData()
	classData   = class.NewCourseData()
	courseData  = course.NewCourseData()
)
