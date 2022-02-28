package main

import "fmt"

func Test1() {
    var myArray [10]int

    myArray2 := [10]int{1, 2, 3}

    fmt.Println(myArray, myArray2)
    fmt.Printf("myArray2 type is %T\n", myArray2)

    for i := 0; i < len(myArray); i++ {
        fmt.Println(myArray[i])
    }

    for index, value := range myArray2 {
        fmt.Println("index =", index, ", value =", value)
    }

    myArray3 := []int{1, 3, 5}

    for i := 0; i < len(myArray3); i++ {
        fmt.Println(myArray3[i])
    }
    fmt.Printf("myArray3 type is %T\n", myArray3)
    for index, value := range myArray3 {
        fmt.Println(index, value)
    }
}
