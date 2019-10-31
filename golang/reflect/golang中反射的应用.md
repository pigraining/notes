### golang中反射的应用

链接：https://juejin.im/post/5a75a4fb5188257a82110544来源：掘金著作权归作者所有。

##### 1.interface 和反射

在讲反射之前，先来看看Golang关于类型设计的一些原则

- 变量包括（type, value）两部分
  - 理解这一点就知道为什么nil != nil了
- `type` 包括 static type和concrete type. 简单来说 static type是你在编码是看见的类型(如int、string)，concrete type是runtime系统看见的类型
- 类型断言能否成功，取决于变量的concrete type，而不是static type. 因此，一个 reader变量如果它的concrete type也实现了write方法的话，它也可以被类型断言为writer.

接下来要讲的反射，就是建立在类型之上的，Golang的指定类型的变量的类型是静态的（也就是指定int、string这些的变量，它的type是static type），在创建变量的时候就已经确定，反射主要与Golang的interface类型相关（它的type是concrete type），只有interface类型才有反射一说。

在Golang的实现中，每个interface变量都有一个对应pair，pair中记录了实际变量的值和类型:

```go
(value, type)
```

value是实际变量值，type是实际变量的类型。一个interface{}类型的变量包含了2个指针，一个指针指向值的类型【对应concrete type】，另外一个指针指向实际的值【对应value】。

## Golang的反射reflect

### reflect的基本功能TypeOf和ValueOf

既然反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制。那么在Golang的reflect反射包中有什么样的方式可以让我们直接获取到变量内部的信息呢？ 它提供了两种类型（或者说两个方法）让我们可以很容易的访问接口变量内容，分别是reflect.ValueOf() 和 reflect.TypeOf()，看看官方的解释

```go
// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i.  ValueOf(nil) returns the zero 
func ValueOf(i interface{}) Value {...}

翻译一下：ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0


// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
func TypeOf(i interface{}) Type {...}
翻译一下：TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
```

###### 代码准备工作(准备好所有用到的结构体)

```go
package main

import (
	"fmt"
	"reflect"
)

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

type Docker struct {
	Person
	Department string
}
```

###### 通过reflect.ValueOf 和relect.TypeOf 获取interface的类型和具体值

```go
//获取接口的type 和 值
func test1(str interface{}) {
	fmt.Println(reflect.ValueOf(str))
	fmt.Println(reflect.TypeOf(str))
}
	str := "this is a string"
	fmt.Println("this is test1")
	test1(str)
	fmt.Println("----------------------------------------------")
```



```go
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
	fmt.Println("this is test2")
	test2(str)
	fmt.Println("----------------------------------------------")

```



```go
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

	fmt.Println("this is test3")
	stu := Student{
		Id:     "002",
		Person: Person{Name: "wei", Age: 21},
	}
	test3(stu)
	fmt.Println("----------------------------------------------")
```



```go
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
	fmt.Println("this is test4")
	doc := Docker{
		Person:     Person{Name: "wei", Age: 21},
		Department: "005部门",
	}
	test4(doc)
	fmt.Println("----------------------------------------------")
```

```go
//通过反射修改字段
func test5() {
	in := 100
	val := reflect.ValueOf(&in)
	val.Elem().SetInt(200)
	fmt.Println(in)
}
fmt.Println("this is test5")
	test5()
	fmt.Println("----------------------------------------------")
```



```go
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
fmt.Println("这是测试6")
	doc1 := &Docker{
		Person:     Person{Name: "wei", Age: 21},
		Department: "005部门",
	}
	test6(doc1)
	fmt.Println("----------------------------------------------")
```

call方法动态调用需要传入一个Value的切片

```go
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
	fmt.Println("this is test7")
	stu1 := &Student{}
	test7(stu1)
	fmt.Println()
	fmt.Println("----------------------------------------------")
```

