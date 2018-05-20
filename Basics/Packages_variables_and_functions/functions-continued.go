package main

import "fmt"

// 2つ以上の引数が同じ型の場合, 最後の型を残して省略可
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
