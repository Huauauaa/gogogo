package main

import (
    "fmt"
    "time"
)

func swap(a int, b int) {
    var temp int = a
    a = b
    b = temp
}
func swap1(pa *int, pb *int) {
	var temp int = *pa
	fmt.Println(temp)
	*pa = *pb
	fmt.Println(*pa)
	*pb = temp
	fmt.Println(*pb)
}


func main() {
    defer fmt.Println("defer 1")
    defer fmt.Println("defer 2")
    fmt.Println("Welcome to the playground!")
    fmt.Println("The time is", time.Now())

    var a int = 1
    var b int = 2

    fmt.Println(a, b)
    swap(a, b)
    fmt.Println(a, b)
	swap1(&a, &b)
    fmt.Println(a, b)
}
