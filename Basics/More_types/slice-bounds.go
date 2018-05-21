package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	// s[0:2]と等価
	s = s[:2]
	fmt.Println(s)

	// s[1:sの要素数]と等価
	s = s[1:]
	fmt.Println(s)
}
