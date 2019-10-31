# 基准测试结合pprof(基础)

 链接https://my.oschina.net/solate/blog/3034188 

golang中基准测试可以结合pprof来生成图形化界面，可以更直观的看接口使用的时间、内存、堆栈等等信息。

```go
package bench
import "testing"
func Fib(n int) int {
    if n < 2 {
      return n
    }
    return Fib(n-1) + Fib(n-2)
}
func BenchmarkFib10(b *testing.B) {
    // run the Fib function b.N times
    for n := 0; n < b.N; n++ {
      Fib(10)
    }
}
```

一个基准测试的实例

基于pprof运行

 go test -bench=. -benchmem -cpuprofile profile.out //只查看cpu的运行命令

 go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out  //查看内存和cpu的运行

go test -v -bench= .  -benchmem  //全部包含

##### 具体分析进程

go tool pprof profile.out    

```go
go tool pprof profile.out
File: bench.test
Type: cpu
Time: Apr 5, 2018 at 4:27pm (EDT)
Duration: 2s, Total samples = 1.85s (92.40%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1.85s, 100% of 1.85s total
      flat  flat%   sum%        cum   cum%
     1.85s   100%   100%      1.85s   100%  bench.Fib
         0     0%   100%      1.85s   100%  bench.BenchmarkFib10
         0     0%   100%      1.85s   100%  testing.(*B).launch
         0     0%   100%      1.85s   100%  testing.(*B).runN
```

list Fib   //使用list命令完成列出Fib 函数的细节时间

```go
(pprof) list Fib
     1.84s      2.75s (flat, cum) 148.65% of Total
         .          .      1:package bench
         .          .      2:
         .          .      3:import "testing"
         .          .      4:
     530ms      530ms      5:func Fib(n int) int {
     260ms      260ms      6:   if n < 2 {
     130ms      130ms      7:           return n
         .          .      8:   }
     920ms      1.83s      9:   return Fib(n-1) + Fib(n-2)
         .          .     10:}
```

web 命令生成堆栈图像信息

![image-20191031201153990](C:\Users\weiyaqi\AppData\Roaming\Typora\typora-user-images\image-20191031201153990.png)

我是没有成功过