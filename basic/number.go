package main

import "fmt"

func Test6() {
    var x complex128 = complex(1, 2)
    fmt.Println(x)
    var y complex128 = complex(3, 4)
    fmt.Println(x * y)
    fmt.Println(real(x * y))
    fmt.Println(imag(x * y))
}
