package main

import "fmt"

// 型は後ろに書く
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
