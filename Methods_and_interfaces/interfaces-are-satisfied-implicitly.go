package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// これでインタフェースIを実装したことになるので
// var i Iに代入可能
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{
		"Hello",
	}
	i.M()
}
