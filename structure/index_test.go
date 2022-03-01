package structure

import (
    "fmt"
    "testing"
)

func TestRelation(t *testing.T) {
    relation()
}

func TestAddr(t *testing.T) {
    addr()
}

func TestPintMsgType(t *testing.T) {
    msg := &struct { // 定义部分
        id   int
        data string
    }{ // 值初始化部分
        1024,
        "hello",
    }
    printMsgType(msg)
}

func TestIndex(t *testing.T) {
    tomCat := &Cat{"tom", "Blue"}
    fmt.Println(tomCat)
    // tomCat = NewCatByColor("Blue")
    // fmt.Println(tomCat)
}
