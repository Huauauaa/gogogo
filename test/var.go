package main

import "fmt"

var a string = "a"

func main() {
	// 标准格式
	var foo int
	// 批量格式
	var (
		bar int8
		baz int16
	)
	// 简短格式
	// 定义变量，同时显式初始化。
	// 不能提供数据类型。
	// 只能用在函数内部。
	hydrogen := 1.0
	fmt.Println(foo, bar, baz, hydrogen, a)

	x := 100
	y, z := 1, "abc"

	fmt.Println(x, y, z)
}
