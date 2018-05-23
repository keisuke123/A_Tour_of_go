package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func convertRot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	default:
		return b
	}
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)

	if err != nil {
		return 0, err
	}

	for idx, v := range b {
		b[idx] = convertRot13(v)
	}

	return n, err
}

func main() {
	fmt.Println()

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
