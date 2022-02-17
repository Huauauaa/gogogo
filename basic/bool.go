package main

import "fmt"

func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}

func main() {
    var aVar = 10
    fmt.Println(aVar == 5, aVar == 10)
    fmt.Println(btoi(true), btoi(false))
}
