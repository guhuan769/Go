package web

import (
	"CourseManagement/common"
	"CourseManagement/data/teacher"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func AddTeacher(w http.ResponseWriter, req *http.Request) {
	Name := req.FormValue("name")
	GenderStr := req.FormValue("gender")
	BirthdayStr := req.FormValue("birthday")
	gender, _ := strconv.Atoi(GenderStr)
	t := &teacher.Teacher{
		Name:   Name,
		Gender: common.Gender(gender),
	}
	birthday, err := common.StrToTime(BirthdayStr)
	if err == nil {
		t.Birthday = birthday
	}
	_, err = teacherData.Add(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w, "%s", err.Error())
		return
	}
	fmt.Fprintf(w, t.String())
}

func GetTeacherList(w http.ResponseWriter, req *http.Request) {
	list, err := teacherData.Get()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	//序列化
	bytes, err := json.Marshal(list)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, string(bytes))
}
