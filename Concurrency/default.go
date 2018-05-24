package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)  // 100msごとに通知
	boom := time.After(500 * time.Millisecond) // 500ms後に通知
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
