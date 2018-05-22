package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

// カスタムエラーを作る
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e)) // キャストしないと無限ループ
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	prev := 0.0
	count := 0
	const EPS = 10e-9
	for math.Abs(z-prev) > EPS {
		count += 1
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
