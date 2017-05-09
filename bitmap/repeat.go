package bitmap

import "fmt"

//在2.5亿个整数中找出不重复的整数，注，内存不足以容纳这2.5亿个整数？.
type RepeatBitMap struct {
	bits []byte
	capacity uint
}

func NewRepeatBitMap(capacity uint) *RepeatBitMap {
	return &RepeatBitMap{
		bits : make([]byte,(capacity / 4)+1),
		capacity: capacity,
	}
}

//添加元素到集合中，添加时先判断是否存在，如果已存在，则修改第二位为1，否则修改第一位为1
func (b *RepeatBitMap) Add(num uint) {
	index := num / 4
	pos := (num % 4) * 2

	if b.Contain(num) {
		b.bits[index] = b.bits[index] | (1 << (pos + 1))
	}else{
		b.bits[index] = b.bits[index] | (1 << pos)
	}
}

func (b *RepeatBitMap) Contain(num uint) bool {
	index := num / 4
	pos := (num % 4) * 2

	 return (b.bits[index] & (1 << pos)) != 0
}

func (b *RepeatBitMap) IsRepeat(num uint) bool  {
	index := num / 4
	pos := ((num % 4) * 2) + 1

	return (b.bits[index] & (1 << pos)) != 0
}

func (b *RepeatBitMap) PrintNoRepeat()  {
	var i uint = 0
	for ; i < b.capacity; i ++ {
		fmt.Println("Number:",i," Is contain:",b.Contain(uint(i))," Is repeat:", b.IsRepeat(i))
	}
}