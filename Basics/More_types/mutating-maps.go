package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("THe value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// キーに対する値があるかどうかは2つ目の値で確認
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
