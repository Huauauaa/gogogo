# gogogo

## reference

- [项目目录结构](https://makeoptim.com/golang/standards/project-layout)
- [Go 语言入门教程](http://c.biancheng.net/golang/)
- [Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html)

## run

1. `go build test/hello.go` `./hello.exe`/`./hello`
2. `go run test/hello.go`

## 数组

### 定长数组

```go
var a = [3]int
```

### 动态数组

```go
a := []int{1,2,3}
```

## 基本语法

### 变量的声明

#### 基本类型

- bool
- string
- 有符号整数类型: int/int8/int16/int32/int64
- 无符号整数类型: uint/unit8/unit16/unit32/unit64/uintptr
- type(uint8 的别名)
- rube(int32 的别名,代表一个 Unicode)
- 小数类型: float32/float64
- 复数: complex64/complex128

### 初始化

- `var [name] [type] = [expression]`
- `[name] := [value]`

### 生命周期

变量的生命周期与变量的作用域有着不可分割的联系：

- 全局变量：它的生命周期和整个程序的运行周期是一致的；
- 局部变量：它的生命周期则是动态的，从创建这个变量的声明语句开始，到这个变量不再被引用为止；
- 形式参数和函数返回值：它们都属于局部变量，在函数被调用的时候创建，函数调用结束后被销毁。

#### heap vs stack

区别在于：

- 堆（heap）：堆是用于存放进程执行中被动态分配的内存段。它的大小并不固定，可动态扩张或缩减。当进程调用 malloc 等函数分配内存时，新分配的内存就被动态加入到堆上（堆被扩张）。当利用 free 等函数释放内存时，被释放的内存从堆中被剔除（堆被缩减）；
- 栈(stack)：栈又称堆栈， 用来存放程序暂时创建的局部变量，也就是我们函数的大括号{ }中定义的局部变量。

## 函数

### 处理运行时错误

> Go 语言的错误处理思想及设计包含以下特征：
>
> - 一个可能造成错误的函数，需要返回值中返回一个错误接口（error），如果调用是成功的，错误接口将返回 nil，否则返回错误。
> - 在函数调用后需要检查错误，如果发生错误，则进行必要的错误处理。

### 测试规则

编写测试用例有以下几点需要注意：

- 测试用例文件不会参与正常源码的编译，不会被包含到可执行文件中；
- 测试用例的文件名必须以`_test.go` 结尾；
- 需要使用 import 导入 testing 包；
- 测试函数的名称要以 Test 或 Benchmark 开头，后面可以跟任意字母组成的字符串，但第一个字母必须大写，例如 TestAbc()，一个测试用例文件中可以包含多个测试函数；
- 单元测试则以`(t *testing.T)`作为参数，性能测试以`(t *testing.B)`做为参数；
- 测试用例文件使用 go test 命令来执行，源码中不需要 main() 函数作为入口，所有以 `_test.go` 结尾的源码文件内以 Test 开头的函数都会自动执行。

- 覆盖率测试 `go test -cover`
