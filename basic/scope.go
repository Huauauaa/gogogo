package main

import (
    "fmt"
    "math"
)

// 全局变量声明必须以 var 关键字开头，如果想要在外部包中使用全局变量的首字母必须大写。
var c int

func Test8() {
    var c = 1
    fmt.Println(c)

    fmt.Println(math.MaxFloat32, math.MaxFloat64)
    // 用 Printf 函数打印浮点数时可以使用“%f”来控制保留几位小数
    fmt.Printf("%f\n", math.Pi)
    fmt.Printf("%.2f\n", math.Pi)

}
