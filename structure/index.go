package structure

import "fmt"

type People struct {
    name  string
    child *People
}

func relation() {
    r := &People{
        name: "grandpa",
        child: &People{
            name: "dad",
            child: &People{
                name: "me",
            },
        },
    }

    fmt.Print(r, r.child, r.child.child)
}

type Address struct {
    Province    string
    City        string
    ZipCode     int
    PhoneNumber string
}

/*
ins := 结构体类型名{
    字段1的值,
    字段2的值,
    …
}
- 必须初始化结构体的所有字段。
- 每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。
- 键值对与值列表的初始化形式不能混用。
*/
func addr() {
    a := Address{
        "Shaanxi",
        "Xi'an",
        710000,
        "15300000000",
    }

    fmt.Println(a)
}

/*
ins := struct {
    // 匿名结构体字段定义
    字段1 字段类型1
    字段2 字段类型2
    …
}{
    // 字段值初始化
    初始化字段1: 字段1的值,
    初始化字段2: 字段2的值,
    …
}
*/
func printMsgType(msg *struct {
    id   int
    data string
}) {
    // 使用动词%T打印msg的类型
    fmt.Printf("%T\n", msg)
}

type Cat struct {
    Color string
    Name  string
}

func NewCatByName(name string) *Cat {
    return &Cat{
        Name: name,
    }
}

func NewCatByColor(color string) *Cat {
    return &Cat{
        Color: color,
    }
}

func index() {}
