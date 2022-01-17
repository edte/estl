// @program:     tiny-stl
// @file:        iterator.go
// @author:      edte
// @create:      2021-12-30 08:19
// @description:
package iterator

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
	// 下一个迭代器
	Next() InputIterator
	// 克隆迭代器
	Clone() InputIterator

	// 判断后面是否还有迭代器
	HasNext() bool
	// 比较迭代器是否相等
	Equal(other InputIterator) bool
	// 判断迭代器是否有效
	IsValid() bool

	// 解引用，获取迭代器的值
	Value() interface{}
}

// OutputIterator 可写向前遍历的迭代器
type OutputIterator interface {
	// 下一个迭代器
	Next() OutputIterator
	// 克隆迭代器
	Clone() OutputIterator

	// 判断后面是否还有迭代器
	HasNext() bool
	// 比较迭代器是否相等
	Equal(other OutputIterator) bool
	// 判断迭代器是否有效
	IsValid() bool

	// 设置迭代器的值
	SetValue(value interface{})
}

// ForwardIterator 可读可写向前遍历的迭代器
type ForwardIterator interface {
	// 下一个迭代器
	Next() ForwardIterator
	// 克隆迭代器
	Clone() ForwardIterator

	// 判断后面是否还有迭代器
	HasNext() bool
	// 比较迭代器是否相等
	Equal(other ForwardIterator) bool
	// 判断迭代器是否有效
	IsValid() bool

	// 解引用，获取迭代器的值
	Value() interface{}
	// 设置迭代器的值
	SetValue(value interface{})
}

// BidirectionalIterator 可读可写双向遍历迭代器
// list, set, multiset, map, multimap
type BidirectionalIterator interface {
	// 下一个迭代器
	Next() BidirectionalIterator
	// 上一个迭代器
	Pre() BidirectionalIterator
	// 克隆迭代器
	Clone() BidirectionalIterator

	// 判断后面是否还有迭代器
	HasNext() bool
	// 判断后面是否还有迭代器
	HasPre() bool
	// 比较迭代器是否相等
	Equal(other BidirectionalIterator) bool
	// 判断迭代器是否有效
	IsValid() bool

	// 解引用，获取迭代器的值
	Value() interface{}
	// 设置迭代器的值
	SetValue(value interface{})
}

// RandomAccessIterator 可读可写随机访问迭代器
// vector, deque, string, array
type RandomAccessIterator interface {
	// 下一个迭代器
	Next() RandomAccessIterator
	// 后第 n 个迭代器
	NextN(n int) RandomAccessIterator
	// 上一个迭代器
	Pre() RandomAccessIterator
	// 前第 n 个迭代器
	PreN(n int) RandomAccessIterator
	// 获取 pos 位置迭代器
	IteratorAt(pos int) RandomAccessIterator
	// 克隆迭代器
	Clone() RandomAccessIterator

	// iter1=iter2
	Equal(other RandomAccessIterator) bool
	// iter1<iter2
	IsFront(other RandomAccessIterator) bool
	// iter1>iter2
	IsBack(other RandomAccessIterator) bool
	// iter1<=iter2
	IsFrontEqual(other RandomAccessIterator) bool
	// iter1>=iter2
	IsBackEqual(other RandomAccessIterator) bool
	// 判断后面是否还有迭代器
	HasNext() bool
	// 判断前面是否还有迭代器
	HasPre() bool
	// 判断迭代器是否有效
	IsValid() bool

	// 解引用，获取迭代器的值
	Value() interface{}
	// 解引用，获取第 n 个迭代器的值
	ValueAt(n int) interface{}
	// 设置迭代器的值
	SetValue(value interface{})
	// 获取两迭代器间的距离
	Distance(other RandomAccessIterator) int
	// 获取当前迭代器位置
	Position() int
}
