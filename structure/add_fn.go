package structure

// Go语言方法和接收器
// 1) 面向过程实现方法
type Bag struct {
    items []int
}

func Inset(b *Bag, itemid int) {
    b.items = append(b.items, itemid)
}

// 2) Go语言的结构体方法

type Sack struct {
    items []int
}

func (s *Sack) Insert(itemId int) {
    s.items = append(s.items, itemId)
}

/*
接收器的格式如下：
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}

对各部分的说明：
接收器变量：接收器中的参数变量名在命名时，官方建议使用接收器类型名的第一个小写字母，而不是 self、this 之类的命名。例如，Socket 类型的接收器变量应该命名为 s，Connector 类型的接收器变量应该命名为 c 等。
接收器类型：接收器类型和参数类似，可以是指针类型和非指针类型。
方法名、参数列表、返回参数：格式与函数定义一致。


*/

type Point struct {
    X int
    Y int
}

// 非指针接收器的加方法

func (p Point) Add(other Point) Point {
    return Point{p.X + other.X, p.Y + other.Y}
}

func p() {}
