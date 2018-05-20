package main

import "fmt"

func main() {
	v := 42 // 右側の定数の精度に基いて, int, float64, complex128のどれかになる
	fmt.Printf("v is of type %T \n", v)
}
