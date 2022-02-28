package main

import "fmt"

func Test9() {
    a := 'a'
    b := 'b'
    fmt.Println(a > b == false)
    // 注意：获取字符串中某个字节的地址属于非法行为，例如 &str[i]。
    ab := "a" + "b"
    ab1 := 'a' + 'b'
    fmt.Println(ab, ab1)

    ab += "!"
    fmt.Println(ab)

    longStr := `line1
    
    line2
    
    
    line3`
    fmt.Println(longStr)
}
