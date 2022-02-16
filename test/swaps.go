package main

import "fmt"

func swap1(a int, b int) (int, int) {
	var t int
	t = a
	a = b
	b = t
	return a, b
}

func swap2(a int, b int) (int, int) {
	a = a ^ b
	b = b ^ a
	a = a ^ b
	return a, b
}

func swap3(a int, b int) (int, int) {
	// b, a = a, b
	// return a, b
	return b, a
}
func main() {
	a := 100
	b := 200
	fmt.Println(a, b)
	fmt.Println(swap1(a, b))
	fmt.Println(swap2(a, b))
	fmt.Println(swap3(a, b))
}
