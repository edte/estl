// @program:     tiny-stl
// @file:        iterator.go
// @author:      edte
// @create:      2021-12-30 12:47
// @description:
package vector

import (
	"fmt"

	"stl/iterator"
)

type Iterator struct {
	InputIterator
	OutputIterator
	ForwardIterator
	data *Vector
	pos  int
}

func NewIterator(data *Vector, pos int) *Iterator {
	return &Iterator{data: data, pos: pos}
}

func (iter *Iterator) Next() iterator.RandomAccessIterator {
	if iter.pos == iter.data.Size() {
		return iter
	}
	iter.pos++
	return iter
}

func (iter *Iterator) NextN(n int) iterator.RandomAccessIterator {
	if iter.pos+n == iter.data.Size() {
		return iter
	}
	iter.pos += n
	return iter
}

func (iter *Iterator) Pre() iterator.RandomAccessIterator {
	if iter.pos == -1 {
		return iter
	}
	iter.pos--
	return iter
}

func (iter *Iterator) PreN(n int) iterator.RandomAccessIterator {
	// if iter.pos-n == -1 {
	// 	return iter
	// }
	iter.pos -= n
	return iter
}

func (iter *Iterator) IteratorAt(pos int) iterator.RandomAccessIterator {
	// if pos < 0 {
	// 	pos = 0
	// }
	// if pos > iter.data.Size() {
	// 	pos = iter.data.Size()
	// }
	iter.pos = pos
	return iter
}

func (iter *Iterator) Clone() iterator.RandomAccessIterator {
	return NewIterator(iter.data, iter.pos)
}

func (iter *Iterator) Equal(other iterator.RandomAccessIterator) bool {
	i, ok := other.(*Iterator)
	if !ok {
		return false
	}
	return i.data == iter.data && i.pos == iter.pos
}

func (iter *Iterator) IsFront(other iterator.RandomAccessIterator) bool {
	i, ok := other.(*Iterator)
	if !ok {
		return false
	}
	return i.data == iter.data && i.pos > iter.pos
}

func (iter *Iterator) IsBack(other iterator.RandomAccessIterator) bool {
	i, ok := other.(*Iterator)
	if !ok {
		return false
	}
	return i.data == iter.data && i.pos < iter.pos

}

func (iter *Iterator) IsFrontEqual(other iterator.RandomAccessIterator) bool {
	i, ok := other.(*Iterator)
	if !ok {
		return false
	}
	return i.data == iter.data && i.pos >= iter.pos
}

func (iter *Iterator) IsBackEqual(other iterator.RandomAccessIterator) bool {
	i, ok := other.(*Iterator)
	if !ok {
		return false
	}
	return i.data == iter.data && i.pos <= iter.pos
}

func (iter *Iterator) HasNext() bool {
	return iter.pos < iter.data.Size()
}

func (iter *Iterator) HasPre() bool {
	return iter.pos > -1
}

func (iter *Iterator) IsValid() bool {
	return iter.pos > -1 && iter.pos <= iter.data.Size()
}

func (iter *Iterator) Value() interface{} {
	return iter.data.At(iter.pos)
}

func (iter *Iterator) ValueAt(n int) interface{} {
	return iter.data.At(n)
}

func (iter *Iterator) SetValue(value interface{}) {
	iter.data.SetAt(iter.pos, value)
}

func (iter *Iterator) Distance(other iterator.RandomAccessIterator) int {
	i, ok := other.(*Iterator)
	if !ok {
		return -1
	}
	// 求绝对值
	return (i.pos - iter.pos) ^ ((i.pos - iter.pos) >> 31) - ((i.pos - iter.pos) >> 31)
}

func (iter *Iterator) Position() int {
	return iter.pos
}

func (iter *Iterator) String() string {
	return fmt.Sprint(iter.Value())
}
