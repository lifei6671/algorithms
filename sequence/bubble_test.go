package sequence

import "testing"

func Test_Bubble(t *testing.T)  {
	array := []int{10,50,60,1,29,95,02,6,025,4521,020,4515,2,5,15,24,65,6,1,051,10,24,45,1,4,51,42}

	list := Bubble(array)

	t.Logf("%+v",list)
}

func Test_BubbleOptimal(t *testing.T) {
	array := []int{10,50,60,1,29,95,02,6,025,4521,020,4515,2,5,15,24,65,6,1,051,10,24,45,1,4,51,42}

	list := BubbleOptimal(array)

	t.Logf("%+v",list)
}