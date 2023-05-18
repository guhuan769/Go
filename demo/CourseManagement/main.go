package main

import (
	_ "CourseManagement/web" //如果只做引用那么前面需要加_
	"net/http"
)

func main() {
	//str := "2022-09-03 21:12:50"
	//t := time.Now()
	//t1, err := common.StrToTime(str, common.DateTime)
	//fmt.Println(t1, err)
	//fmt.Println(common.TimeToStr(t, common.DateTimeZone))

	//启动HTTPserve
	http.ListenAndServe(":8081", nil)
}
