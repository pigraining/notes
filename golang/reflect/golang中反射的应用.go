package main

import (
	"fmt"
	"reflect"
)

//获取接口的type 和 值
func test1(str interface{}) {
	fmt.Println(reflect.ValueOf(str))
	fmt.Println(reflect.TypeOf(str))
}

//强转一个接口到其它类型值
func test2(str interface{}) {
	fmt.Println("str type :", reflect.ValueOf(str))
	fmt.Println("str type :", reflect.TypeOf(str))

	str1 := reflect.ValueOf(str)
	fmt.Println("str1 type :", reflect.ValueOf(str1))
	fmt.Println("str1 type :", reflect.TypeOf(str1))

	str2 := str1.Interface().(string)
	fmt.Println("str2 type :", reflect.ValueOf(str2))
	fmt.Println("str2 type :", reflect.TypeOf(str2))
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Print(name string, age int) {
	fmt.Printf("person name = %s,age = %d", p.Name, p.Age)
}

type Student struct {
	Id     string
	Person Person
}

func (s Student) Print(stu Student) {
	fmt.Printf("student id = %s,student name = %s,student age = %d", stu.Id, stu.Person.Name, stu.Person.Age)
}

//解析传过来的结构体
func test3(input interface{}) {
	val := reflect.ValueOf(input)
	type1 := reflect.TypeOf(input)
	if type1.Kind() != reflect.Struct {
		fmt.Println("this is err")
	}
	for i := 0; i < type1.NumField(); i++ {
		field := type1.Field(i)
		value := val.Field(i).Interface()
		if field.Type.Kind() == reflect.Struct {
			test3(value)
		}
		fmt.Printf("name = %s,type = %s,value = %v", field.Name, field.Type, value)
		fmt.Println()
	}
	//解析结构体中方法的名称和类型
	for i := 0; i < type1.NumMethod(); i++ {
		m := type1.Method(i)
		fmt.Printf("func name = %s,func type = %s", m.Name, m.Type)
		fmt.Println()
	}
}

type Docker struct {
	Person
	Department string
}

//直接解析结构体中套结构体
//或者是结构体中含有匿名结构
func test4(input interface{}) {

	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	field, _ := t.FieldByName("Name")
	value := v.FieldByIndex([]int{0, 0})
	fmt.Printf("name = %s,type = %s,value = %v", field.Name, field.Type, value)
	fmt.Println()
}

//通过反射修改字段
func test5() {
	in := 100
	val := reflect.ValueOf(&in)
	val.Elem().SetInt(200)
	fmt.Println(in)
}

//通过反射修改struct中的字段
func test6(input interface{}) {
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("input 为指针类型，且不能修改")
		return
	} else {
		v = v.Elem()
	}

	feild := v.FieldByName("Name")
	if !feild.IsValid() {
		fmt.Println("没有这个值")
		return
	}
	if feild.Kind() == reflect.String {
		feild.SetString("zhu yu ning")
	}
	fmt.Println(input)
}

//通过反射动态调用struct方法
func test7(input interface{}) {
	m := Student{
		Id:     "111",
		Person: Person{"zzzzzz", 18},
	}

	v := reflect.ValueOf(input)
	mv := v.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf(m)}
	mv.Call(args)
}

func main() {
	str := "this is a string"
	fmt.Println("this is test1")
	test1(str)
	fmt.Println("----------------------------------------------")

	fmt.Println("this is test2")
	test2(str)
	fmt.Println("----------------------------------------------")

	fmt.Println("this is test3")
	stu := Student{
		Id:     "002",
		Person: Person{Name: "wei", Age: 21},
	}
	test3(stu)
	fmt.Println("----------------------------------------------")

	fmt.Println("this is test4")
	doc := Docker{
		Person:     Person{Name: "wei", Age: 21},
		Department: "005部门",
	}
	test4(doc)
	fmt.Println("----------------------------------------------")

	fmt.Println("this is test5")
	test5()
	fmt.Println("----------------------------------------------")

	fmt.Println("这是测试6")
	doc1 := &Docker{
		Person:     Person{Name: "wei", Age: 21},
		Department: "005部门",
	}
	test6(doc1)
	fmt.Println("----------------------------------------------")

	fmt.Println("this is test7")
	stu1 := &Student{}
	test7(stu1)
	fmt.Println()
	fmt.Println("----------------------------------------------")
}
