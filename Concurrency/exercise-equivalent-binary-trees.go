package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)

}

func Same(t1, t2 *tree.Tree) bool {
	ch_t1 := make(chan int)
	ch_t2 := make(chan int)
	go Walk(t1, ch_t1)
	go Walk(t2, ch_t2)
	for i := 0; i < 10; i++ {
		t1_v := <-ch_t1
		t2_v := <-ch_t2
		if t1_v != t2_v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
