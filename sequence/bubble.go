package sequence

func Bubble(list []int) ([]int) {

	v := 0
	for i := 0; i < len(list) -1 ; i ++ {
		for j := len(list) - 1; j >= i; j-- {
			v = list[i]
			if list[j] > v {
				list[i] = list[j]
				list[j] = v
			}
		}
	}
	return list
}


func BubbleOptimal(list []int) ([]int) {

	v := 0
	bChange := false
	for i := 0; i < len(list) -1 ; i ++ {
		for j := len(list) - 1; j >= i; j-- {
			v = list[i]
			if list[j] > v {
				list[i] = list[j]
				list[j] = v
				bChange = true
			}
		}
		if bChange == false {
			break
		}
	}
	return list
}