// @program:     tiny-stl
// @file:        main.go
// @author:      edte
// @create:      2021-12-24 19:24
// @description:
package main

type Iterator interface {
	// 下一个迭代器
	Next() Iterator
	// 克隆迭代器
	Clone() Iterator

	// 判断后面是否还有迭代器
	HasNext() bool
	// 比较迭代器是否相等
	Equal(other Iterator) bool
	// 判断迭代器是否有效
	IsValid() bool
}

// InputIterator 可读向前遍历的迭代器
// find 和 accumulate
type InputIterator interface {
	Iterator

	// 解引用，获取迭代器的值
	Value() interface{}
}

// OutputIterator 可写向前遍历的迭代器
type OutputIterator interface {
	Iterator

	// 设置迭代器的值
	SetValue(value interface{})
}

// ForwardIterator 可读可写向前遍历的迭代器
type ForwardIterator interface {
	Iterator
	InputIterator
	OutputIterator
}

// BidirectionalIterator 可读可写双向遍历迭代器
// list, set, multiset, map, multimap
type BidirectionalIterator interface {
	Iterator
	InputIterator
	OutputIterator
	ForwardIterator

	// 上一个迭代器
	Pre() Iterator

	// 判断后面是否还有迭代器
	HasPre() bool
}

// RandomAccessIterator 可读可写随机访问迭代器
// vector, deque, string, array
type RandomAccessIterator interface {
	Iterator
	InputIterator
	OutputIterator
	ForwardIterator
	BidirectionalIterator

	// 后第 n 个迭代器
	NextN(n int) Iterator
	// 前第 n 个迭代器
	PreN(n int) Iterator
	// 获取 pos 位置迭代器
	IteratorAt(pos int) Iterator

	// iter1<iter2
	IsFront(other Iterator) bool
	// iter1>iter2
	IsBack(other Iterator) bool
	// iter1<=iter2
	IsFrontEqual(other Iterator) bool
	// iter1>=iter2
	IsBackEqual(other Iterator) bool
	// 判断前面是否还有迭代器
	HasPre() bool

	// 解引用，获取第 n 个迭代器的值
	ValueAt(n int) interface{}
	// 获取两迭代器间的距离
	Distance(other Iterator) int
	// 获取当前迭代器位置
	Position() int
}

func main() {
}
