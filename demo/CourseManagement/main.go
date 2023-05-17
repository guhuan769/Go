package main

import (
	"CourseManagement/data/stu"
)

type s struct {
	ID   int
	Name string
	Age  int
}

func main() {
	d := stu.NewStuData()
	stuobj := &stu.Stu{
		Name: "elon1",
	}
	d.Add(stuobj)

	//fmt.Println("hello golang")
	//s1 := &s{
	//	Name: "elon",
	//	Age:  18,
	//}
	//s2 := &s{
	//	Name: "elon",
	//	Age:  18,
	//}
	//s3 := &s{
	//	Name: "elon",
	//	Age:  18,
	//}
	//s4 := &s{
	//	Name: "elon",
	//	Age:  18,
	//}
	//data.Add(data.Course, s1)
	//data.Add(data.Course, s2)
	//data.Add(data.Course, s3)
	//data.Add(data.Course, s4)
	//
	//data.ShowAllData(data.Course)
	//s1.ID = 3
	//s1.Name = "李四"
	//data.Edit(data.Course, 1, s1)
	//data.ShowAllData(data.Course)
	//data.Delete(data.Course, 3)
	//data.ShowAllData(data.Course)
	//res, _ := data.Get(data.Course)
	//fmt.Println(res)
	//fmt.Println("hello golang")

}
