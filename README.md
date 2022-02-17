# gogogo

## reference

- [项目目录结构](https://makeoptim.com/golang/standards/project-layout)
- [Go 语言入门教程](http://c.biancheng.net/golang/)

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
- int/int8/int16/int64
- uint/unit8/unit16/unit32/unit64/uintptr
- type(uint8 的别名)
- rube(int32 的别名,代表一个 Unicode)
- float32/float64
- complex64/complex128

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
