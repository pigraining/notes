package main

import (
	"fmt"
	"unsafe"
)

type Emtyp struct {
	num string
	len int
}

type Emtyp1 struct {
}

func test() {
	actionMap := make(map[string][]int)
	actionMap["1"] = []int{1, 2, 3}
	list := actionMap["1"]
	actionMap["1"] = append(list, 12)
	fmt.Println(actionMap)

}
func test2() {
	n := &Emtyp{}
	fmt.Println(unsafe.Sizeof(n))

	m := &Emtyp1{}
	fmt.Println(unsafe.Sizeof(m))

	var f Emtyp1
	fmt.Println(unsafe.Sizeof(f))
	v := &Emtyp{
		num:"10000000000000000000000000000000000000000000000000000000000000000000000000000000",
		len:10,
	}
	fmt.Println(unsafe.Sizeof(v))

	if n == v {
		fmt.Println(true)
	}
	fmt.Println(false)

}
func main() {
	test2()
}
