package main

import "fmt"

func getData() (int, int) {
    return 100, 200
}

func Test() {
    // 匿名变量的特点是一个下画线“_”，“_”本身就是一个特殊的标识符，被称为空白标识符
    // 匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
    a, _ := getData()
    _, b := getData()

    fmt.Println(a, b)
}
