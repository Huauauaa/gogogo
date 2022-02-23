package main

import "fmt"

func sum(args ...int) int {
    result := 0
    for _, v := range args {
        result += v
    }
    return result
}

func fibonacci(n int) (res int) {
    if n <= 2 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
    // fmt.Println(sum(4, 5))
    // fmt.Println(sum(7, 8, 9))
    fmt.Println(fmt.Println("a"))

    for i := 1; i < 10; i++ {
        fmt.Print(fibonacci(i), " ")
    }
}
