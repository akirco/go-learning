- [golang learning](#golang-learning)
  - [intro](#intro)
  - [install](#install)
  - [hello world](#hello-world)
  - [基础语法](#基础语法)
    - [Go 标记](#go-标记)
    - [注释](#注释)
    - [标识符](#标识符)
    - [字符串连接](#字符串连接)
    - [关键字](#关键字)
    - [标识符](#标识符-1)
    - [格式化字符串](#格式化字符串)
    - [数据类型](#数据类型)
    - [常量](#常量)
      - [`iota`的使用及用法](#iota的使用及用法)
    - [运算符](#运算符)


# golang learning


## intro

Go 语言最主要的特性：

- 自动垃圾回收
- 更丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性

## install

```shell
scoop install go
paru -S go
```

## hello world

```go
package main

import "fmt"

func main() {
	fmt.Println("hello")
}
```

- Go 语言的基础组成有以下几个部分：

  - 包声明
  - 引入包
  - 函数
  - 变量
  - 语句 & 表达式
  - 注释

- 包说明

  - 文件名与包名没有直接关系，不一定要将文件名与包名定成同一个。
  - 文件夹名与包名没有直接关系，并非需要一致。
  - 同一个文件夹下的文件只能有一个包名，否则编译报错。

- run

```shell
go run .
#hello
```

- build

```shell
go build .
```

## 基础语法

### Go 标记

Go 程序可以由多个标记组成，可以是`关键字`，`标识符`，`常量`，`字符串`，`符号`。如以下 GO 语句由 6 个标记组成：

```go
fmt.Println("Hello, World!")
```

### 注释

```go
// 单行注释
/*
 我是多行注释
 */
```

### 标识符

标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z 和 a~z)数字(0~9)、下划线\_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

### 字符串连接

```go
package main
import "fmt"
func main() {
    fmt.Println("hello" + "world")
}
```

### 关键字

25 个关键字及简单使用示例：

1. package：用于定义代码所属的包。

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

2. import：用于导入其他包中的代码。

```go
import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

3. func：用于定义函数。

```go
func add(a, b int) int {
    return a + b
}
```

4. var：用于声明变量。

```go
var x int = 5
```

5. const：用于声明常量。

```go
const pi = 3.14
```

6. type：用于定义自定义类型。

```go
type Point struct {
    x, y int
}
```

7. struct：用于定义结构体。

```go
type Point struct {
    x, y int
}
```

8. interface：用于定义接口。

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

9. map：用于定义映射（键值对）。

```go
m := make(map[string]int)
m["apple"] = 1
```

10. slice：用于定义切片（动态数组）。

```go
s := []int{1, 2, 3, 4, 5}
```

11. range：用于迭代数组、切片、映射等。

```go
s := []int{1, 2, 3, 4, 5}
for i, v := range s {
    fmt.Println(i, v)
}
```

12. for：用于循环执行代码块。

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

13. if：用于条件判断。

```go
x := 5
if x > 0 {
    fmt.Println("Positive")
} else if x < 0 {
    fmt.Println("Negative")
} else {
    fmt.Println("Zero")
}
```

14. else：用于 if 语句的可选分支。

```go
x := 5
if x > 0 {
    fmt.Println("Positive")
} else {
    fmt.Println("Non-positive")
}
```

15. switch：用于多条件判断。

```go
x := 5
switch x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

16. select：用于选择可用的通道操作。

```go
ch1 := make(chan int)
ch2 := make(chan int)

select {
case <-ch1:
    fmt.Println("Received from ch1")
case <-ch2:
    fmt.Println("Received from ch2")
}
```

17. go：用于并发执行函数。

```go
go func() {
    fmt.Println("Hello, World!")
}()
```

18. defer：用于延迟执行函数。

```go
defer fmt.Println("Goodbye, World!")

fmt.Println("Hello, World!")
```

19. fallthrough：用于在 switch 语句中继续执行下一个分支。

```go
x := 2
switch x {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
}
```

20. break：用于跳出循环或 switch 语句。

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        break
    }
    fmt.Println(i)
}
```

21. continue：用于跳过当前循环的剩余代码，继续下一次循环。

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        continue
    }
    fmt.Println(i)
}
```

22. return：用于从函数中返回结果。

```go
func add(a, b int) int {
    return a + b
}
```

23. chan: 声明一个通道（channel）。通道是用来在不同的 Go 协程之间进行通信的一种机制

```go
package main

import "fmt"

func main() {
    // 创建一个非缓冲的通道
    ch := make(chan int)

    // 启动一个协程，向通道发送数据
    go func() {
        ch <- 10
    }()

    // 从通道接收数据
    x := <-ch
    fmt.Println(x)
}
```

24. default: 默认选项，用在 switch 和 select 语句种。

```go
package main

import "fmt"

func main() {
    num := 3

    switch num {
    case 1:
        fmt.Println("数字是1")
    case 2:
        fmt.Println("数字是2")
    default:
        fmt.Println("数字不是1或2")
    }
}
```

25. go：用于启动并发执行的 goroutine。

```go
go func() {
    fmt.Println("Hello, World!")
}()
```

### 标识符

1. `append`：用于向切片追加元素。
2. `bool`：布尔类型。
3. `byte`：uint8 的别名。
4. `cap`：返回切片或映射的容量。
5. `close`：关闭一个通道。
6. `complex`：复数类型。
7. `complex64`：复数类型，由 float32 表示实部和虚部。
8. `complex128`：复数类型，由 float64 表示实部和虚部。
9. `copy`：将源切片的元素复制到目标切片。
10. `delete`：从映射中删除指定的键值对。
11. `error`：错误类型。
12. `false`：布尔类型的假值。
13. `float32`：32 位浮点数类型。
14. `float64`：64 位浮点数类型。
15. `imag`：返回复数的虚部。
16. `int`：整数类型。
17. `int8`：8 位整数类型。
18. `int16`：16 位整数类型。
19. `int32`：32 位整数类型。
20. `int64`：64 位整数类型。
21. `iota`：常量生成器。
22. `len`：返回切片、数组、字符串、映射或通道的长度。
23. `make`：创建切片、映射或通道。
24. `new`：分配内存。
25. `nil`：空值。
26. `panic`：引发运行时错误。
27. `print`：打印到标准输出。
28. `println`：打印到标准输出并换行。
29. `real`：返回复数的实部。
30. `recover`：从 panic 中恢复。
31. `rune`：int32 的别名，表示 Unicode 码点。
32. `string`：字符串类型。
33. `true`：布尔类型的真值。
34. `uint`：无符号整数类型。
35. `uint8`：8 位无符号整数类型。
36. `uint16`：16 位无符号整数类型。
37. `uint32`：32 位无符号整数类型。
38. `uint64`：64 位无符号整数类型。
39. `uintptr`：无符号整数类型，用于存储指针。

### 格式化字符串

1. `fmt.Sprintf`

`Sprintf` 根据格式化参数生成格式化的字符串并返回该字符串。

```go
package main

import (
    "fmt"
)

func main() {
   // %d 表示整型数字，%s 表示字符串
    var stockcode=123
    var enddate="2020-12-31"
    var url="Code=%d&endDate=%s"
    var target_url=fmt.Sprintf(url,stockcode,enddate)
    fmt.Println(target_url)
}
```

2. `fmt.Printf`

`Printf` 根据格式化参数生成格式化的字符串并写入标准输出。

```go
package main

import (
    "fmt"
)

func main() {
   // %d 表示整型数字，%s 表示字符串
    var stockcode=123
    var enddate="2020-12-31"
    var url="Code=%d&endDate=%s"
    fmt.Printf(url,stockcode,enddate)
}
```

### 数据类型

- 布尔类型

  布尔型的值只可以是常量 true 或者 false。

- 数字类型

  整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。

  ```go
  /**
  uint8无符号 8 位整型 (0 到 255)

  uint16
  无符号 16 位整型 (0 到 65535)

  uint32
  无符号 32 位整型 (0 到 4294967295)

  uint64
  无符号 64 位整型 (0 到 18446744073709551615)

  int8
  有符号 8 位整型 (-128 到 127)

  int16
  有符号 16 位整型 (-32768 到 32767)

  int32
  有符号 32 位整型 (-2147483648 到 2147483647)

  int64
  有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

  float32
  IEEE-754 32位浮点型数

  float64
  IEEE-754 64位浮点型数

  complex64
  32 位实数和虚数

  complex128
  64 位实数和虚数

  byte
  类似 uint8

  rune
  类似 int32

  uint
  32 或 64 位

  int
  与 uint 一样大小

  uintptr
  无符号整型，用于存放一个指针
  */
  ```

- 字符串类型

  字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。

- 派生类型

  (a) 指针类型（Pointer）

  (b) 数组类型

  (c) 结构化类型(struct)

  (d) Channel 类型

  (e) 函数类型

  (f) 切片类型

  (g) 接口类型（interface）

  (h) Map 类型

- 变量

```go
// var varName varType;

// 例如，声明一个整数变量x：

var x int

// 也可以同时声明多个变量：

var x, y int

// 如果要给变量赋初值，可以使用等号进行赋值操作：

var x int = 10

// 也可以使用类型推断，省略变量类型(Go 1.10)：

var x = 10

// 还可以使用:=简化变量声明和赋值的过程：

x := 10

// :=符号是用于声明和初始化变量的简写形式。它可以同时完成变量的声明和赋值操作。

g,h := 123,"hello"
```

- `:=`符号的限制

```go
1. `:=`只能用于函数内部，不能用于全局变量的声明和初始化。
2. 变量必须是新声明的，不能是已经存在的变量。
3. 只能在同一个作用域内使用，不能在外层作用域重新声明已经存在的变量。
4. 只能用于初始化变量，不能用于赋值操作。
```

### 常量

```go
// const varName varType

const identifier [type] = value
```

其中，`identifier`是常量的名称，`type`是常量的类型（可选），`value`是常量的值。

以下是一些常量声明的示例：

```go
const pi = 3.14159
const maxAge int = 100
const message string = "Hello, World!"
```

在上面的示例中，`pi`是一个未指定类型的常量，`maxAge`是一个整数类型的常量，`message`是一个字符串类型的常量。

常量还可以使用`iota`关键字进行枚举，例如：

```go
const (
    Monday = iota
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)
```

在上面的示例中，`Monday`的值为 0，`Tuesday`的值为 1，以此类推，`Sunday`的值为 6。

需要注意的是，常量的值必须是编译时可确定的表达式，例如字面值、常量表达式、函数调用等。不能使用运行时才能确定的表达式作为常量的值。

#### `iota`的使用及用法

在 Go 语言中，`iota`是一个预定义的标识符，用于枚举常量的自增值。

`iota`在常量声明中使用，它的初始值为 0，每次在常量声明中使用时，它的值都会自动递增。当遇到新的常量声明时，`iota`会被重置为 0。

以下是`iota`的一些常见用法：

1. 枚举常量：可以使用`iota`来定义一组自增的常量。例如：

```go
const (
    Monday = iota // 0
    Tuesday      // 1
    Wednesday    // 2
    Thursday     // 3
    Friday       // 4
    Saturday     // 5
    Sunday       // 6
)
```

2. 位运算：可以使用`iota`和位运算符来定义一组按位掩码常量。例如：

```go
const (
    ReadPermission = 1 << iota // 1
    WritePermission            // 2
    ExecutePermission          // 4
)
```

3. 跳过值：可以使用`_`来跳过`iota`的某些值。例如：

```go
const (
    _  = iota // 跳过0
    KB = 1 << (10 * iota) // 1 << (10 * 1)
    MB = 1 << (10 * iota) // 1 << (10 * 2)
    GB = 1 << (10 * iota) // 1 << (10 * 3)
    TB = 1 << (10 * iota) // 1 << (10 * 4)
)
```

在上述例子中，`iota`的值在每个常量声明中都会自动递增，从而实现了按照指定规则生成常量的功能。

需要注意的是，`iota`只在常量声明中有效，不能在其他地方使用。

### 运算符

- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符
- 其他运算符
