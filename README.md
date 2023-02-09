### hello world
我们的第一个程序将打印经典的“hello world”消息。  这是完整的源代码。
要运行该程序，请将代码放入 hello-world.go 并使用 go run。
有时我们想将我们的程序构建成二进制文件。  我们可以使用 go build 来做到这一点。
然后我们可以直接执行构建的二进制文件。
现在我们可以运行和构建基本的 Go 程序，让我们进一步了解这门语言。

```
package main

import "fmt"

func main() {
  fmt.Println("hello world")
}

```

```
$ go run hello-world.go
hello world

$ go build hello-world.go
$ ls
hello-world    hello-world.go

$ ./hello-world
hello world

```

### Values
Go 有多种值类型，包括字符串、整数、浮点数、布尔值等。下面是一些基本示例。
字符串，可以与 +.
整数和浮点数。
布尔值，如您所料，具有布尔运算符。

```
package main

import "fmt"

func main() {
  fmt.Println("go" + "long")
  fmt.Println("1 + 1 =", 1 + 1)
  fmt.Println("70 / 10 =" + 70 / 10)
  fmt.Println(true && false)
  fmt.Println(true || false)
  fmt.Println(!true)
}
```

```
$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```
