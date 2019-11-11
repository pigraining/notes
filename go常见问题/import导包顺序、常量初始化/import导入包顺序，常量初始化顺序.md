## import导入包顺序，常量初始化顺序

### 搜索路径

import用于导入包：

```go
import (
    "fmt"
    "net/http"
    "mypkg"
)
```

编译器会根据上面指定的相对路径去搜索包然后导入，这个相对路径是从GOROOT或GOPATH(workspace)下的src下开始搜索的。

假如go的安装目录为`/usr/local/go`，也就是说`GOROOT=/usr/local/go`，而GOPATH环境变量`GOPATH=~/mycode:~/mylib`，那么要搜索`net/http`包的时候，将按照**如下顺序**进行搜索：

```
/usr/local/go/srcnet/http
~/mycode/src/net/http
~/mylib/src/net/http
```

以下是`go install`搜索不到mypkg包时的一个报错信息：

```
can't load package: package mypkg: cannot find package "mypkg" in any of:
        /usr/lib/go-1.6/src/mypkg (from $GOROOT)
        /golang/src/mypkg (from $GOPATH)
```

也就是说，go总是先从`GOROOT`出先搜索，再从`GOPATH`列出的路径顺序中搜索，只要一搜索到合适的包就理解停止。当搜索完了仍搜索不到包时，将报错。

包导入后，就可以使用这个包中的属性。使用`包名.属性`的方式即可。例如，调用fmt包中的Println函数`fmt.Println`。

### 包导入的过程

![img](https://img2018.cnblogs.com/blog/733013/201810/733013-20181023224911978-1960747966.png)

首先从main包开始，如果main包中有import语句，则会导入这些包，如果要导入的这些包又有要导入的包，则继续先导入所依赖的包。重复的包只会导入一次，就像很多包都要导入fmt包一样，但它只会导入一次。

每个被导入的包在导入之后，都会先将包的可导出函数(大写字母开头)、包变量、包常量等声明并初始化完成，然后如果这个包中定义了init()函数，则自动调用init()函数。init()函数调用完成后，才回到导入者所在的包。同理，这个导入者所在包也一样的处理逻辑，声明并初始化包变量、包常量等，再调用init()函数(如果有的话)，依次类推，直到回到main包，main包也将初始化包常量、包变量、函数，然后调用init()函数，调用完init()后，调用main函数，于是开始进入主程序的执行逻辑。

## 常量和变量的初始化

Go中的常量在编译期间就会创建好，即使是那些定义为函数的本地常量也如此。常量只允许是数值、字符(runes)、字符串或布尔值。

由于编译期间的限制，定义它们的表达式必须是编译器可评估的常量表达式(constant expression)。例如，`1<<3`是一个常量表达式，而`math.Sin(math.Pi/4)`则不是常量表达式，因为涉及了函数math.Sin()的调用过程，而函数调用是在运行期间进行的。

变量的初始化和常量的初始化差不多，但初始化的变量允许是"需要在执行期间计算的一般表达式"。例如：

```
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```

## init()函数

Go中除了保留了main()函数，还保留了一个init()函数，这两个函数都不能有任何参数和返回值。它们都是在特定的时候自动调用的，无需我们手动去执行。

还是这张图：

![img](https://img2018.cnblogs.com/blog/733013/201810/733013-20181023224911978-1960747966.png)

每个包中都可以定义init函数，甚至可以定义多个，但建议每个包只定义一个。每次导入包的时候，在导入完成后，且变量、常量等声明并初始化完成后，将会调用这个包中的init()函数。

对于main包，如果main包也定义了init()，那么它会在main()函数之前执行。当main包中的init()执行完之后，就会立即执行main()函数，然后进入主程序。

所以，init()经常用来初始化环境、安装包或其他需要在程序启动之前先执行的操作。如果import导入包的时候，发现前面命名为下划线`_`了，一般就说明所导入的这个包有init()函数，且导入的这个包除了init()函数外，没有其它作用。