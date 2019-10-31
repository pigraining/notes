package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Id   string
	Addr []string
}

//使用  xxx == Person{} 来判断结构体是否为空
func test1() {
	var person Person
	per := Person{}
	if person == per {
		fmt.Println(person, "is null")
	}
}

//test1这个方法有一个问题，不能判断存在内嵌结构体的情况
/*func test2() {
	var stu Student
	if stu == (Student{}) {
		fmt.Println(stu, "is null")
	}
}*/

//方法2  使用reflect来进行判断比较
func test3() {
	var stu Student
	if reflect.DeepEqual(stu, (Student{})) {
		fmt.Println("struct is null")
	}
	stu1 := Student{
		Id:   "111",
		Addr: []string{"wo", "ning"},
	}
	if !reflect.DeepEqual(stu1, (Student{})) {
		fmt.Println("struct is not null")
	}

}

func main() {
	test1()
	//错误示范
	//	test2()

	test3()
}
