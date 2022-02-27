package fn

import (
    "fmt"
    "testing"
)

func TestDuration(t *testing.T) {
    duration()
}

func TestSum(t *testing.T) {
    total := sum(4, 5)
    if total != 9 {
        t.Error("Test fail")
    }
}

func TestFibonacci(t *testing.T) {
    for i := 1; i < 10; i++ {
        fmt.Print(fibonacci(i), " ")
    }
}
