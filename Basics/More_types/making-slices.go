package main

import "fmt"

func printSlice(str string, s []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", str, len(s), cap(s), s)
}

func main() {
	a := make([]int, 5)
	// a len=5 cap=5 [0, 0, 0, 0, 0]
	printSlice("a", a)

	b := make([]int, 0, 5)
	// b len=0 cap=5 []
	printSlice("b", b)

	c := b[:2]
	// c len=2 cap=5 [0 0]
	printSlice("c", c)

	d := c[2:5]
	// d len=3 cap=3 [0, 0, 0]
	printSlice("d", d)
}
