package main

import (
	"fmt"
	"github.com/lifei6671/algorithms/bitmap"
)

func main() {
	fmt.Println(0x07)

	bm := bitmap.NewRepeatBitMap(100)

	for i := 1;i <= 99 ;i ++ {
		bm.Add(uint(i))
	}

	bm.Add(10)
	bm.Add(11)

	bm.PrintNoRepeat()

}
