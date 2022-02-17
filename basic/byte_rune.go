package main

import "fmt"

func main() {
    var a byte = 'a'
    var a1 uint8 = 'a'
    fmt.Println(a == a1)

    var ch0 byte = 65
    var ch1 byte = 'A'
    var ch2 byte = '\x41'
    fmt.Println(ch0 == ch1, ch2 == ch1)
}
