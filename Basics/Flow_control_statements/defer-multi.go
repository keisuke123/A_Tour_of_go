package main

import "fmt"

func main() {
	fmt.Println("counting")

	// deferはスタックされて, 関数が終わると後ろから実行(LIFO)
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done!")
}
