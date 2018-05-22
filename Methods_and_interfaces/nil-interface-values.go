package main

import "fmt"

// 具体的な メソッドを示す型がインターフェースのタプル内に存在しないためエラー
type I interface {
	M()
}	

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i
	describe(i)
	i.M()
}
