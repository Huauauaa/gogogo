package main

import "fmt"

type Weekday int

func main() {
    const (
        Sunday Weekday = iota
        Monday
        Tuesday
        wednesday
        Thursday
        Friday
        Saturday
    )

    fmt.Println(Friday)

}
