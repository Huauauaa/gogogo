package main

import "fmt"

func main() {
    a := [3]int{1, 2, 3}
    fmt.Println(a, a[1:2], a[2:], a[:2])

    var highRiseBuilding [30]int

    for i := 0; i < 30; i++ {
        highRiseBuilding[i] = i + 1
    }
    fmt.Println(highRiseBuilding)
    // 直接声明新的切片
    var foo = [2]int{}
    var bar = []int{}
    fmt.Print("", foo, bar)

    var a1 []int
    a1 = append(a1, 1)

    fmt.Println(a1)
    a1 = append(a1, 1, 2, 3)
    fmt.Println(a1)
    a1 = append(a1, []int{10, 20}...)
    fmt.Println(a1)
}
