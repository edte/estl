// @program:     tiny-stl
// @file:        inputIterator.go
// @author:      edte
// @create:      2021-12-30 15:43
// @description:
package vector

import (
	"fmt"

	"stl/iterator"
)

type InputIterator struct {
	data *Vector
	pos  int
}

func NewInputIterator(data *Vector, pos int) *InputIterator {
	return &InputIterator{data: data, pos: pos}
}

func (iter *InputIterator) Next() iterator.InputIterator {
	if iter.pos == iter.data.Size() {
		return iter
	}
	iter.pos++
	return iter
}

func (iter *InputIterator) Clone() iterator.InputIterator {
	return NewInputIterator(iter.data, iter.pos)
}

func (iter *InputIterator) HasNext() bool {
	return iter.pos == iter.data.Size()
}

func (iter *InputIterator) Equal(other iterator.InputIterator) bool {
	i, ok := other.(*InputIterator)
	if !ok {
		return false
	}
	return i.data == iter.data && i.pos == iter.pos
}

func (iter *InputIterator) IsValid() bool {
	return iter.pos >= -1 && iter.pos <= iter.data.Size()
}

func (iter *InputIterator) Value() interface{} {
	return iter.data.At(iter.pos)
}

func (iter *InputIterator) String() string {
	return fmt.Sprint(iter.data.At(iter.pos))
}
