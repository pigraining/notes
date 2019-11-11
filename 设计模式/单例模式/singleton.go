package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct {
	Name string
	Age  int
}

func (s *singleton) print() {
	fmt.Printf("%s,age=%d", s.Name, s.Age)
	fmt.Println()
}

var instance *singleton
var n int

//实现方式一 没有加锁 ，会有线程安全问题
func GetInstance1() *singleton {
	if instance == nil {
		n++
		instance = new(singleton)
		instance.Name = "this is a singleton"
		instance.Age = n
		return instance
	}
	fmt.Println("instance is exist")
	return instance
}

var mu sync.Mutex

//实现方式二 加全局锁，保证全局唯一
func GetInstance2() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		n++
		instance = new(singleton)
		instance.Name = "this is a singleton"
		instance.Age = n
		return instance
	}
	fmt.Println("instance is exist")
	return instance
}

//实现方式三，加上局域锁DCL
func GetInstance3() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			n++
			instance = new(singleton)
			instance.Name = "this is a singleton"
			instance.Age = n
			return instance
		}
	}
	fmt.Println("instance is exist")
	return instance
}

var initialized uint32

//使用golang提供的特性
func GetInstance4() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = new(singleton)
		n++
		instance.Name = "this is a singleton"
		instance.Age = n
		atomic.StoreUint32(&initialized, 1)
		return instance
	}
	fmt.Println("instance is exist")
	return instance
}

var once sync.Once

//使用golang特性 once Do
func GetInstance5() *singleton {
	once.Do(func() {
		instance = new(singleton)
		n++
		instance.Name = "this is a singleton"
		instance.Age = n
	})
	return instance
}

//饿汉模式

var insta *singleton

func init() {
	insta = new(singleton)
	insta.Name = "e han singleton"
	insta.Age = 1
}

func GetInstance6() *singleton {
	return insta
}

func main() {
	//单例方式一
	/*
		ins1 := GetInstance1()
		ins1.print()
	*/
	ins4 := GetInstance6()
	ins4.print()
}
