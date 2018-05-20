package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!" // 型推論
	fmt.Println(i, j, c, python, java)       // 1 2 true false no!
}
