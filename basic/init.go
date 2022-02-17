package main

import "fmt"

func main() {
	var hp int

	// hp := 10
	// error: no new variables on left side of :=
	// 由于使用了:=，而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，将会发生编译错误。

	fmt.Println(hp)
}
