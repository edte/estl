// @program:     tiny-stl
// @file:        list.go
// @author:      edte
// @create:      2021-12-25 00:33
// @description:
package list

import (
	"stl/allocator"
	"stl/containers"
	"stl/iterator"
)

type List struct {
}

func (l List) Swap(t *containers.Container) {
	//TODO implement me
	panic("implement me")
}

func (l List) Clone() containers.Container {
	//TODO implement me
	panic("implement me")
}

func (l List) Clear() {
	//TODO implement me
	panic("implement me")
}

func (l List) Size() int {
	//TODO implement me
	panic("implement me")
}

func (l List) MaxSize() int {
	//TODO implement me
	panic("implement me")
}

func (l List) Capacity() int {
	//TODO implement me
	panic("implement me")
}

func (l List) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (l List) GetAllocator() allocator.Allocator {
	//TODO implement me
	panic("implement me")
}

func (l List) Construct() {
	//TODO implement me
	panic("implement me")
}

func (l List) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (l List) String() string {
	//TODO implement me
	panic("implement me")
}

func (l List) Begin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) End() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) RBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) REnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) CBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) CEnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) CRBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) CREnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) ReSize(size int, value ...int) {
	//TODO implement me
	panic("implement me")
}

func (l List) Reserve(n int) {
	//TODO implement me
	panic("implement me")
}

func (l List) ShrinkToFit() {
	//TODO implement me
	panic("implement me")
}

func (l List) At(i int) interface{} {
	//TODO implement me
	panic("implement me")
}

func (l List) Front() interface{} {
	//TODO implement me
	panic("implement me")
}

func (l List) Back() interface{} {
	//TODO implement me
	panic("implement me")
}

func (l List) Data() []interface{} {
	//TODO implement me
	panic("implement me")
}

func (l List) SetAt(i int, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l List) Assign(size int, value int) {
	//TODO implement me
	panic("implement me")
}

func (l List) AssignIter(a, b iterator.RandomAccessIterator) {
	//TODO implement me
	panic("implement me")
}

func (l List) PushBack(data interface{}) {
	//TODO implement me
	panic("implement me")
}

func (l List) Insert(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) InsertSize(pos iterator.RandomAccessIterator, size int, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) InsertIter(pos iterator.RandomAccessIterator, first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) Erase(pos iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) EraseIter(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) Emplace(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) EmplaceBack(value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (l List) PopBack() {
	//TODO implement me
	panic("implement me")
}

func (l List) PopFront() {
	//TODO implement me
	panic("implement me")
}

func (l List) Slice() {
	//TODO implement me
	panic("implement me")
}

func (l List) Remove() {
	//TODO implement me
	panic("implement me")
}

func (l List) RemoveIf() {
	//TODO implement me
	panic("implement me")
}

func (l List) Unique() {
	//TODO implement me
	panic("implement me")
}

func (l List) Merge() {
	//TODO implement me
	panic("implement me")
}

func (l List) Sort() {
	//TODO implement me
	panic("implement me")
}
