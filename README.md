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
