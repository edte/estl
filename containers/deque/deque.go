// @program:     tiny-stl
// @file:        deque.go
// @author:      edte
// @create:      2022-01-15 17:37
// @description:
package deque

import (
	"stl/allocator"
	"stl/containers"
	"stl/iterator"
)

type Deque struct {
}

func New() *Deque {
	return &Deque{}
}

func (d Deque) Swap(t *containers.Container) {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Clone() containers.Container {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Clear() {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Size() int {
	//TODO implement me
	panic("implement me")
}

func (d Deque) MaxSize() int {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Capacity() int {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (d Deque) GetAllocator() allocator.Allocator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Construct() {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (d Deque) String() string {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Begin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) End() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) RBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) REnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) CBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) CEnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) CRBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) CREnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) ReSize(size int, value ...int) {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Reserve(n int) {
	//TODO implement me
	panic("implement me")
}

func (d Deque) ShrinkToFit() {
	//TODO implement me
	panic("implement me")
}

func (d Deque) At(i int) interface{} {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Front() interface{} {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Back() interface{} {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Data() []interface{} {
	//TODO implement me
	panic("implement me")
}

func (d Deque) SetAt(i int, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Assign(size int, value int) {
	//TODO implement me
	panic("implement me")
}

func (d Deque) AssignIter(a, b iterator.RandomAccessIterator) {
	//TODO implement me
	panic("implement me")
}

func (d Deque) PushBack(data interface{}) {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Insert(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) InsertSize(pos iterator.RandomAccessIterator, size int, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) InsertIter(pos iterator.RandomAccessIterator, first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Erase(pos iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) EraseIter(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) Emplace(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) EmplaceBack(value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d Deque) PopBack() {
	//TODO implement me
	panic("implement me")
}

func (d Deque) PopFront() {
	//TODO implement me
	panic("implement me")
}
