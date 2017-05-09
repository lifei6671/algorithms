package bitmap

import "fmt"

//实现BitMap算法.
//@link http://www.jianshu.com/p/6082a2f7df8e#.
type BitMap struct {
	bits []byte
	capacity int
}

// 初始化一个 BitMap 类型，因为一个 byte 可以存放 8 个bit，所以只需要创建容量为 (N/8)+1 的数组即可.
func NewBitMap(capacity int) *BitMap {
	return &BitMap{
		bits : make([]byte,(capacity >> 3)+1),
		capacity: capacity,
	}
}

//添加一个元素，该元素所在 byte 数组的位置为 num/8 ，在 byte 字节中的位置为 num%8 ，1 左移 pos 位为该元素在该byte中所在的位置，
// 和 byte 做或运算，就是当前byte字节的值。
func (bm *BitMap) Add(num uint) {
	index := num >> 3
	pos := num & 0x07

	bm.bits[index] = bm.bits[index] | (1 << pos)
}

//移除一个元素，将byte字节和 1 左移取非 做与运算.
func (bm *BitMap) Remove(num uint)  {
	index := num >> 3
	pos := num & 0x07

	bm.bits[index] = bm.bits[index] & ^(1 << pos)
}

//判断元素是否存在，byte和 1 左移 pos 位置做与运算，如果不等于0表示存在.
func (b *BitMap) Contain(num uint) bool {
	index := num >> 3
	pos := num & 0x07

	return b.bits[index] & (1<<pos) != 0
}



func (b *BitMap)String() string {
	return fmt.Sprint(b.bits)
}