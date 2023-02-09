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

### for
for 是 Go 唯一的循环结构。  以下是 for 循环的一些基本类型。
最基本的类型，只有一个条件。
经典的 initial/condition/after for 循环。
没有条件的 for 将重复循环，直到您跳出循环或从封闭函数返回。
您还可以继续循环的下一次迭代。
稍后当我们查看范围语句、通道和其他数据结构时，我们将看到一些其他形式。

```
package main

import (
  "fmt"
)

func main() {

  i := 1
  fmt.Println("i is ", i)

  for i <= 3 {
    fmt.Println("i is ", i)
    i = i + 1
  }

  for j := 7; j <= 9; j++ {
    fmt.Println("j is ", j)
  }

  for {
    fmt.Println("loop")
    break
  }

  for n :=0; n < 5; n++ {
    if n % 2 == 0 {
      continue
    }
    fmt.Println("n is ", n)
  }
}
```

```
$ go run for.go
i is  1
i is  1
i is  2
i is  3
j is  7
j is  8
j is  9
loop
n is  1
n is  3

```

### if else
在 Go 中用 if 和 else 分支是很直接的。
这是一个基本示例。
你可以有一个没有 else 的 if 语句。
语句可以先于条件；  此语句中声明的任何变量在当前和所有后续分支中都可用。
请注意，您不需要在 Go 中将条件括起来，但大括号是必需的。
Go 中没有三元 if，因此即使是基本条件，您也需要使用完整的 if 语句。

<a href="https://en.wikipedia.org/wiki/Ternary_conditional_operator">三元if</a>

```
package main

import (
  "fmt"
)

func main() {
  if 7 % 2 == 0 {
    fmt.Println("7 is enve")
  } else {
    fmt.Println("7 is odd")
  }

  if 8 % 2 == 0 {
    fmt.Println("8 is enve")
  } else {
    fmt.Println("8 is odd")
  }

  if 8 % 4 == 0 {
    fmt.Println("8 is divisible by 4")
  }

  if 9 % 2 == 0 {
    fmt.Println("9 is enve")
  } else {
    fmt.Println("9 is odd")
  }

  if num := 9; num < 0 {
    fmt.Println(num, "is nagative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
}

```

```
go run if_else.go
7 is odd
8 is enve
8 is divisible by 4
9 is odd
9 has 1 digit
```

###
