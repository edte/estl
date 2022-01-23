// @program:     tiny-stl
// @file:        string.go
// @author:      edte
// @create:      2022-01-22 22:17
// @description:
package string

import (
	"stl/allocator"
	"stl/containers"
	"stl/containers/vector"
	"stl/iterator"
)

type String struct {
	data *vector.Vector
}

func (s String) Swap(t *containers.Container) {
	//TODO implement me
	panic("implement me")
}

func (s String) Clone() containers.Container {
	//TODO implement me
	panic("implement me")
}

func (s String) Clear() {
	//TODO implement me
	panic("implement me")
}

func (s String) Size() int {
	//TODO implement me
	panic("implement me")
}

func (s String) MaxSize() int {
	//TODO implement me
	panic("implement me")
}

func (s String) Capacity() int {
	//TODO implement me
	panic("implement me")
}

func (s String) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (s String) GetAllocator() allocator.Allocator {
	//TODO implement me
	panic("implement me")
}

func (s String) Construct() {
	//TODO implement me
	panic("implement me")
}

func (s String) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (s String) String() string {
	//TODO implement me
	panic("implement me")
}

func (s String) Begin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) End() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) RBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) REnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) CBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) CEnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) CRBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) CREnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) ReSize(size int, value ...int) {
	//TODO implement me
	panic("implement me")
}

func (s String) Reserve(n int) {
	//TODO implement me
	panic("implement me")
}

func (s String) ShrinkToFit() {
	//TODO implement me
	panic("implement me")
}

func (s String) At(i int) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s String) Front() interface{} {
	//TODO implement me
	panic("implement me")
}

func (s String) Back() interface{} {
	//TODO implement me
	panic("implement me")
}

func (s String) Data() []interface{} {
	//TODO implement me
	panic("implement me")
}

func (s String) SetAt(i int, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s String) Assign(size int, value int) {
	//TODO implement me
	panic("implement me")
}

func (s String) AssignIter(a, b iterator.RandomAccessIterator) {
	//TODO implement me
	panic("implement me")
}

func (s String) PushBack(data interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s String) Insert(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) InsertSize(pos iterator.RandomAccessIterator, size int, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) InsertIter(pos iterator.RandomAccessIterator, first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) Erase(pos iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) EraseIter(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) Emplace(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) EmplaceBack(value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (s String) PopBack() {
	//TODO implement me
	panic("implement me")
}

func (s String) PopFront() {
	//TODO implement me
	panic("implement me")
}
