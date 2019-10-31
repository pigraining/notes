# golang基准测试

连接 https://blog.csdn.net/hjmnasdkl/article/details/81304329 

 基准测试以Benchmark开头，接收一个指针型参数（*testing.B） 

```go
package main
 
import (
	"fmt"
	"strconv"
	"testing"
)
 
func BenchmarkA(b *testing.B) {
	number := 10
	b.ResetTimer()
 
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}
 
func BenchmarkB(b *testing.B) {
	number := 10
	b.ResetTimer()
 
	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}				
```

 b.ResetTimer() 代表从此处计算时间，如果前期准备时间过长，可以使用

运行测试   go test -bench=. -benchmem     (其中的 . 可以换成具体的函数名称，匹配正则) 

也可以 -benchmem = 3（3为时间 s） 来保证运行基准测试的运行时间为3s，默认一次基准测试的时间为1s

```go
$ go test -v -bench=. -benchmem
goos: windows
goarch: amd64
BenchmarkA-4    20000000               122 ns/op              16 B/op          2 allocs/op
BenchmarkB-4    100000000               11.0 ns/op             0 B/op          0 allocs/op
PASS
ok      _/C_/Users/pxlol/Desktop/demo   3.791s
```

-4  代表机器为4核

 122ns/op 表示每次操作耗时122纳秒， 16B表示每次操作用了16字节，2 allocs表示每次操作分配内存2次。 

