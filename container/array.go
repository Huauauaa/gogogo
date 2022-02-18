package main

import "fmt"

func main() {
    var a [3]int

    fmt.Println(a, a[0], a[len(a)-1])

    for i, v := range a {
        fmt.Println(i, v)
    }

    for _, v := range a {
        fmt.Printf("%d\n", v)
    }

    q := [3]int{1, 2, 3}
    r := [3]int{1, 2}
    // index 2 is out of bounds (>= 2)
    // r1 := [2]int{1, 2, 3}
    // 在数组的定义中，如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算
    q1 := [...]int{2, 3}

    fmt.Println(q, r, q1, len(q1))
}
