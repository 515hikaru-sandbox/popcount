package main

import (
	"fmt"
	"github.com/515hikaru/popcount"
)

func main() {
	var x uint64
	x = 32
	fmt.Printf("%d\n", popcount.PopCount(x))
}
