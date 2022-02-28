package container

import (
    "container/list"
    "fmt"
)

func listDemo() {
    l := list.New()

    l.PushBack(1)
    l.PushBack("2")
    l.PushFront(false)

    fmt.Println(l)
    fmt.Printf("%T\n", l)

    for i := l.Front(); i != nil; i = i.Next() {
        fmt.Println(i.Value)
    }
}
