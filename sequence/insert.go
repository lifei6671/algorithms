package sequence

func InsertSort(list []int) []int {
	// 因为前i-1个数都是从小到大的有序序列，所以只要当前比较的数(list[j])比temp大，就把这个数后移一位
	for i := 1; i < len(list) - 1; i ++ {
		j := 0
		// 取出第i个数，和前i-1个数比较后，插入合适位置
		temp := list[i]
		// 因为前i-1个数都是从小到大的有序序列，所以只要当前比较的数(list[j])比temp大，就把这个数后移一位
		for j = i -1; j > 0 && temp  < list[j]; j -- {
			list[j + 1] = list[j]
		}
		list[j + 1] = temp
	}
	return list
}
