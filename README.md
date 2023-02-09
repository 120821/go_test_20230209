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

### variables
在 Go 中，编译器显式声明和使用变量，例如检查函数调用的类型正确性。
var 声明 1 个或多个变量。
您可以一次声明多个变量。
Go 会推断初始化变量的类型。
没有相应初始化而声明的变量为零值。  例如，int 的零值为 0。
:= 语法是声明和初始化变量的简写，例如在本例中为 var f string = "apple"。  此语法仅在函数内部可用。

```
package main

import "fmt"

func main() {
  var a = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e)

  f := "apple"
  fmt.Println(f)

}

```

```
$ go run variables.go
initial
1 2
true
0
apple

```

### constants
Go 支持字符、字符串、布尔值和数值常量。
const 声明一个常量值。
const 语句可以出现在 var 语句可以出现的任何地方。
常量表达式以任意精度执行算术运算。
数字常量在被赋予类型之前是没有类型的，例如通过显式转换。
可以通过在需要类型的上下文中使用数字来为数字赋予类型，例如变量赋值或函数调用。  例如，此处 math.Sin 需要 float64。

```
package main

import (
  "fmt"
  "math"
)

const s string = "constant"

func main() {

  fmt.Println(s)

  const n = 500000000

  const d = 3e20 / n
  fmt.Println(d)
  fmt.Println(int64(d))
  fmt.Println(math.Sin(n))
}
```

```
$ go run constant.go
constant
6e+11
600000000000
-0.28470407323754404
```
