package main

import (
	"fmt"
	"reflect"
)

type IPAddr [4]byte

// IPAddrにString() stringメソッドを実装せよ
func (host IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", host[0], host[1], host[2], host[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, addr := range hosts {
		fmt.Printf("%v: %v\n", name, addr)
	}
}
