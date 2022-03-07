package structure

import (
    "fmt"
    "testing"
)

func TestInsert(t *testing.T) {
    bag := new(Bag)
    Inset(bag, 1001)
    Inset(bag, 1002)

    fmt.Println(bag)
}

func Test(t *testing.T) {
    s := new(Sack)

    s.Insert(1)
    s.Insert(2)

    fmt.Println(s)
}

func TestP(t *testing.T) {
    p1 := Point{1, 1}
    p2 := Point{2, 2}
    result := p1.Add(p2)

    fmt.Println(p1, p2, result)
}
