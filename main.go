package main

import (
	"algorithms/string"
	"fmt"
)

func main() {

	str := algorithms.LeftMoveRight("abcdefg",3);

	fmt.Println("字符串转移",str);

	has := algorithms.StringContain("ABCDEFG","EDC");

	fmt.Println("是否包含",has);

	brother := algorithms.FindBrothers([]string{"ABD","DEFG","KIJ","JIK","IJK","ULFA","KIJJ",},"KIJ");

	fmt.Println("查找兄弟字符串:",brother);
}
