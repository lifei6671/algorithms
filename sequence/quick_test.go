package sequence

import "testing"

func Test_QuickSort(t *testing.T) {
	list := []int{10,50,60,1,29,95,02,6,025,4521,20,4515,2,5,15,24,65,6,1,051,10,24,45,1,4,51,42}

	QuickSort(list,0,len(list) - 1)
}
