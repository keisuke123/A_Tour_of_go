package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// レシーバを伴うメソッドの宣言は, レシーバ型が同じパッケージにある必要がある
// よって, 組み込みのintや他のパッケージに定義している型については定義不可
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
