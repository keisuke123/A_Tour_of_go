package main

import "fmt"

func main() {
	// 空のインタフェースは任意の型の値を保持できる
	// 未知の型を扱うコードで使用される
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "Hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
