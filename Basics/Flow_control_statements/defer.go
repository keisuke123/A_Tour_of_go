package main

import "fmt"

func main() {
	// 呼び出し元の関数がreturnするまで実行が遅延する
	// 引数自体は即時評価される
	defer fmt.Println("world")

	fmt.Println("hello")
}
