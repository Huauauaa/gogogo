package main

import (
    "flag"
    "fmt"
)

// 通过 flag.String，定义一个 mode 变量，这个变量的类型是 *string。后面 3 个参数分别如下：
//      参数名称：在命令行输入参数时，使用这个名称。
//      参数值的默认值：与 flag 所使用的函数创建变量类型对应，String 对应字符串、Int 对应整型、Bool 对应布尔型等。
//      参数说明：使用 -help 时，会出现在说明中。
var mode = flag.String("mode", "default value", "process mode")

func main() {
    // 指针（pointer）在Go语言中可以被拆分为两个核心概念：
    // 类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
    // 切片，由指向起始元素的原始指针、元素数量和容量组成。

    cat := 1
    str := "banana"

    fmt.Println(&cat, &str)

    x, y := 1, 2
    fmt.Println(x, y)
    swap(&x, &y)
    fmt.Println(x, y)

    flag.Parse()

    fmt.Println(*mode)

    // 创建方式
    // 1. [ptr] := &[v]
    // 2. new(type)

    str1 := new(string)

    fmt.Println(str1, "---", *str1)
    *str1 = "hello"
    fmt.Println(str1, "---", *str1)
}

func swap(a, b *int) {
    t := *a
    *a = *b
    *b = t
}
