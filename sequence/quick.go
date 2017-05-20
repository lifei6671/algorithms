package sequence

import "fmt"

func division(list []int, left, right int) int {
	//定义一个基准变量
	base := list[left]

	//当 left 小于 right 时循环遍历
	for left < right {
		//从序列右端开始，向左遍历，直到找到小于base的数
		for left < right && list[right] > base {
			right --
		}
		// 找到了比base小的元素，将这个元素放到最左边的位置
		list[left] = list[right]
		for left < right && list[left] <= base {
			left ++
		}
		list[right] = list[left]
	}

	// 最后将base放到left位置。此时，left位置的左侧数值应该都比left小；
	// 而left位置的右侧数值应该都比left大。
	list[left] = base;
	return left;
}

func QuickSort(list []int,left, right int) {
	// 左下标一定小于右下标，否则就越界了
	if left < right {
		base := division(list,left,right)

		fmt.Println("base =",list[base])
		fmt.Println("left ",left,"right",right)
		QuickSort(list,left,base - 1)
		QuickSort(list,base + 1,right)
	}
	defer func(list []int) {
		fmt.Printf("list => %+v",list)
		fmt.Println()
	}(list)

}