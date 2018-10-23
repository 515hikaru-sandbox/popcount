package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/515hikaru/popcount"
)

func main() {
	item1 := "x"
	item2 := "X"
	c1 := sha256.Sum256([]byte(item1))
	c2 := sha256.Sum256([]byte(item2))
	fmt.Printf("%x\n%x\n", c1, c2)
	sum := 0
	for i := range c1 {
		sum += popcount.PopCountUint8(c1[i] ^ c2[i])
	}
	fmt.Printf("sum=%d\n", sum)
}
