package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)

	var hoge []uint8
	for i := 0; i < dy; i++ {
		ret[i] = make([]uint8, dx)
	}

	for j := 0; j < dy; j++ {
		for i := 0; i < dx; i++ {
			ret[i][j] = uint8((i + j) / 2)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
