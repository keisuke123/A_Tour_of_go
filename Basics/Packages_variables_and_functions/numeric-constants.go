package main

import "fmt"

// 数値の定数は高精度な値なので, すごい大きな値とかはconstで持つ
// （型を指定しない場合, 状況に応じて適切な型を取る）
const (
	// Big 1を100bit左シフト(intの範囲外だけどconstならいける)
	Big = 1 << 100
	// Small Bigを右に99bitシフト(つまり, 1<<1=2と同等)
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
