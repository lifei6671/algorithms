package main

import (
	"fmt"
	"github.com/lifei6671/algorithms/bitmap"
)

func main() {
	fmt.Println(0x07)

	bm := bitmap.NewBitMap(100)

	bm.Add(10)

	bm.Remove(5)

	fmt.Println(bm.Contain(10))
	bm.Remove(10)
	fmt.Println(bm.Contain(10))
}
