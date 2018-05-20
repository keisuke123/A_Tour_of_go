package main

import (
	"fmt"
	"math"
)

func main() {
	// アクセス不可能(1文字目が大文字の場合エクスポートされている)
	// fmt.Println(math.pi)

	// アクセス可能
	fmt.Println(math.Pi)
}
