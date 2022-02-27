package fn

import (
    "fmt"
    "time"
)

func duration() {
    start := time.Now()
    sum := 0
    for i := 0; i < 100000000; i++ {
        sum++
    }
    elapsed := time.Now().Sub(start)
    fmt.Println("duration is", elapsed)
}

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
