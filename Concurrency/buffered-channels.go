package main

import (
	"fmt"
)

func main() {
	// このコードだと例外を吐いて死ぬ
	// バッファがいっぱいの場合, Lockされている
	// ch := make(chan int, 1)
	// ch <- 1
	// ch <- 2
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// OK
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch) // バッファが開放される
	ch <- 2
	fmt.Println(<-ch)
}
