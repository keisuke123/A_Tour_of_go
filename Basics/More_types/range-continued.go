package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// indexのみ取得
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	// valueのみ取得(_で捨てる)
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
